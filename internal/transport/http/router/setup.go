package router

import (
	"context"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"vago/internal/app"
	"vago/internal/config/route"
	"vago/internal/infra/token"
	apiq "vago/internal/transport/http/api/question"
	"vago/internal/transport/http/handler"
	"vago/internal/transport/http/middleware"
	webq "vago/internal/transport/http/web/question"

	"github.com/gin-gonic/gin"
)

func SetupRouter(goCtx context.Context, ctx *app.Context, tokenProvider *token.JWTProvider) *gin.Engine {
	deps := buildDeps(goCtx, ctx, tokenProvider)

	gin.SetMode(ctx.Cfg.GinMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	// Статика и шаблоны
	r.Static("/static", "web/static")
	// Шаблоны
	r.SetHTMLTemplate(loadTemplates("web/templates"))
	_ = r.SetTrustedProxies(nil)
	// Favicon: отдаём напрямую, чтобы не было 404
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("web/static/favicon.ico")
	})

	registerReactApp(r)

	// Middleware
	r.Use(middleware.SessionMiddleware())
	r.Use(middleware.CheckJWT(tokenProvider, ctx.Cfg.RefreshTTLInt()))
	r.Use(middleware.LoadUserContext(deps.Services.User, deps.Cache))
	r.Use(middleware.NoCache)
	r.Use(middleware.TemplateContext)

	// Публичные маршруты
	r.GET(route.Index, handler.ShowIndex())
	r.GET(route.Book, handler.ShowBook)
	r.GET(route.Login, handler.ShowLogin)
	r.POST(route.Login, deps.Handlers.Auth.Login)
	r.GET(route.Register, handler.ShowSignup)
	r.POST(route.Register, handler.SignUp(deps.Services.User))
	r.GET(route.SignOut, handler.SignOut)

	r.GET(route.Test, deps.Handlers.Test.ShowRandom)
	r.GET(route.Test+"/:id", deps.Handlers.Test.ShowByID())
	r.POST(route.Test+"/check", deps.Handlers.Test.CheckAnswer)

	webQ := webq.New(deps.Loaders.Question)
	r.GET(route.Questions, webQ.Page)

	// Защищенные маршруты
	auth := r.Group("/")
	auth.Use(middleware.RequireAuthAndRedirect)
	{
		admin := auth.Group(route.Admin)
		{
			admin.GET("", deps.Handlers.Admin.ShowAdmin)
			admin.GET(route.User, deps.Handlers.Admin.ShowUser)
			admin.GET(route.Comments, deps.Handlers.Admin.ShowComments)
			admin.GET(route.Users, deps.Handlers.Admin.Users)
			admin.GET(route.Messages, deps.Handlers.Admin.ShowMessages)
			admin.GET(route.Grpc, deps.Handlers.Admin.ShowGrpc)
		}

		auth.GET(route.Tasks, handler.Tasks(deps.Services.Task))
		auth.POST(route.Tasks, handler.PostTask(deps.Services.Task))
		auth.DELETE(route.Tasks+"/:id", handler.DeleteTask(ctx.DB))
		auth.PUT(route.Tasks+"/:id", handler.UpdateTask(deps.Services.Task))
		auth.DELETE("/users/:id", handler.DeleteUser(deps.Services.User))

		auth.GET("/ws", handler.ServeSW(deps.Hub, ctx.Log, tokenProvider, deps.Services.Chat))
		auth.GET("/chat", handler.ShowChat(ctx.Cfg.Port, deps.Services.Chat))

		messagesHandler := handler.NewMessageHandler(deps.Services.Chat, deps.Services.User)
		auth.POST("/messages", messagesHandler.AddMessage())
		auth.POST("/messagesDeleteAll", messagesHandler.DeleteAll())
		auth.DELETE("/messages/:id", messagesHandler.Delete())

		auth.GET(route.AddQuestions, deps.Handlers.TestEditor.ShowAddQuestion)
		auth.POST(route.AddQuestions, deps.Handlers.TestEditor.AddQuestion)
		auth.POST(route.RunQuestionsSeed, deps.Handlers.TestEditor.RunGoQuestionsSeed)
		auth.POST(route.RunGoTopicsSeed, deps.Handlers.TestEditor.RunGoTopicsSeed)

		auth.POST(route.Comments, deps.Handlers.Comment.PostComment)
	}

	// ===== temp block =====

	// ========= API =========
	apiGroup := r.Group("/api")
	apiGroup.GET(route.Me, deps.Handlers.Auth.MeAPI)
	apiGroup.POST(route.SignIn, deps.Handlers.Auth.SignInAPI)
	apiGroup.POST(route.SignUp, handler.SignUpApi(deps.Services.User))
	apiGroup.GET(route.SignOut, handler.SignOut)

	apiQ := apiq.New(deps.Loaders.Question)
	apiGroup.GET(route.Questions, apiQ.Get)

	// Защищенные маршруты (API)
	apiGroup.Use(middleware.RequireAuthApi)
	{
		apiGroup.GET(route.Users, deps.Handlers.Admin.UsersApi)
		apiGroup.DELETE(route.Users+"/:id", handler.DeleteUser(deps.Services.User))

		apiGroup.GET(route.Tasks, handler.TasksAPI(deps.Services.Task))
		apiGroup.POST(route.Tasks, handler.PostTaskAPI(deps.Services.Task))
		apiGroup.DELETE(route.Tasks+"/:id", handler.DeleteTaskAPI(deps.Services.Task))
		apiGroup.PUT(route.Tasks+"/:id", handler.UpdateTaskAPI(deps.Services.Task))

		apiGroup.GET(route.Test, deps.Handlers.Test.RandomQuestionIdAPI)
		apiGroup.GET(route.Test+"/:id", deps.Handlers.Test.QuestionByIdAPI)
		apiGroup.POST(route.Test+"/check", deps.Handlers.Test.CheckAnswerAPI)
	}

	r.NoRoute(handler.NotFoundHandler)

	return r
}

func loadTemplates(root string) *template.Template {
	tmpl := template.New("").Funcs(template.FuncMap{
		"dict": dict,
		"menuActive": func(path, href string) bool {
			if href == "/" {
				return path == "/"
			}
			return strings.HasPrefix(path, href)
		},
	})

	walkErr := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".gohtml") {
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}

			relPath = filepath.ToSlash(relPath)

			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			_, err = tmpl.New(relPath).Parse(string(content))
			if err != nil {
				return err
			}
		}
		return nil
	})

	if walkErr != nil {
		panic(walkErr)
	}

	return tmpl
}

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("invalid dict call: odd number of args")
	}
	m := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}
		m[key] = values[i+1]
	}
	return m, nil
}

func registerReactApp(r *gin.Engine) {
	distDir := "./web/v2/dist"

	r.GET("/v2/*path", func(c *gin.Context) {
		// Например: "/assets/index-xxxx.js" или "/courses" или "/"
		p := strings.TrimPrefix(c.Param("path"), "/")

		// Если запросили конкретный файл — отдадим его
		if p != "" {
			fp := filepath.Join(distDir, p)
			if st, err := os.Stat(fp); err == nil && !st.IsDir() {
				c.File(fp)
				return
			}
		}

		// Иначе — SPA fallback
		c.File(filepath.Join(distDir, "index.html"))
	})
}
