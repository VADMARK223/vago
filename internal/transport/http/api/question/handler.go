package question

import (
	"net/http"
	"vago/internal/transport/http/api/response"
	"vago/pkg/strx"

	shared "vago/internal/transport/http/shared/question"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Loader shared.Loader
}

func New(loader shared.Loader) *Handler {
	return &Handler{Loader: loader}
}

func (h *Handler) Get(c *gin.Context) {
	data, err := h.Loader.Load(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, strx.Capitalize(err.Error()))
		return
	}

	response.OK(c, "Вопросы", ToDTO(data.Chapters, data.Topics, data.Questions))
}
