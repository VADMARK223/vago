package chat

type MessageCreateDTO struct {
	AuthorID    int64  `json:"author_id"`
	Body        string `json:"body"`
	MessageType string `json:"type"`
}

type MessageDTO struct {
	ID          int64  `json:"id"`
	AuthorID    int64  `json:"author_id"`
	Username    string `json:"username"`
	Body        string `json:"body"`
	SentAt      string `json:"sent_at"`
	MessageType string `json:"type"`
}
