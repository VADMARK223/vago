package message

import (
	"net/http"
	"strconv"
	"vago/internal/application/message"
	"vago/internal/config/code"
	"vago/internal/domain"
	"vago/internal/transport/http/api/response"

	//"vago/internal/transport/http/handler"

	"github.com/gin-gonic/gin"

	shared "vago/internal/transport/http/shared/message"
)

type Handler struct {
	Loader     shared.Loader
	messageSvc *message.Service
}

func New(loader shared.Loader, messageSvc *message.Service) *Handler {
	return &Handler{
		Loader:     loader,
		messageSvc: messageSvc,
	}
}

func (h *Handler) GetAllWithUsername(c *gin.Context) {
	messages, err := h.Loader.Load(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Список сообщений")
		return
	}

	response.OK(c, "Список сообщений", MessagesApiDTO{Messages: MessagesToDTO(messages)})
}

func (h *Handler) Delete(c *gin.Context) {
	role, errRole := c.Get(code.Role)
	if !errRole {
		response.Error(c, http.StatusBadRequest, "Роль пользователя неизвестна")
		return
	}

	if role != string(domain.RoleAdmin) {
		response.Error(c, http.StatusForbidden, "У вас нет прав на удаление сообщения")
		return
	}

	parseUint, parseUintErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseUintErr != nil {
		response.Error(c, http.StatusInternalServerError, "Ошибка конвертации идентификатора")
		return
	}

	errDelete := h.messageSvc.DeleteMessage(parseUint)

	if errDelete != nil {
		response.Error(c, http.StatusInternalServerError, `Ошибка удаления сообщения`)
		return
	}

	response.OKNoData(c, "Сообщение удалено")
}

func (h *Handler) DeleteAll(c *gin.Context) {
	role, errRole := c.Get(code.Role)
	if !errRole {
		response.Error(c, http.StatusInternalServerError, "Ошибка удаления всех сообщений")

		//handler.ShowError(c, "Ошибка удаления всех сообщений", "Роль пользователя неизвестна")
		return
	}

	if role != string(domain.RoleAdmin) {
		response.Error(c, http.StatusInternalServerError, "У вас нет прав на удаление сообщений")
		//handler.ShowError(c, "Ошибка удаления всех сообщений", "У вас нет прав на удаление пользователей")
		return
	}

	err := h.messageSvc.DeleteAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Ошибка удаления всех сообщений")
		//handler.ShowError(c, "Ошибка удаления всех сообщений", err.Error())
	}
	response.OKNoData(c, "Все сообщения удалены")
}
