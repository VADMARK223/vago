package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	chapterIDStr := c.Query("chapter_id")

	var (
		chapterID uint64
		err       error
	)

	if chapterIDStr != "" {
		chapterID, err = strconv.ParseUint(chapterIDStr, 10, 64)
		if err != nil {
			ShowError(c, "Ошибка", err.Error())
			return
		}
	}

	data := tplWithCapture(c, "Книга по Golang")
	data["chapter_id"] = chapterID
	c.HTML(http.StatusOK, "book.html", data)
}
