package handler

import (
	"time"
	"vago/internal/domain"
)

type UserApiDTO struct {
	ID        int64     `json:"id"`
	Login     string    `json:"login"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"createdAt"`
}

type UsersApiDTO struct {
	Users []UserApiDTO `json:"users"`
}

func userToDTO(u domain.User) UserApiDTO {
	return UserApiDTO{
		ID:        u.ID,
		Login:     u.Login,
		Username:  u.Username,
		Role:      string(u.Role),
		Color:     u.Color,
		CreatedAt: u.CreatedAt,
	}
}

func usersToDTO(users []domain.User) []UserApiDTO {
	result := make([]UserApiDTO, 0, len(users))
	for _, u := range users {
		result = append(result, userToDTO(u))
	}
	return result
}

type TaskApiDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	Completed   bool      `json:"completed"`
}

type TasksApiDTO struct {
	Tasks []TaskApiDTO `json:"tasks"`
}

func taskToDTO(t domain.Task) TaskApiDTO {
	return TaskApiDTO{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		Completed:   t.Completed,
	}
}

func tasksToDTO(tasks []domain.Task) []TaskApiDTO {
	result := make([]TaskApiDTO, 0, len(tasks))
	for _, t := range tasks {
		result = append(result, taskToDTO(t))
	}

	return result
}

type PostTaskDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type UpdateTaskDTO struct {
	Completed bool `json:"completed"`
}

type TopicsAndQuestionsApiDTO struct {
	Topics    []TopicApiDTO    `json:"topics"`
	Questions []QuestionApiDTO `json:"questions"`
}

type TopicApiDTO struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	QuestionsCount int    `json:"questionsCount"`
}

type QuestionApiDTO struct {
	ID          int64  `json:"id"`
	TopicID     int64  `json:"topicId"`
	Text        string `json:"text"`
	Code        string `json:"code"`
	Explanation string `json:"explanation"`
}

func topicToDTO(t domain.TopicWithCount) TopicApiDTO {
	return TopicApiDTO{
		ID:             t.ID,
		Name:           t.Name,
		QuestionsCount: t.QuestionsCount,
	}
}

func topicsToDTO(topics []domain.TopicWithCount) []TopicApiDTO {
	result := make([]TopicApiDTO, 0, len(topics))
	for _, t := range topics {
		result = append(result, topicToDTO(t))
	}

	return result
}

func topicWithQuestionsToDTO(topics []domain.TopicWithCount, questions []*domain.Question) TopicsAndQuestionsApiDTO {
	return TopicsAndQuestionsApiDTO{Topics: topicsToDTO(topics), Questions: questionsToDTO(questions)}
}

func questionsToDTO(questions []*domain.Question) []QuestionApiDTO {
	result := make([]QuestionApiDTO, 0, len(questions))
	for _, t := range questions {
		result = append(result, questionToDTO(t))
	}

	return result
}

func questionToDTO(t *domain.Question) QuestionApiDTO {
	return QuestionApiDTO{
		ID:          t.ID,
		Text:        t.Text,
		Code:        t.Code,
		Explanation: t.Explanation,
		TopicID:     t.TopicID,
	}
}
