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
func NewBool(name string, v bool) *Bool {
	return &single[bool]{
		name:     name,
		value:    v,
		kind:     KindBool,
		stringer: boolToString,
	}
}

// ParseBool parses string representation of the boolean tag.
func ParseBool(name, val string, _ ...Option) (Tag, error) {
	v, err := strconv.ParseBool(val)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, ErrInvFormat)
	}
	return NewBool(name, v), nil
}

// boolToString converts bool to its string representation.
func boolToString(v bool) string { return strconv.FormatBool(v) }

// asBool casts the value to bool or when the value is a string, parses it
// using [strconv.ParseBool]. Returns the bool and nil error on success.
// Returns false and [ErrInvType] if the value is not a supported type.
func asBool(val any, _ Options) (bool, error) {
	switch v := val.(type) {
	case bool:
		return v, nil
	case string:
		vv, err := strconv.ParseBool(v)
		if err != nil {
			return false, ErrInvFormat
		}
		return vv, nil
	}
	return false, ErrInvType
}
