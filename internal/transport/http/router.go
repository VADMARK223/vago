package http

import (
	"context"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"vago/internal/app"
	"vago/internal/application/chat"
	"vago/internal/application/quiz"
	"vago/internal/application/task"
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
	questionSvc := quiz.NewService(gorm.NewQuestionRepo(ctx.DB), topicRepo)
	topicSvc := topic.NewService(topicRepo)
	authH := handler.NewAuthHandler(userSvc, ctx.Cfg.JwtSecret, ctx.Cfg.RefreshTTLInt(), ctx.Log)
	quizHandler := handler.NewQuizHandler(questionSvc, topicSvc, ctx.Cfg.PostgresDsn)
	adminHandler := handler.NewAdminHandler(tokenProvider, userSvc, chatSvc)

	gin.SetMode(ctx.Cfg.GinMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	// Шаблоны
	r.SetHTMLTemplate(loadTemplates("web/templates"))
	_ = r.SetTrustedProxies(nil)
	// Статика и шаблоны
	r.Static("/static", "/app/web/static")
	// Favicon: отдаём напрямую, чтобы не было 404
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("web/static/favicon.ico")
	})

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
	r.POST(route.Register, handler.PerformRegister(userSvc))
	r.POST(route.Logout, handler.Logout)

	r.GET("/quiz", quizHandler.ShowQuizRandom())
	r.GET("/quiz/:id", quizHandler.ShowQuizByID())
	r.POST("/quiz/check", quizHandler.Check())

	r.GET("/questions", quizHandler.ShowQuestions)

	// Защищенные маршруты
	auth := r.Group("/")
	auth.Use(middleware.CheckAuthAndRedirect())
	{
		admin := auth.Group(route.Admin)
		{
			admin.GET(route.User, adminHandler.ShowUser)
			admin.GET(route.Users, adminHandler.ShowUsers)
			admin.GET(route.Messages, adminHandler.ShowMessages)
			admin.GET(route.Grpc, adminHandler.ShowGrpc)
		}

		auth.GET(route.Tasks, handler.Tasks(taskSvc))
		auth.POST(route.Tasks, handler.AddTask(ctx))
		auth.DELETE("/tasks/:id", handler.DeleteTask(ctx))
		auth.PUT("/tasks/:id", handler.UpdateTask(ctx, taskSvc))
		auth.DELETE("/users/:id", handler.DeleteUser(userSvc))

		auth.GET("/ws", handler.ServeSW(hub, ctx.Log, tokenProvider, chatSvc))
		auth.GET("/chat", handler.ShowChat(ctx.Cfg.Port, chatSvc))

		messagesHandler := handler.NewMessageHandler(chatSvc, userSvc)
		auth.POST("/messages", messagesHandler.AddMessage())
		auth.POST("/messagesDeleteAll", messagesHandler.DeleteAll())
		auth.DELETE("/messages/:id", messagesHandler.Delete())

		auth.GET("/add_questions", quizHandler.ShowAddQuestion())
		auth.POST("/add_questions", quizHandler.AddQuestion())

		auth.POST("/runTopicsSeed", quizHandler.RunTopicsSeed())
		auth.POST("/run_questions_seed", quizHandler.RunQuestionsSeedNew())
	}

	r.NoRoute(handler.NotFoundHandler)

	return r
}

func loadTemplates(root string) *template.Template {
	tmpl := template.New("").Funcs(template.FuncMap{
		"dict": dict,
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
