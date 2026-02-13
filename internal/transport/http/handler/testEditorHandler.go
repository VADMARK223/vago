package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"vago/internal/application/test"
	"vago/internal/application/topic"
	"vago/internal/config/code"
	"vago/internal/domain"
	"vago/internal/seed"
	"vago/internal/transport/http/api/response"

	"github.com/gin-gonic/gin"
)

type TestEditorHandler struct {
	testSvc  *test.Service
	topicSvc *topic.Service
	dsn      string
}

func NewTestEditorHandler(
	testSvc *test.Service,
	topicSvc *topic.Service,
	dsn string,
) *TestEditorHandler {
	return &TestEditorHandler{testSvc: testSvc, topicSvc: topicSvc, dsn: dsn}
}

func (h *TestEditorHandler) ShowAddQuestion(c *gin.Context) {
	topics, _ := h.topicSvc.AllWithCount()
	questions, _ := h.testSvc.AllQuestions()

	data := TplWithMetaData(c, "Добавление вопроса")
	data[code.Topics] = topics
	data[code.QuestionsCount] = len(questions)

	c.HTML(http.StatusOK, "add_question.html", data)
}

func (h *TestEditorHandler) AddQuestion(c *gin.Context) {
	role, ok := c.Get(code.Role)
	r, ok := role.(string)
	if !ok {
		response.Error(c, http.StatusBadRequest, "Роль пользователя неизвестна")
		return
	}

	if r != string(domain.RoleAdmin) {
		response.Error(c, http.StatusForbidden, "У вас нет прав на это действие")
		return
	}

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

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	err := seed.AddQuestion(ctx, h.dsn, seed.Question{
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

	c.Redirect(http.StatusSeeOther, "/add_questions")
}

func (h *TestEditorHandler) RunGoTopicsSeed(c *gin.Context) {
	role, ok := c.Get(code.Role)
	r, ok := role.(string)
	if !ok {
		response.Error(c, http.StatusBadRequest, "Роль пользователя неизвестна")
		return
	}

	if r != string(domain.RoleAdmin) {
		response.Error(c, http.StatusForbidden, "У вас нет прав на это действие")
		return
	}

	err := seed.GoTopics(h.dsn)
	if err != nil {
		ShowError(c, "Ошибка сидирования", err.Error())
		return
	}
	c.Redirect(http.StatusSeeOther, "/add_questions")
}

func (h *TestEditorHandler) RunGoQuestionsSeed(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	role, ok := c.Get(code.Role)
	r, ok := role.(string)
	if !ok {
		response.Error(c, http.StatusBadRequest, "Роль пользователя неизвестна")
		return
	}

	if r != string(domain.RoleAdmin) {
		response.Error(c, http.StatusForbidden, "У вас нет прав на это действие")
		return
	}

	if err := seed.SyncQuestions(ctx, h.dsn); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
