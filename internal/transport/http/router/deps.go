package router

import (
	"context"
	"vago/internal/app"
	"vago/internal/application/chapter"
	"vago/internal/application/chat"
	"vago/internal/application/comment"
	"vago/internal/application/task"
	"vago/internal/application/test"
	"vago/internal/application/topic"
	"vago/internal/application/user"
	"vago/internal/infra/gorm"
	"vago/internal/infra/token"
	"vago/internal/transport/http/handler"
	questionLoader "vago/internal/transport/http/shared/question"
	"vago/internal/transport/ws"
)

type Services struct {
	Task    *task.Service
	User    *user.Service
	Chat    *chat.Service
	Topic   *topic.Service
	Chapter *chapter.Service
	Test    *test.Service
	Comment *comment.Service
}
type Handlers struct {
	Auth       *handler.AuthHandler
	Test       *handler.TestHandler
	TestEditor *handler.TestEditorHandler
	Admin      *handler.AdminHandler
	Comment    *handler.CommentHandler
	Message    *handler.MessageHandler
}

type Loaders struct {
	Question questionLoader.Loader
}

type Deps struct {
	Hub       *ws.Hub
	Services  Services
	Handlers  Handlers
	Loaders   Loaders
	Cache     *app.LocalCache
	TokenProv *token.JWTProvider
}

func buildDeps(goCtx context.Context, ctx *app.Context, tokenProvider *token.JWTProvider) *Deps {
	// WS hub
	hub := ws.NewHub(ctx.Log)
	go hub.Run(goCtx)

	// repos
	taskRepo := gorm.NewTaskRepo(ctx.DB)
	messageRepo := gorm.NewMessageRepo(ctx.DB)
	topicRepo := gorm.NewTopicRepo(ctx.DB)
	questionRepo := gorm.NewQuestionRepo(ctx.DB)
	chapterRepo := gorm.NewChapterRepo(ctx.DB)
	commentRepo := gorm.NewCommentRepo(ctx.DB)
	userRepo := gorm.NewUserRepo(ctx)

	// services
	taskSvc := task.NewService(taskRepo)
	chatSvc := chat.NewService(messageRepo, userRepo)
	userSvc := user.NewService(userRepo, tokenProvider)
	testSvc := test.NewService(questionRepo, topicRepo)
	chapterSvc := chapter.NewService(chapterRepo)
	topicSvc := topic.NewService(topicRepo)
	commentSvc := comment.NewService(commentRepo)

	// cache
	localCache := app.NewLocalCache()

	// handlers
	authH := handler.NewAuthHandler(userSvc, ctx.Cfg.JwtSecret, ctx.Cfg.RefreshTTLInt(), ctx.Log)
	testH := handler.NewTestHandler(testSvc, topicSvc, commentSvc)
	testEditorH := handler.NewTestEditorHandler(testSvc, topicSvc, ctx.Cfg.PostgresDsn)
	adminH := handler.NewAdminHandler(tokenProvider, userSvc, chatSvc, commentSvc)
	commentH := handler.NewCommentHandler(commentSvc)

	// loaders
	qLoader := questionLoader.Loader{
		ChapterSvc: chapterSvc,
		TopicSvc:   topicSvc,
		TestSvc:    testSvc,
	}

	return &Deps{
		Hub: hub,
		Services: Services{
			Task:    taskSvc,
			User:    userSvc,
			Chat:    chatSvc,
			Topic:   topicSvc,
			Chapter: chapterSvc,
			Test:    testSvc,
			Comment: commentSvc,
		},
		Handlers: Handlers{
			Auth:       authH,
			Test:       testH,
			TestEditor: testEditorH,
			Admin:      adminH,
			Comment:    commentH,
		},
		Loaders: Loaders{
			Question: qLoader,
		},
		Cache:     localCache,
		TokenProv: tokenProvider,
	}
}
