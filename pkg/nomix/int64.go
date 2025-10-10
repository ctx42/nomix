// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"strconv"
)

// Int64 is a tag for a single int64 value.
type Int64 = single[int64]

// NewInt64 returns a new instance of [Int64].
func NewInt64(name string, v int64) *Int64 {
	return &single[int64]{
		name:     name,
		value:    v,
		kind:     KindInt64,
		stringer: int64ToString,
	}
}

// ParseInt64 parses string representation of the 64-bit integer tag.
func ParseInt64(name, v string, opts ...Option) (*Int64, error) {
	def := DefaultOptions()
	for _, opt := range opts {
		opt(def)
	}
	val, err := strconv.ParseInt(v, def.intBase, 64)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, ErrInvFormat)
	}
	return NewInt64(name, val), nil
}

// int64ToString converts int64 to its string representation.
func int64ToString(v int64) string { return strconv.FormatInt(v, 10) }
