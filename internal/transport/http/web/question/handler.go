package question

import (
	"net/http"
	"vago/internal/config/code"
	"vago/internal/transport/http/handler"
	shared "vago/internal/transport/http/shared/question"
	"vago/internal/transport/http/shared/template"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Loader shared.Loader
}

func New(loader shared.Loader) *Handler {
	return &Handler{Loader: loader}
}

func (h *Handler) Page(c *gin.Context) {
	data, err := h.Loader.Load(c)
	if err != nil {
		handler.ShowError(c, "Ошибка выборки", err.Error())
		return
	}

	vm := ToViewModel(data)

	tpl := template.TplWithMetaData(c, "Редактор вопросов") // твоя функция
	tpl[code.Topics] = vm.Topics
	tpl["topic_id"] = vm.TopicID
	tpl[code.Questions] = vm.Questions

	if qs, ok := data.Questions, true; ok {
		tpl[code.QuestionsCount] = len(qs)
	}

	c.HTML(http.StatusOK, "questions.html", tpl)
}
