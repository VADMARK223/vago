package chat

import (
	"context"
	"vago/internal/domain"
	"vago/pkg/timex"
)

type Service struct {
	msgRepo  domain.MessageRepository
	userRepo domain.UserRepository
}

func NewService(messageRepo domain.MessageRepository, userRepo domain.UserRepository) *Service {
	return &Service{msgRepo: messageRepo, userRepo: userRepo}
}

func (s *Service) CreateMessage(ctx context.Context, dto MessageCreateDTO) (MessageDTO, error) {
	msg := domain.NewMessage(domain.UserID(dto.AuthorID), domain.Body(dto.Body), dto.MessageType)
	id, err := s.msgRepo.Save(ctx, msg)
	if err != nil {
		return MessageDTO{}, err
	}

	user, err := s.userRepo.GetByID(dto.AuthorID)
	if err != nil {
		return MessageDTO{}, err
	}

	return MessageDTO{
		ID:          id,
		AuthorID:    domain.UserID(user.ID),
		Username:    user.Username,
		Body:        dto.Body,
		MessageType: dto.MessageType,
		SentAt:      timex.Format(msg.SentAt()),
	}, nil
}

func (s *Service) ListMessagesWithAuthors(ctx context.Context) ([]MessageDTO, error) {
	msgs, err := s.msgRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(msgs) == 0 {
		return []MessageDTO{}, nil
	}

	// Collect unique user ids (dedup)
	seen := make(map[int64]struct{}, len(msgs))
	ids := make([]int64, 0, len(msgs))
	for _, m := range msgs {
		uid := int64(m.Author())

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

	// Build DTO
	result := make([]MessageDTO, 0, len(msgs))
	for _, m := range msgs {
		uid := int64(m.Author())
		username := userMap[uid]
		if username == "" {
			username = "Unknown"
		}

		result = append(result, MessageDTO{
			ID:          m.ID,
			AuthorID:    domain.UserID(uid),
			Username:    username,
			Body:        string(m.Body()),
			SentAt:      timex.Format(m.SentAt()),
			MessageType: m.MessageType,
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
