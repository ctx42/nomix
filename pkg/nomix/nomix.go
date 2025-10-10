package nomix

import (
	"errors"
)

// Metadata parsing and casting errors.
var (
	// ErrMissing represents a missing set element.
	ErrMissing = errors.New("missing element")

	// ErrInvType represents an invalid element type.
	ErrInvType = errors.New("invalid element type")

	// ErrInvFormat represents an invalid element format.
	ErrInvFormat = errors.New("invalid element format")

	// ErrInvValue represents an invalid element value.
	ErrInvValue = errors.New("invalid element value")
)
