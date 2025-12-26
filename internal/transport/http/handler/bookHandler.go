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
		{ID: 101, HideID: true, Name: "Практика"},
		{ID: 1, Name: "Общие вопросы"},
		{ID: 2, Name: "Срезы (Slices)"},
		{ID: 3, Name: "Массивы (Array)"},
		{ID: 4, Name: "Карты (Maps)"},
		{ID: 5, Name: "Функции и методы (Functions and methods)"},
		{ID: 6, Name: "Интерфейсы (Interfaces)"},
		{ID: 7, Name: "Горутины (Goroutines)"},
		{ID: 8, Name: "Планировщик (Scheduler)"},
		{ID: 9, Name: "Каналы (Channels)"},
		{ID: 10, Name: "Синхронизация (Sync)"},
		{ID: 11, Name: "Ошибки и паники (Error and panics)"},
		{ID: 12, Name: "Отложенный вызов (Defer)"},
		{ID: 13, Name: "Стек и куча (Stack and heap)"},
		{ID: 14, Name: "Указатели (Pointers)"},
		{ID: 15, Name: "Сборщик мусора (Garbage collector)"},
		{ID: 16, Name: "Контекст (Context)"},
		{ID: 17, Name: "Строки (Strings)"},
		{ID: 18, Name: "Обобщения (Generics)"},
		{ID: 19, Name: "Мультиплексор событий (Select)"},
		{ID: 20, Name: "Тестирование (Testing)"},
		{ID: 21, Name: "Методология CI/CD"},
		{ID: 22, Name: "Модель памяти и гонки данных (Memory model & Data races)"},
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

	data := tplWithMetaData(c, "Книга по Golang")
	data["chapter_id"] = chapterID
	data["chapters"] = bookChapters
	c.HTML(http.StatusOK, "book.html", data)
}
