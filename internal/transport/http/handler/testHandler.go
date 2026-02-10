package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"vago/internal/application/comment"
	"vago/internal/application/test"
	"vago/internal/application/topic"
	"vago/internal/config/code"
	"vago/internal/domain"
	"vago/internal/seed"
	"vago/internal/transport/http/api"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	testSvc    *test.Service
	topicSvc   *topic.Service
	commentSvc *comment.Service
	dsn        string
}

type CheckRequest struct {
	QuestionID int64 `json:"question_id"`
	AnswerID   int64 `json:"answer_id"`
}

type CheckResponse struct {
	Correct     bool   `json:"correct"`
	Explanation string `json:"explanation"`
}

func NewTestHandler(testSvc *test.Service, topicSvc *topic.Service, commentSvc *comment.Service, dsn string) *TestHandler {
	return &TestHandler{testSvc: testSvc, topicSvc: topicSvc, commentSvc: commentSvc, dsn: dsn}
}

func (h *TestHandler) ShowTestRandom() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := h.testSvc.RandomID()
		if err != nil {
			ShowError(c, "Ошибка генерации случайного вопроса", err.Error())
			return
		}
		//id := 1
		c.Redirect(http.StatusFound, fmt.Sprintf("/test/%d", id))
	}
}

func (h *TestHandler) ShowTestByID() func(c *gin.Context) {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id64, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.String(400, "invalid id")
			return
		}

		q := h.testSvc.RandomPublicQuestion(&id64)

		renderTestPage(c, h, q)
	}
}

func renderTestPage(c *gin.Context, h *TestHandler, q test.QuestionPublic) {
	comments, count, err := h.commentSvc.ListByQuestionID(q.ID)
	if err != nil {
		ShowError(c, "Ошибка загрузки комментариев", err.Error())
		return
	}

	data := tplWithMetaData(c, "Тест")
	data[code.Question] = q
	data[code.CommentsCount] = count
	data[code.Comments] = comments

	c.HTML(http.StatusOK, "test.html", data)
}

func (h *TestHandler) Check() func(c *gin.Context) {
	return func(c *gin.Context) {
		var req CheckRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid"})
			return
		}

		correct, explanation, err := h.testSvc.CheckAnswer(req.QuestionID, req.AnswerID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, CheckResponse{Correct: correct, Explanation: explanation})
	}
}

func (h *TestHandler) ShowAddQuestion() func(c *gin.Context) {
	return func(c *gin.Context) {
		topics, _ := h.topicSvc.AllWithCount()

		questions, _ := h.testSvc.AllQuestions()

		data := tplWithMetaData(c, "Добавление вопроса")
		data[code.Topics] = topics
		data[code.QuestionsCount] = len(questions)

		c.HTML(http.StatusOK, "add_question.html", data)
	}
}

func (h *TestHandler) ShowQuestions(c *gin.Context) {
	topicIDStr := c.Query("topic_id")

	topics, _ := h.topicSvc.AllWithCount()

	var (
		topicID   int64
		questions []*domain.Question
		err       error
	)

	if topicIDStr != "" {
		topicID, err = strconv.ParseInt(topicIDStr, 10, 64)
		if err != nil {
			ShowError(c, "Ошибка", err.Error())
			return
		}
		questions, err = h.testSvc.GetQuestionsByTopic(topicID)
	} else {
		questions, err = h.testSvc.AllQuestions()
	}

	if err != nil {
		ShowError(c, "Ошибка выборки", err.Error())
		return
	}

	data := tplWithMetaData(c, "Редактор вопросов")
	data[code.Topics] = topics
	data["topic_id"] = topicID
	data[code.Questions] = questions
	data[code.QuestionsCount] = len(questions)

	c.HTML(http.StatusOK, "questions.html", data)
}

func (h *TestHandler) AddQuestion() func(c *gin.Context) {
	return func(c *gin.Context) {
		role, errRole := c.Get(code.Role)
		if !errRole {
			api.Error(c, http.StatusBadRequest, "Роль пользователя неизвестна")
			return
		}

		if role != domain.RoleAdmin {
			api.Error(c, http.StatusForbidden, "У вас нет прав на это действие")
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
}

func (h *TestHandler) RunGoTopicsSeed() func(c *gin.Context) {
	return func(c *gin.Context) {
		role, errRole := c.Get(code.Role)
		if !errRole {
			api.Error(c, http.StatusBadRequest, "Роль пользователя неизвестна")
			return
		}

		if role != domain.RoleAdmin {
			api.Error(c, http.StatusForbidden, "У вас нет прав на это действие")
			return
		}

		err := seed.GoTopics(h.dsn)
		if err != nil {
			ShowError(c, "Ошибка сидирования", err.Error())
			return
		}
		c.Redirect(http.StatusSeeOther, "/add_questions")
	}
}

func (h *TestHandler) RunGoQuestionsSeed() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
		defer cancel()

		role, errRole := c.Get(code.Role)
		if !errRole {
			api.Error(c, http.StatusBadRequest, "Роль пользователя неизвестна")
			return
		}

		if role != domain.RoleAdmin {
			api.Error(c, http.StatusForbidden, "У вас нет прав на это действие")
			return
		}

		if err := seed.SyncQuestions(ctx, h.dsn); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
