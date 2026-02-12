package dto

import (
	"vago/internal/application/test"
	"vago/internal/domain"
)

type Me struct {
	Username string      `json:"username"`
	Role     domain.Role `json:"role"`
}

type QuestionPublicResponse struct {
	ID        int64                  `json:"id"`
	Text      string                 `json:"text"`
	Code      string                 `json:"code"`
	TopicName string                 `json:"topicName"`
	Answers   []AnswerPublicResponse `json:"answers"`
}

type AnswerPublicResponse struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
}

func QuestionPublicToDTO(q test.QuestionPublic) QuestionPublicResponse {
	return QuestionPublicResponse{
		ID:        q.ID,
		Text:      q.Text,
		Code:      q.Code,
		TopicName: q.TopicName,
		Answers:   answersPublicToDTO(q.Answers),
	}
}

func answersPublicToDTO(answers []test.AnswerPublic) []AnswerPublicResponse {
	result := make([]AnswerPublicResponse, 0, len(answers))
	for _, answer := range answers {
		result = append(result, answerPublicToDTO(answer))
	}
	return result
}

func answerPublicToDTO(q test.AnswerPublic) AnswerPublicResponse {
	return AnswerPublicResponse{
		ID:   q.ID,
		Text: q.Text,
	}
}
