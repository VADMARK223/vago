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

type CheckRequest struct {
	QuestionID uint `json:"question_id"`
	AnswerID   uint `json:"answer_id"`
}

type CheckResponse struct {
	Correct bool `json:"correct"`
}

func NewQuizHandler(quizSvc *quiz.Service, topicSvc *topic.Service) *QuizHandler {
	return &QuizHandler{quizSvc: quizSvc, topicSvc: topicSvc}
}

func (h *QuizHandler) ShowQuiz() func(c *gin.Context) {
	return func(c *gin.Context) {
		q, _ := h.quizSvc.RandomQuestion()
		public := h.quizSvc.ToPublic(q)

		data := tplWithCapture(c, "Викторина")
		data[code.Question] = public

		c.HTML(http.StatusOK, "quiz.html", data)
	}
}

func (h *QuizHandler) Check() func(c *gin.Context) {
	return func(c *gin.Context) {
		var req CheckRequest
		app.Dump("asd", req)
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid"})
			return
		}

		correct := h.quizSvc.CheckAnswer(req.QuestionID, req.AnswerID)
		c.JSON(200, CheckResponse{Correct: correct})
	}
}

func (h *QuizHandler) ShowQuizAdmin() func(c *gin.Context) {
	return func(c *gin.Context) {
		topics, _ := h.topicSvc.AllWithCount()

		app.Dump("Topics", topics)

		questions, _ := h.quizSvc.AllQuestions()

		data := tplWithCapture(c, "Админка викторины")
		data[code.Topics] = topics
		data[code.Questions] = questions
		data[code.QuestionsCount] = len(questions)

		c.HTML(http.StatusOK, "questions.html", data)
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
