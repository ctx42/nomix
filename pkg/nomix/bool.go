// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"strconv"
)

// Bool is a tag for a single bool value.
type Bool = single[bool]

// NewBool returns a new instance of [Bool].
func NewBool(name string, val bool) *Bool {
	return &single[bool]{
		name:     name,
		value:    val,
		kind:     KindBool,
		stringer: boolToString,
	}
}

// CreateBool casts the given value to bool. Returns the [Bool] instance with
// the given name and nil error on success. Returns nil and [ErrInvType] if the
// value is not the bool type.
func CreateBool(name string, val any, _ ...Option) (*Bool, error) {
	v, err := createBool(val, defaultOptions)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewBool(name, v), nil
}

// ParseBool parses string representation of the boolean tag.
func ParseBool(name, val string, _ ...Option) (*Bool, error) {
	// TODO(rz): maybe all of constructor functions like this can return concrete types?
	vv, err := strconv.ParseBool(val)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, ErrInvFormat)
	}
	return NewBool(name, vv), nil
}

// boolToString converts bool to its string representation.
func boolToString(v bool) string { return strconv.FormatBool(v) }

// createBool casts the given value to bool. Returns the bool and nil error on
// success. Returns false and [ErrInvType] if the value is not the bool type.
func createBool(val any, _ Options) (bool, error) {
	if v, ok := val.(bool); ok {
		return v, nil
	}
	return false, ErrInvType
}
