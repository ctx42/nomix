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
	vv, err := strconv.ParseBool(val)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, ErrInvFormat)
	}
	return NewBool(name, vv), nil
}

// boolToString converts bool to its string representation.
func boolToString(v bool) string { return strconv.FormatBool(v) }

// asBool casts the value to bool. Returns the bool and nil error on success.
// Returns false and [ErrInvType] if the value is not a supported type.
func asBool(val any, _ Options) (bool, error) {
	if v, ok := val.(bool); ok {
		return v, nil
	}
	return false, ErrInvType
}
