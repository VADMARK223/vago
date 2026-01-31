package handler

import (
	"net/http"
	user2 "vago/internal/application/user"
	"vago/internal/config/code"
	"vago/internal/config/route"
	"vago/internal/domain"
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

type LoginReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
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

func (h *AuthHandler) LoginAPI(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректные данные",
		})
		return
	}

	_, tokens, err := h.service.Login(req.Login, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": strx.Capitalize(err.Error()),
		})
		return
	}

	domain.SetTokenCookies(c, tokens, h.refreshTTL)

	c.JSON(http.StatusOK, gin.H{
		"message": "Успешный вход!",
	})
}

func PerformRegister(service *user2.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		login := c.PostForm(code.Login)
		password := c.PostForm(code.Password)
		username := c.PostForm(code.Username)
		email := c.PostForm(code.Email)

		role := c.PostForm(code.Role)
		color := c.PostForm(code.Color)

		err := service.CreateUser(domain.DTO{Login: login, Email: email, Password: password, Role: domain.Role(role), Color: color, Username: username})

		if err != nil {
			c.Set(code.Error, strx.Capitalize(err.Error()))
			ShowSignup(c)
			return
		}

		c.Redirect(http.StatusFound, "/login")
	}
}
