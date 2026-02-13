package router

import (
	"vago/internal/config/route"
	apiq "vago/internal/transport/http/api/question"
	"vago/internal/transport/http/handler"
	"vago/internal/transport/http/middleware"

	"github.com/gin-gonic/gin"
)

func registerAPIRoutes(api *gin.RouterGroup, deps *Deps) {
	apiHandlerQuestion := apiq.New(deps.Loaders.Question)

	// Public
	api.GET(route.Me, deps.Handlers.Auth.MeAPI)
	api.POST(route.SignIn, deps.Handlers.Auth.SignInAPI)
	api.POST(route.SignUp, handler.SignUpApi(deps.Services.User))
	api.GET(route.SignOut, handler.SignOut)
	api.GET(route.Questions, apiHandlerQuestion.Get)

	// Protected
	apiAuth := api.Group("")
	apiAuth.Use(middleware.RequireAuthApi)

	apiAuth.GET(route.Users, deps.Handlers.Admin.UsersApi)
	apiAuth.DELETE(route.Users+"/:id", handler.DeleteUser(deps.Services.User))

	apiAuth.GET(route.Tasks, handler.TasksAPI(deps.Services.Task))
	apiAuth.POST(route.Tasks, handler.PostTaskAPI(deps.Services.Task))
	apiAuth.DELETE(route.Tasks+"/:id", handler.DeleteTaskAPI(deps.Services.Task))
	apiAuth.PUT(route.Tasks+"/:id", handler.UpdateTaskAPI(deps.Services.Task))

	apiAuth.GET(route.Test, deps.Handlers.Test.RandomQuestionIdAPI)
	apiAuth.GET(route.Test+"/:id", deps.Handlers.Test.QuestionByIdAPI)
	apiAuth.POST(route.Test+"/check", deps.Handlers.Test.CheckAnswerAPI)
}
