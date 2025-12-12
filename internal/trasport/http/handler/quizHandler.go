package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"vago/internal/application/quiz"
	"vago/internal/application/topic"
	"vago/internal/config/code"
	"vago/internal/seed"

	"github.com/gin-gonic/gin"
)

type QuizHandler struct {
	quizSvc  *quiz.Service
	topicSvc *topic.Service
	dsn      string
}

type CheckRequest struct {
	QuestionID uint `json:"question_id"`
	AnswerID   uint `json:"answer_id"`
}

type CheckResponse struct {
	Correct     bool   `json:"correct"`
	Explanation string `json:"explanation"`
}

func NewQuizHandler(quizSvc *quiz.Service, topicSvc *topic.Service, dsn string) *QuizHandler {
	return &QuizHandler{quizSvc: quizSvc, topicSvc: topicSvc, dsn: dsn}
}

func (h *QuizHandler) ShowQuizRandom() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := h.quizSvc.RandomID()
		if err != nil {
			ShowError(c, "Ошибка генерации случайного вопроса", err.Error())
			return
		}
		c.Redirect(http.StatusFound, fmt.Sprintf("/quiz/%d", id))
	}
}

func (h *QuizHandler) ShowQuizByID() func(c *gin.Context) {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id64, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.String(400, "invalid id")
			return
		}
		temp := uint(id64)
		q := h.quizSvc.RandomPublicQuestion(&temp)

		renderQuiz(c, q)
	}
}

func renderQuiz(c *gin.Context, q quiz.QuestionPublic) {
	data := tplWithCapture(c, "Викторина")
	data[code.Question] = q

	c.HTML(http.StatusOK, "quiz.html", data)
}

func (h *QuizHandler) Check() func(c *gin.Context) {
	return func(c *gin.Context) {
		var req CheckRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid"})
			return
		}

		correct, explanation := h.quizSvc.CheckAnswer(req.QuestionID, req.AnswerID)
		c.JSON(200, CheckResponse{Correct: correct, Explanation: explanation})
	}
}

func (h *QuizHandler) ShowQuizAdmin() func(c *gin.Context) {
	return func(c *gin.Context) {
		topics, _ := h.topicSvc.AllWithCount()

		questions, _ := h.quizSvc.AllQuestions()

		data := tplWithCapture(c, "Админка викторины")
		data[code.Topics] = topics
		data[code.Questions] = questions
		data[code.QuestionsCount] = len(questions)

		c.HTML(http.StatusOK, "questions.html", data)
	}
}

func (h *QuizHandler) AddQuestion() func(c *gin.Context) {
	return func(c *gin.Context) {
		text := c.PostForm("text")
		codeStr := c.PostForm("code")
		answer1 := c.PostForm("answer1")
		answer2 := c.PostForm("answer2")
		answer3 := c.PostForm("answer3")
		answer4 := c.PostForm("answer4")
		correctAnswerIdxStr := c.PostForm("correct_answer_index")
		topicIdStr := c.PostForm("topic_id")
		explanation := c.PostForm("explanation")

		topicId, _ := strconv.Atoi(topicIdStr)
		correctIdx, _ := strconv.Atoi(correctAnswerIdxStr)

		answers := []seed.Answer{
			{Text: answer1},
			{Text: answer2},
			{Text: answer3},
			{Text: answer4},
		}

		if correctIdx >= 0 && correctIdx < len(answers) {
			answers[correctIdx].Correct = true
		}

		err := seed.AddQuestion(seed.Question{
			TopicID:     topicId,
			Text:        text,
			Code:        codeStr,
			Explanation: explanation,
			Answers:     answers,
		})

		if err != nil {
			ShowError(c, "Ошибка добавления вопроса", err.Error())
			return
		}

		c.Redirect(http.StatusSeeOther, "/questions")
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

func (h *QuizHandler) RunQuestionsSeed() func(c *gin.Context) {
	return func(c *gin.Context) {
		err := seed.SyncQuestions(h.dsn)
		if err != nil {
			ShowError(c, "Ошибка сидирования", err.Error())
			return
		}
		c.Redirect(http.StatusSeeOther, "/questions")
	}
}

func (h *QuizHandler) RunTopicsSeed() func(c *gin.Context) {
	return func(c *gin.Context) {
		err := seed.SyncTopics(h.dsn)
		if err != nil {
			ShowError(c, "Ошибка сидирования", err.Error())
			return
		}
		c.Redirect(http.StatusSeeOther, "/questions")
	}
}
