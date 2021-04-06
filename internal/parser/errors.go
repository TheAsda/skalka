package parser

import "github.com/TheAsda/skalka/internal"

type Error struct {
	internal.Error
}

func NewError(message string) *Error {
	return &Error{Error: internal.Error{
		message,
		"parser",
	}}
}
