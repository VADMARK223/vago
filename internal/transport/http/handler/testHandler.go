package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"vago/internal/application/comment"
	"vago/internal/application/test"
	"vago/internal/application/topic"
	"vago/internal/config/code"
	"vago/internal/config/route"
	"vago/internal/transport/http/api/response"
	"vago/internal/transport/http/dto"
	"vago/internal/transport/http/shared/template"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	testSvc    *test.Service
	topicSvc   *topic.Service
	commentSvc *comment.Service
}

type CheckRequest struct {
	QuestionID int64 `json:"question_id"`
	AnswerID   int64 `json:"answer_id"`
}

type CheckRequestAPI struct {
	QuestionID int64 `json:"questionId"`
	AnswerID   int64 `json:"answerId"`
}

type CheckResponse struct {
	Correct     bool   `json:"correct"`
	Explanation string `json:"explanation,omitempty"`
}

func NewTestHandler(
	testSvc *test.Service,
	topicSvc *topic.Service,
	commentSvc *comment.Service,
) *TestHandler {
	return &TestHandler{testSvc: testSvc, topicSvc: topicSvc, commentSvc: commentSvc}
}

func (h *TestHandler) ShowRandom(c *gin.Context) {
	id, err := h.testSvc.RandomID()
	if err != nil {
		ShowError(c, "Ошибка генерации случайного вопроса", err.Error())
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf(route.Test+"/%d", id))
}

func (h *TestHandler) RandomQuestionIdAPI(c *gin.Context) {
	id, err := h.testSvc.RandomID()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Ошибка генерации случайного вопроса"+err.Error())
		return
	}

	response.OK(c, fmt.Sprintf("Идентификатор вопроса: %d", id), id)
}

func (h *TestHandler) QuestionByIdAPI(c *gin.Context) {
	questionId, parseIdErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseIdErr != nil {
		response.Error(c, http.StatusBadRequest, "Некорректные данные")
		return
	}

	q := h.testSvc.PublicQuestion(&questionId)

	response.OK(c, fmt.Sprintf("Вопрос: %d", questionId), dto.QuestionPublicToDTO(q))
}

func (h *TestHandler) ShowByID(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.String(400, "invalid id")
		return
	}

	q := h.testSvc.PublicQuestion(&id64)

	renderTestPage(c, h, q)
}

func (h *TestHandler) CheckAnswerAPI(c *gin.Context) {
	var req CheckRequestAPI
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	correct, explanation, err := h.testSvc.CheckAnswer(req.QuestionID, req.AnswerID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if correct {
		response.OK(c, "Правильный ответ", CheckResponse{Correct: correct, Explanation: explanation})
	} else {
		response.OK(c, "Неправильный ответ", CheckResponse{Correct: correct})
	}
}

func (h *TestHandler) CheckAnswer(c *gin.Context) {
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

func renderTestPage(c *gin.Context, h *TestHandler, q test.QuestionPublic) {
	comments, count, err := h.commentSvc.ListByQuestionID(q.ID)
	if err != nil {
		ShowError(c, "Ошибка загрузки комментариев", err.Error())
		return
	}

	data := template.TplWithMetaData(c, "Тест")
	data[code.Question] = q
	data[code.CommentsCount] = count
	data[code.Comments] = comments

	c.HTML(http.StatusOK, "test.html", data)
}
