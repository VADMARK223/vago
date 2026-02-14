package handler

import (
	"net/http"
	"strconv"
	"vago/internal/application/user"
	"vago/internal/config/code"
	"vago/internal/domain"
	"vago/internal/transport/http/api/response"
	"vago/pkg/strx"

	"github.com/gin-gonic/gin"
)

func DeleteUser(service *user.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		parseId, parseIdErr := strconv.ParseInt(c.Param("id"), 10, 64)
		if parseIdErr != nil {
			response.Error(c, http.StatusBadRequest, "Некорректные данные")
			return
		}

		currentId, errGetUSerId := c.Get(code.UserId)
		if !errGetUSerId {
			response.Error(c, http.StatusUnauthorized, "Пользователь не аутентифицировался")
			return
		}

		if currentId == parseId {
			response.Error(c, http.StatusBadRequest, "Вы пытаетесь удалить свой собственный аккаунт")
			return
		}

		role, errRole := c.Get(code.Role)
		if !errRole {
			response.Error(c, http.StatusBadRequest, "Роль пользователя неизвестна")
			return
		}

		if role != string(domain.RoleAdmin) {
			// Пользователь аутентифицирован, но не авторизован для этого действия
			response.Error(c, http.StatusForbidden, "У вас нет прав на удаление пользователей")
			return
		}

		err := service.DeleteUser(parseId)
		if err != nil {
			response.Error(c, mapErrorToHTTP(err), strx.Capitalize(err.Error()))
			return
		}

		response.OKNoData(c, "Пользователь удален")
	}
}
