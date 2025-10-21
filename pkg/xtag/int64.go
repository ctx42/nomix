// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"fmt"
	"strconv"

	"github.com/ctx42/nomix/pkg/nomix"
)

// Int64 is a tag representing a single int64 value.
type Int64 = nomix.Single[int64]

// int64Spec defines the [nomix.Spec] for [Int64] type.
var int64Spec = nomix.NewSpec(
	nomix.KindInt64,
	nomix.TagCreateFunc(CreateInt64),
	nomix.TagParseFunc(ParseInt64),
)

// Int64Spec returns a [nomix.Spec] for [Int64] type.
func Int64Spec() nomix.Spec { return int64Spec }

// NewInt64 returns a new instance of [Int64].
func NewInt64(name string, val int64) *Int64 {
	return nomix.NewSingle(name, val, nomix.KindInt64, strValueInt64, nil)
}

// CreateInt64 casts the value to int64. Returns the [Int64] instance with the
// given name and nil error if the value is a byte, int, int8, int16, int32, or
// int64. Returns nil and [nomix.ErrInvType] if the value's type is not a
// supported numeric type.
func CreateInt64(name string, val any, _ ...nomix.Option) (*Int64, error) {
	v, err := nomix.CreateInt64(val)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewInt64(name, v), nil
}

// ParseInt64 parses string representation of the 64-bit integer tag.
func ParseInt64(name, v string, opts ...nomix.Option) (*Int64, error) {
	def := nomix.NewOptions(opts...)
	val, err := strconv.ParseInt(v, def.Radix, 64)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, nomix.ErrInvFormat)
	}
	return NewInt64(name, val), nil
}

// strValueInt64 converts int64 to its string representation.
func strValueInt64(v int64) string { return strconv.FormatInt(v, 10) }
