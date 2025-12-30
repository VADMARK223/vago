package handler

import (
	"net/http"
	"strconv"
	"vago/internal/application/comment"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentSvc *comment.Service
}

func NewCommentHandler(commentSvc *comment.Service) *CommentHandler {
	return &CommentHandler{commentSvc: commentSvc}
}

func (h *CommentHandler) PostComment(c *gin.Context) {
	ctx := c.Request.Context()

	questionID, err := strconv.ParseInt(c.PostForm("question_id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var parentID *int64
	if v := c.PostForm("parent_id"); v != "" {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		parentID = &id
	}

	userID := c.MustGet("user_id").(int64) // или как ты достаёшь юзера

	dto := comment.CreateCommentDTO{
		QuestionID: questionID,
		ParentID:   parentID,
		AuthorID:   userID,
		Content:    c.PostForm("content"),
	}

	_, err = h.commentSvc.Create(ctx, dto)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Redirect(http.StatusSeeOther, "/test/"+strconv.FormatInt(questionID, 10))
}
