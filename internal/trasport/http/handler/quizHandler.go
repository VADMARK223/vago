package handler

import (
	"net/http"
	"vago/internal/app"
	"vago/internal/application/quiz"
	"vago/internal/application/topic"
	"vago/internal/config/code"

	"github.com/gin-gonic/gin"
)

type QuizHandler struct {
	quizSvc  *quiz.Service
	topicSvc *topic.Service
}

func NewQuizHandler(quizSvc *quiz.Service, topicSvc *topic.Service) *QuizHandler {
	return &QuizHandler{quizSvc: quizSvc, topicSvc: topicSvc}
}

func (h *QuizHandler) ShowQuizAdmin() func(c *gin.Context) {
	return func(c *gin.Context) {
		topics, _ := h.topicSvc.All()

		app.Dump("Topics", topics)

		questions, _ := h.quizSvc.AllQuestions()

		data := tplWithCapture(c, "Админка викторины")
		data[code.Topics] = topics
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
