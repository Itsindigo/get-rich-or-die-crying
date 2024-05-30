package trading

import (
	"fmt"
	"strings"
)

type Error struct {
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

type ErrorWithLabel struct {
	Label string
	Error string
}

func (e *ErrorWithLabel) String() string {
	return fmt.Sprintf("%q: %q", e.Label, e.Error)
}

type CreateOrderError struct {
	ErrorLabels []ErrorWithLabel
}

func (e CreateOrderError) Error() string {
	errors := []string{}

	for _, errorWithLabel := range e.ErrorLabels {
		if len(errorWithLabel.Error) > 0 {
			errors = append(errors, errorWithLabel.String())
		}
	}

	return fmt.Sprintf("failed to create coinbase order: `{ %s }`", strings.Join(errors, ", "))
}
