package internal

import "fmt"

type Error struct {
	Message string
	Scope   string
}

func (e Error) Error() string {
	return fmt.Sprintf("[%s]: %s", e.Scope, e.Message)
}
