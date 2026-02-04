package handler

import (
	"net/http"
	user2 "vago/internal/application/user"
	"vago/internal/config/code"
	"vago/internal/config/route"
	"vago/internal/domain"
	"vago/internal/transport/http/api"
	"vago/internal/transport/http/dto"
	"vago/pkg/strx"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	service    *user2.Service
	secret     string
	refreshTTL int
	log        *zap.SugaredLogger
}

type SignInReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignUpReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func NewAuthHandler(service *user2.Service, secret string, refreshTTL int, log *zap.SugaredLogger) *AuthHandler {
	return &AuthHandler{
		service:    service,
		secret:     secret,
		refreshTTL: refreshTTL,
		log:        log,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")

	_, tokens, err := h.service.Login(login, password)
	if err != nil {
		c.Set(code.Error, strx.Capitalize(err.Error()))
		ShowLogin(c)
		return
	}

	domain.SetTokenCookies(c, tokens, h.refreshTTL)

	session := sessions.Default(c)

	redirectTo := session.Get(code.RedirectTo)
	if redirectTo == nil {
		redirectTo = route.Index
	} else {
		session.Delete(code.RedirectTo)
	}

	_ = session.Save()

	c.Redirect(http.StatusFound, redirectTo.(string))
}

func (h *AuthHandler) MeAPI(c *gin.Context) {
	//time.Sleep(3 * time.Second)
	uAny, ok := c.Get(code.CurrentUser)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	u, ok := uAny.(domain.User)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid user in context"})
		return
	}

	api.OK(c, "Пользователь", dto.MeDTO{Username: u.Username, Role: u.Role})
}

func (h *AuthHandler) SignInAPI(c *gin.Context) {
	//time.Sleep(3 * time.Second)
	var req SignInReq
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, "Некорректные данные")
		return
	}

	_, tokens, err := h.service.Login(req.Login, req.Password)
	if err != nil {
		api.Error(c, http.StatusUnauthorized, strx.Capitalize(err.Error()))
		return
	}

	domain.SetTokenCookies(c, tokens, h.refreshTTL)
	api.OKNoData(c, "Успешный вход!")
}

func SignUp(service *user2.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		login := c.PostForm(code.Login)
		password := c.PostForm(code.Password)
		username := c.PostForm(code.Username)

		role := c.PostForm(code.Role)
		color := c.PostForm(code.Color)

		err := service.CreateUser(domain.DTO{Login: login, Password: password, Role: domain.Role(role), Color: color, Username: username})

		if err != nil {
			c.Set(code.Error, strx.Capitalize(err.Error()))
			ShowSignup(c)
			return
		}

		c.Redirect(http.StatusFound, "/login")
	}
}

func SignUpApi(service *user2.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SignUpReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Некорректные данные",
			})
			return
		}

		role := "user"
		color := "#FF5733"

		err := service.CreateUser(domain.DTO{Login: req.Login, Password: req.Password, Role: domain.Role(role), Color: color, Username: req.Username})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": strx.Capitalize(err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Успешная регистрация!",
		})
	}
}
