package schema

import (
	"errors"
)

var (
	// ErrInvalidTokenSize is returned when session token size is invalid
	ErrInvalidTokenSize = errors.New("invalid token size")

	// ErrContainsSpaces is returned when field contains spaces
	ErrContainsSpaces = errors.New("field should not contain spaces")
)
