package handler

import (
	"net/http"
	"vago/internal/domain"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	domain.ClearTokenCookies(c)
	c.Redirect(http.StatusFound, "/")
}
