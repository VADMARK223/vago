package handler

import (
	"net/http"
	"strconv"
	"vago/internal/application/user"
	"vago/internal/config/code"
	"vago/internal/domain"
	"vago/internal/transport/http/api"
	"vago/pkg/strx"

	"github.com/gin-gonic/gin"
)

func DeleteUser(service *user.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		parseId, parseIdErr := strconv.ParseInt(c.Param("id"), 10, 64)
		if parseIdErr != nil {
			return
		}

		currentId, errGetUSerId := c.Get(code.UserId)
		if !errGetUSerId {
			api.Error(c, http.StatusUnauthorized, "Пользователь не аутентифицировался")
			return
		}

		if currentId == parseId {
			api.Error(c, http.StatusBadRequest, "Вы пытаетесь удалить свой собственный аккаунт")
			return
		}

		role, errRole := c.Get(code.Role)
		if !errRole {
			api.Error(c, http.StatusBadRequest, "Роль пользователя неизвестна")
			return
		}

		if role != domain.RoleAdmin {
			// Пользователь аутентифицирован, но не авторизован для этого действия
			api.Error(c, http.StatusForbidden, "У вас нет прав на удаление пользователей")
			return
		}

		err := service.DeleteUser(parseId)
		if err != nil {
			api.Error(c, mapErrorToHTTP(err), strx.Capitalize(err.Error()))
			return
		}

		api.OKNoData(c, "Пользователь удален")
	}
}
