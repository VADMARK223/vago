package message

import (
	"net/http"
	"strconv"
	"vago/internal/application/message"
	"vago/internal/config/code"
	"vago/internal/config/route"
	"vago/internal/domain"
	"vago/internal/transport/http/api/response"
	"vago/internal/transport/http/handler"
	"vago/internal/transport/http/shared/template"

	shared "vago/internal/transport/http/shared/message"

	"github.com/gin-gonic/gin"
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

func (h *Handler) Page(c *gin.Context) {
	all, err := h.Loader.Load(c)
	if err != nil {
		handler.ShowError(c, "Ошибка получения списка сообщений", err.Error())
		return
	}

	data := template.BaseAdminData(c, "Сообщения")
	data[code.Messages] = all
	data[code.MessagesCount] = len(all)
	data["Active"] = "messages"
	c.HTML(http.StatusOK, "admin/layout", data)
}

func (h *Handler) Delete(c *gin.Context) {
	role, errRole := c.Get(code.Role)
	if !errRole {
		response.Error(c, http.StatusBadRequest, "Роль пользователя неизвестна")
		return
	}

	if role != string(domain.RoleAdmin) {
		response.Error(c, http.StatusForbidden, "У вас нет прав на удаление пользователей")
		return
	}

	parseUint, parseUintErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseUintErr != nil {
		handler.ShowError(c, "Ошибка конвертации идентификатора", parseUintErr.Error())
		return
	}

	errDelete := h.messageSvc.DeleteMessage(parseUint)

	if errDelete != nil {
		handler.ShowError(c, "Ошибка удаления сообщения", errDelete.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, route.Admin+route.Messages)
}

func (h *Handler) DeleteAll(c *gin.Context) {
	role, errRole := c.Get(code.Role)
	if !errRole {
		handler.ShowError(c, "Ошибка удаления всех сообщений", "Роль пользователя неизвестна")
		return
	}

	if role != string(domain.RoleAdmin) {
		handler.ShowError(c, "Ошибка удаления всех сообщений", "У вас нет прав на удаление пользователей")
		return
	}

	err := h.messageSvc.DeleteAll()
	if err != nil {
		handler.ShowError(c, "Ошибка удаления всех сообщений", err.Error())
	}
	c.Redirect(http.StatusSeeOther, route.Admin+route.Messages)
}
