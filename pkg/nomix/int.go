// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"strconv"
)

// Int is a tag for a single int value.
type Int = single[int]

// NewInt returns a new instance of [Int].
func NewInt(name string, v int) *Int {
	return &single[int]{
		name:     name,
		value:    v,
		kind:     KindInt,
		stringer: intToString,
	}
}

// ParseInt parses string representation of the integer tag.
func ParseInt(name, v string, opts ...Option) (*Int64, error) {
	def := defaultOptions
	for _, opt := range opts {
		opt(&def)
	}
	val, err := strconv.ParseInt(v, def.intBase, 0)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, ErrInvFormat)
	}
	return NewInt64(name, val), nil
}

// intToString converts int to its string representation.
func intToString(v int) string { return strconv.Itoa(v) }

// asInt casts the value to int. Returns the int and nil error on success.
// Returns false and [ErrInvType] if the value is not a supported type.
func asInt(val any, _ Options) (int, error) {
	if v, ok := val.(int); ok {
		return v, nil
	}
	return 0, ErrInvType
}
