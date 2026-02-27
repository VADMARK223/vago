package router

import (
	"vago/internal/config/route"
	"vago/internal/transport/http/handler"
	"vago/internal/transport/http/middleware"
	webm "vago/internal/transport/http/web/message"
	webq "vago/internal/transport/http/web/question"

	"github.com/gin-gonic/gin"
)

func registerWebRoutes(web *gin.RouterGroup, deps *Deps) {
	// Public
	web.GET(route.Index, handler.ShowIndex)
	web.GET(route.Book, handler.ShowBook)
	web.GET(route.Login, handler.ShowLogin)
	web.POST(route.Login, deps.Handlers.Auth.Login)
	web.GET(route.Register, handler.ShowSignup)
	web.POST(route.Register, handler.SignUp(deps.Services.User))
	web.GET(route.SignOut, handler.SignOut)

	web.GET(route.Test, deps.Handlers.Test.ShowRandom)
	web.GET(route.Test+"/:id", deps.Handlers.Test.ShowByID)
	web.POST(route.Test+"/check", deps.Handlers.Test.CheckAnswer)

	webHandlerQuestion := webq.New(deps.Loaders.Question)
	web.GET(route.Questions, webHandlerQuestion.Page)

	webHandlerMessage := webm.New(deps.Loaders.Message, deps.Services.Message)
	// Protected
	auth := web.Group("")
	auth.Use(middleware.RequireAuthAndRedirect)
	{
		admin := auth.Group(route.Admin)
		{
			admin.GET("", deps.Handlers.Admin.ShowAdmin)
			admin.GET(route.User, deps.Handlers.Admin.ShowUser)
			admin.GET(route.Comments, deps.Handlers.Admin.ShowComments)
			admin.GET(route.Users, deps.Handlers.Admin.Users)
			admin.GET(route.Messages, webHandlerMessage.Page)
			admin.GET(route.Grpc, deps.Handlers.Admin.ShowGrpc)
		}
		auth.DELETE("/users/:id", handler.DeleteUser(deps.Services.User))

		auth.GET(route.Tasks, handler.Tasks(deps.Services.Task))
		auth.POST(route.Tasks, handler.PostTask(deps.Services.Task))
		auth.DELETE(route.Tasks+"/:id", handler.DeleteTask(deps.Services.Task))
		auth.PUT(route.Tasks+"/:id", handler.UpdateTask(deps.Services.Task))

		auth.GET("/ws", handler.ServeSW(deps.Hub, deps.Log, deps.TokenProv, deps.Services.Chat))
		auth.GET("/chat", handler.ShowChat(deps.TokenProv.Port, deps.Services.Message))

		auth.GET(route.AddQuestions, deps.Handlers.TestEditor.ShowAddQuestion)
		auth.POST(route.AddQuestions, deps.Handlers.TestEditor.AddQuestion)
		auth.POST(route.RunQuestionsSeed, deps.Handlers.TestEditor.RunGoQuestionsSeed)
		auth.POST(route.RunGoTopicsSeed, deps.Handlers.TestEditor.RunGoTopicsSeed)

		auth.POST(route.Comments, deps.Handlers.Comment.PostComment)

		auth.POST("/messagesDeleteAll", webHandlerMessage.DeleteAll)
		auth.DELETE("/messages/:id", webHandlerMessage.Delete)
	}
}
