package handler

import (
	"net/http"
	"strconv"
	"vago/internal/application/chapter"
	"vago/internal/application/test"
	"vago/internal/application/topic"
	"vago/internal/config/code"
	"vago/internal/domain"
	"vago/internal/transport/http/api/question"
	"vago/internal/transport/http/api/response"
	"vago/pkg/strx"

	"github.com/gin-gonic/gin"
)

// QuestionHandler страница глав, топиков и вопросов
type QuestionHandler struct {
	chapterSvc *chapter.Service
	topicSvc   *topic.Service
	testSvc    *test.Service
}

func NewQuestionHandler(
	chapterSvc *chapter.Service,
	topicSvc *topic.Service,
	testSvc *test.Service,
) *QuestionHandler {
	return &QuestionHandler{
		chapterSvc: chapterSvc,
		topicSvc:   topicSvc,
		testSvc:    testSvc,
	}
}

func (h *QuestionHandler) ShowQuestionsAPI(c *gin.Context) {
	var (
		topicID   int64
		questions []*domain.Question
		err       error
	)

	chapters, errChapters := h.chapterSvc.All()
	if errChapters != nil {
		response.Error(c, http.StatusInternalServerError, strx.Capitalize(errChapters.Error()))
		return
	}

	topicIDStr := c.Query("topic_id")
	topics, _ := h.topicSvc.AllWithCount()

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

	response.OK(c, "Вопросы", question.ToDTO(chapters, topics, questions))
}

func (h *QuestionHandler) ShowQuestions(c *gin.Context) {
	var (
		topicID   int64
		questions []*domain.Question
		err       error
	)

	topicIDStr := c.Query("topic_id")
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

	topics, _ := h.topicSvc.AllWithCount()

	data := tplWithMetaData(c, "Редактор вопросов")
	data[code.Topics] = topics
	data["topic_id"] = topicID
	data[code.Questions] = questions
	data[code.QuestionsCount] = len(questions)

	c.HTML(http.StatusOK, "questions.html", data)
}
