package message

import (
	"context"
	"vago/internal/domain"
)

type Service struct {
	msgRepo  domain.MessageRepository
	userRepo domain.UserRepository
}

func NewService(messageRepo domain.MessageRepository, userRepo domain.UserRepository) *Service {
	return &Service{msgRepo: messageRepo, userRepo: userRepo}
}

func (s *Service) ListMessagesWithAuthors(ctx context.Context) ([]WithUsername, error) {
	msgs, err := s.msgRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(msgs) == 0 {
		return []WithUsername{}, nil
	}

	// Collect unique user ids (dedup)
	seen := make(map[int64]struct{}, len(msgs))
	ids := make([]int64, 0, len(msgs))
	for _, msg := range msgs {
		uid := int64(msg.Author())

		if _, exists := seen[uid]; exists {
			continue
		}

		seen[uid] = struct{}{}
		ids = append(ids, uid)
	}

	// Load users
	users, err := s.userRepo.GetByIDs(ids)
	if err != nil {
		return nil, err
	}

	// Build user map
	userMap := map[int64]string{}
	for _, u := range users {
		userMap[u.ID] = u.Username
	}

	result := make([]WithUsername, 0, len(msgs))
	for _, msg := range msgs {
		uid := int64(msg.Author())
		username := userMap[uid]
		if username == "" {
			username = "Unknown"
		}

		result = append(result, WithUsername{
			ID:          msg.ID,
			AuthorID:    domain.UserID(uid),
			Username:    username,
			Body:        string(msg.Body()),
			SentAt:      msg.SentAt(),
			MessageType: msg.MessageType,
		})
	}

	return result, nil
}

func (s *Service) DeleteMessage(id int64) error {
	return s.msgRepo.DeleteMessage(id)
}

func (s *Service) DeleteAll() error {
	return s.msgRepo.DeleteAll()
}
