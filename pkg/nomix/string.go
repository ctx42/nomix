// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
)

// String is a tag for a single byte value.
type String = single[string]

// NewString returns a new instance of [String].
func NewString(name, v string) *String {
	return &single[string]{
		name:     name,
		value:    v,
		kind:     KindString,
		stringer: stringToString,
	}
}

// CreateString casts the value to a string. Returns the [String] instance with
// the given name and nil error on success. Returns nil and [ErrInvType] if the
// value is not the string type.
func CreateString(name string, v any, _ ...Option) (*String, error) {
	vv, err := createString(v, defaultOptions)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewString(name, vv), nil
}

// ParseString creates string type tag. Never returns an error.
func ParseString(name, v string, _ ...Option) (*String, error) {
	return NewString(name, v), nil
}

// stringToString converts string to string.
func stringToString(v string) string { return v }

// createString casts the value to a string. Returns the string and nil error
// if the value is a string. Returns an empty string and [ErrInvType] if the
// value is not the string type.
func createString(val any, _ Options) (string, error) {
	if v, ok := val.(string); ok {
		return v, nil
	}
	return "", ErrInvType
}
