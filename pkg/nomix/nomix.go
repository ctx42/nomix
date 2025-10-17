// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

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

	// ErrNoCreator represents a missing tag creator for a type.
	ErrNoCreator = errors.New("creator not found")

	// ErrNoSpec represents a missing tag spec for a type.
	ErrNoSpec = errors.New("spec not found")

	// ErrNotImpl represents a missing method implementation or
	// functionality for a type.
	ErrNotImpl = errors.New("not implemented")
)
