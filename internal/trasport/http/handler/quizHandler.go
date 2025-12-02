package handler

import (
	"net/http"
	"vago/internal/application/quiz"
	"vago/internal/config/code"

	"github.com/gin-gonic/gin"
)

type QuizHandler struct {
	quizSvc *quiz.Service
}

func NewQuizHandler(quizSvc *quiz.Service) *QuizHandler {
	return &QuizHandler{quizSvc: quizSvc}
}

func (h *QuizHandler) ShowQuizAdmin() func(c *gin.Context) {
	return func(c *gin.Context) {
		questions, _ := h.quizSvc.AllQuestions()

		data := tplWithCapture(c, "Админка викторины")
		data[code.Questions] = questions
		data[code.QuestionsCount] = len(questions)

		c.HTML(http.StatusOK, "quiz.html", data)
	}
}

func (h *QuizHandler) DeleteAllQuestions() func(c *gin.Context) {
	return func(c *gin.Context) {
		err := h.quizSvc.DeleteAll()
		if err != nil {
			ShowError(c, "Ошибка удаления всех вопросов", err.Error())
		}
		c.Redirect(http.StatusSeeOther, "/quiz")
	}
}
