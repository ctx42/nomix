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
func ParseBool(name, v string, _ ...Option) (Tag, error) {
	val, err := strconv.ParseBool(v)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, ErrInvFormat)
	}
	return NewBool(name, val), nil
}

// boolToString converts bool to its string representation.
func boolToString(v bool) string { return strconv.FormatBool(v) }
