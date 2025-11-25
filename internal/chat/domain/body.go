package domain

import (
	"errors"
)

var (
	errEmptyBody = errors.New("empty body")
	errTooLong   = errors.New("empty body too long")
)

func NewBody(s string) (Body, error) {
	if len(s) == 0 {
		return "", errEmptyBody
	}
	if len(s) > 2000 {
		return "", errTooLong
	}
	return Body(s), nil
}
