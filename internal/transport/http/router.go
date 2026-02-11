package http

import (
	"context"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"vago/internal/app"
	"vago/internal/application/chapter"
	"vago/internal/application/chat"
	"vago/internal/application/comment"
	"vago/internal/application/task"
	"vago/internal/application/test"
	"vago/internal/application/topic"
	"vago/internal/application/user"
	"vago/internal/config/route"
	"vago/internal/infra/persistence/gorm"
	"vago/internal/infra/token"
	"vago/internal/transport/http/handler"
	"vago/internal/transport/http/middleware"
	"vago/internal/transport/ws"

	"github.com/gin-gonic/gin"
)

func SetupRouter(goCtx context.Context, ctx *app.Context, tokenProvider *token.JWTProvider) *gin.Engine {
	// WS
	hub := ws.NewHub(ctx.Log)
	go hub.Run(goCtx)
	// Сервисы
	taskSvc := task.NewService(gorm.NewTaskRepo(ctx.DB))
	messageRepo := gorm.NewMessageRepo(ctx.DB)
	userRepo := gorm.NewUserRepo(ctx)
	chatSvc := chat.NewService(messageRepo, userRepo)

	userSvc := user.NewService(userRepo, tokenProvider)
	localCache := app.NewLocalCache()

	// Хендлеры
	topicRepo := gorm.NewTopicRepo(ctx.DB)
	questionSvc := test.NewService(gorm.NewQuestionRepo(ctx.DB), topicRepo)
	chapterSvc := chapter.NewService(gorm.NewChapterRepo(ctx.DB))
	topicSvc := topic.NewService(topicRepo)
	commentSvc := comment.NewService(gorm.NewCommentRepo(ctx.DB))

	authH := handler.NewAuthHandler(userSvc, ctx.Cfg.JwtSecret, ctx.Cfg.RefreshTTLInt(), ctx.Log)
	testH := handler.NewTestHandler(questionSvc, chapterSvc, topicSvc, commentSvc, ctx.Cfg.PostgresDsn)
	adminH := handler.NewAdminHandler(tokenProvider, userSvc, chatSvc, commentSvc)
	commentH := handler.NewCommentHandler(commentSvc)

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
	r.Use(middleware.LoadUserContext(userSvc, localCache))
	r.Use(middleware.NoCache)
	r.Use(middleware.TemplateContext)

	// Публичные маршруты
	r.GET(route.Index, handler.ShowIndex())
	r.GET(route.Book, handler.ShowBook)
	r.GET(route.Login, handler.ShowLogin)
	r.POST(route.Login, authH.Login)
	r.GET(route.Register, handler.ShowSignup)
	r.POST(route.Register, handler.SignUp(userSvc))
	r.GET(route.SignOut, handler.SignOut)

	r.GET("/test", testH.ShowTestRandom())
	r.GET("/test/:id", testH.ShowTestByID())
	r.POST("/test/check", testH.Check())

	r.GET(route.Questions, testH.ShowQuestions)

	// Защищенные маршруты
	auth := r.Group("/")
	auth.Use(middleware.RequireAuthAndRedirect)
	{
		admin := auth.Group(route.Admin)
		{
			admin.GET("", adminH.ShowAdmin)
			admin.GET(route.User, adminH.ShowUser)
			admin.GET(route.Comments, adminH.ShowComments)
			admin.GET(route.Users, adminH.Users)
			admin.GET(route.Messages, adminH.ShowMessages)
			admin.GET(route.Grpc, adminH.ShowGrpc)
		}

		auth.GET(route.Tasks, handler.Tasks(taskSvc))
		auth.POST(route.Tasks, handler.PostTask(taskSvc))
		auth.DELETE(route.Tasks+"/:id", handler.DeleteTask(ctx))
		auth.PUT(route.Tasks+"/:id", handler.UpdateTask(taskSvc))
		auth.DELETE("/users/:id", handler.DeleteUser(userSvc))

		auth.GET("/ws", handler.ServeSW(hub, ctx.Log, tokenProvider, chatSvc))
		auth.GET("/chat", handler.ShowChat(ctx.Cfg.Port, chatSvc))

		messagesHandler := handler.NewMessageHandler(chatSvc, userSvc)
		auth.POST("/messages", messagesHandler.AddMessage())
		auth.POST("/messagesDeleteAll", messagesHandler.DeleteAll())
		auth.DELETE("/messages/:id", messagesHandler.Delete())

		auth.GET("/add_questions", testH.ShowAddQuestion())
		auth.POST("/add_questions", testH.AddQuestion())

		auth.POST("/run_questions_seed", testH.RunGoQuestionsSeed())
		auth.POST(route.RunGoTopicsSeed, testH.RunGoTopicsSeed())

		auth.POST("/comments", commentH.PostComment)
	}

	// ========= API =========
	apiGroup := r.Group("/api")
	apiGroup.GET(route.Me, authH.MeAPI)
	apiGroup.POST(route.SignIn, authH.SignInAPI)
	apiGroup.POST(route.SignUp, handler.SignUpApi(userSvc))
	apiGroup.GET(route.SignOut, handler.SignOut)
	apiGroup.GET(route.Questions, testH.ShowQuestionsAPI)

	// Защищенные маршруты (API)
	apiGroup.Use(middleware.RequireAuthApi)
	{
		apiGroup.GET(route.Users, adminH.UsersApi)
		apiGroup.DELETE(route.Users+"/:id", handler.DeleteUser(userSvc))

		apiGroup.POST(route.Tasks, handler.PostTaskAPI(taskSvc))
		apiGroup.DELETE(route.Tasks+"/:id", handler.DeleteTaskAPI(taskSvc))
		apiGroup.PUT(route.Tasks+"/:id", handler.UpdateTaskAPI(taskSvc))

		apiGroup.GET(route.Tasks, handler.TasksAPI(taskSvc))
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
