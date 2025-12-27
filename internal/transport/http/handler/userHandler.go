package handler

import (
	"strconv"
	"vago/internal/application/user"
	"vago/internal/config/code"

	"github.com/gin-gonic/gin"
)

func DeleteUser(service *user.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		parseUint, parseUintErr := strconv.ParseInt(c.Param("id"), 10, 32)
		if parseUintErr != nil {
			return
		}

		currentId, errGetUSerId := c.Get(code.UserId)
		if !errGetUSerId {
			c.JSON(500, gin.H{
				"error": "Not fount user id",
			})
			return
		}

		if currentId == uint(parseUint) {
			c.JSON(400, gin.H{
				"error": "Вы пытаетесь удалить свой собственный аккаунт",
			})
			return
		}

		role, errRole := c.Get(code.Role)
		if !errRole {
			c.JSON(500, gin.H{
				"error": "Not fount user role",
			})
			return
		}

		if role != "admin" {
			c.JSON(400, gin.H{
				"error": "У вас нет прав на удаление аккаунтов",
			})
			return
		}

		err := service.DeleteUser(parseUint)
		if err != nil {
			ShowError(c, "Ошибка удаления пользователя", err.Error())
			return
		}

		c.JSON(200, gin.H{})
	}
}
