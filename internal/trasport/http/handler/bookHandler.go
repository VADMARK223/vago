package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookChapter struct {
	ID     int64
	Name   string
	HideID bool
}

func ShowBook(c *gin.Context) {
	chapterIDStr := c.Query("chapter_id")

	bookChapters := []bookChapter{
		{ID: 100, HideID: true, Name: "Шпаргалка"},
		{ID: 1, Name: "Общие вопросы"},
		{ID: 2, Name: "Срезы (Slices)"},
		{ID: 4, Name: "Карты (Maps)"},
		{ID: 9, Name: "Каналы"},
		{ID: 11, Name: "Ошибка (Errors)"},
		{ID: 12, Name: "Defer"},
		{ID: 14, Name: "Указатели (Pointers)"},
		{ID: 16, Name: "Контекст (Context)"},
		{ID: 17, Name: "Строки (Strings)"},
	}

	bookTasks := []bookChapter{
		{ID: -1, Name: "Задача 1"},
		{ID: -2, Name: "Fan-out / Fan-in"},
	}

	var (
		chapterID int64
		err       error
	)

	if chapterIDStr != "" {
		chapterID, err = strconv.ParseInt(chapterIDStr, 10, 64)
		if err != nil {
			ShowError(c, "Ошибка", err.Error())
			return
		}
	}

	data := tplWithCapture(c, "Книга по Golang")
	data["chapter_id"] = chapterID
	data["chapters"] = bookChapters
	data["tasks"] = bookTasks
	c.HTML(http.StatusOK, "book.html", data)
}
