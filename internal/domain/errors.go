package domain

import "errors"

var (
	ErrLoginExists  = errors.New("пользователь с таким логином уже существует")
	ErrValueTooLong = errors.New("значение слишком длинное")

	ErrUserNotFound = errors.New("пользователь не найден")

	ErrIncorrectPassword = errors.New("неверный пароль")

	ErrNoQuestion = errors.New("вопросов нет")
)
