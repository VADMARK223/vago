package question

import "vago/internal/domain"

type ChapterDTO struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Order int64  `json:"order"`
}

type TopicDTO struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	QuestionsCount int64  `json:"questionsCount"`
}

type QuestionDTO struct {
	ID          int64       `json:"id"`
	TopicID     int64       `json:"topicId"`
	Text        string      `json:"text"`
	Code        string      `json:"code"`
	Explanation string      `json:"explanation"`
	Answers     []AnswerDTO `json:"answers"`
}

type AnswerDTO struct {
	ID        int64  `json:"id"`
	Text      string `json:"text"`
	IsCorrect bool   `json:"isCorrect"`
}

type QuestionsPageDTO struct {
	Chapters  []ChapterDTO  `json:"chapters"`
	Topics    []TopicDTO    `json:"topics"`
	Questions []QuestionDTO `json:"questions"`
}

func ToDTO(chapters []*domain.Chapter, topics []domain.TopicWithCount, questions []*domain.Question) QuestionsPageDTO {
	return QuestionsPageDTO{Chapters: chaptersToDTO(chapters), Topics: topicsToDTO(topics), Questions: questionsToDTO(questions)}
}

func chaptersToDTO(chapters []*domain.Chapter) []ChapterDTO {
	result := make([]ChapterDTO, 0, len(chapters))
	for _, t := range chapters {
		result = append(result, chapterToDTO(t))
	}

	return result
}

func chapterToDTO(c *domain.Chapter) ChapterDTO {
	return ChapterDTO{
		ID:    c.ID,
		Name:  c.Name,
		Order: c.Order,
	}
}

func topicsToDTO(topics []domain.TopicWithCount) []TopicDTO {
	result := make([]TopicDTO, 0, len(topics))
	for _, t := range topics {
		result = append(result, topicToDTO(t))
	}

	return result
}

func topicToDTO(t domain.TopicWithCount) TopicDTO {
	return TopicDTO{
		ID:             t.ID,
		Name:           t.Name,
		QuestionsCount: t.QuestionsCount,
	}
}

func questionsToDTO(questions []*domain.Question) []QuestionDTO {
	result := make([]QuestionDTO, 0, len(questions))
	for _, q := range questions {
		result = append(result, questionToDTO(q))
	}

	return result
}

func questionToDTO(q *domain.Question) QuestionDTO {
	result := QuestionDTO{
		ID:          q.ID,
		Text:        q.Text,
		Code:        q.Code,
		Explanation: q.Explanation,
		TopicID:     q.TopicID,
	}
	result.Answers = answersToDTO(q.Answers)
	return result
}

func answersToDTO(answers []domain.Answer) []AnswerDTO {
	result := make([]AnswerDTO, 0, len(answers))
	for _, a := range answers {
		result = append(result, answerToDTO(a))
	}

	return result
}

func answerToDTO(t domain.Answer) AnswerDTO {
	return AnswerDTO{
		ID:        t.ID,
		Text:      t.Text,
		IsCorrect: t.IsCorrect,
	}
}
