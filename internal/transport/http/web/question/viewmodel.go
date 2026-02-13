package question

import shared "vago/internal/transport/http/shared/question"

type ViewModel struct {
	TopicID   int64
	Topics    any // если у тебя шаблон ожидает доменный тип — можешь оставить домен
	Questions any
}

func ToViewModel(d shared.Data) ViewModel {
	return ViewModel{
		TopicID:   d.TopicID,
		Topics:    d.Topics,    // можно не маппить, если HTML и так жрет домен
		Questions: d.Questions, // аналогично
	}
}
