package question

import (
	"strconv"
	"vago/internal/application/chapter"
	"vago/internal/application/test"
	"vago/internal/application/topic"
	"vago/internal/domain"

	"github.com/gin-gonic/gin"
)

type Loader struct {
	ChapterSvc *chapter.Service
	TopicSvc   *topic.Service
	TestSvc    *test.Service
}

type Data struct {
	TopicID   int64
	Chapters  []*domain.Chapter
	Questions []*domain.Question
	Topics    []domain.TopicWithCount
}

func (l Loader) Load(c *gin.Context) (Data, error) {
	var out Data

	chapters, err := l.ChapterSvc.All()
	if err != nil {
		return out, err
	}
	out.Chapters = chapters

	topics, err := l.TopicSvc.AllWithCount()
	if err != nil {
		return out, err
	}
	out.Topics = topics

	topicIDStr := c.Query("topic_id")
	if topicIDStr != "" {
		id, err := strconv.ParseInt(topicIDStr, 10, 64)
		if err != nil {
			return out, err
		}
		out.TopicID = id

		out.Questions, err = l.TestSvc.GetQuestionsByTopic(id)
		if err != nil {
			return out, err
		}
		return out, nil
	}

	out.Questions, err = l.TestSvc.AllQuestions()
	if err != nil {
		return out, err
	}
	return out, nil
}
