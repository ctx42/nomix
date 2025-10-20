// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"strconv"
)

// Int64 is a tag for a single int64 value.
type Int64 = Single[int64]

// int64Spec defines the [KindSpec] for [Int64] type.
var int64Spec = KindSpec{
	knd: KindInt64,
	tcr: CreateFunc(CreateInt64),
	tpr: ParseFunc(ParseInt64),
}

// Int64Spec returns a [KindSpec] for [Int64] type.
func Int64Spec() KindSpec { return int64Spec }

// NewInt64 returns a new instance of [Int64].
func NewInt64(name string, val int64) *Int64 {
	return NewSingle(name, val, KindInt64, strValueInt64, nil)
}

// CreateInt64 casts the value to int64. Returns the [Int64] instance with the
// given name and nil error if the value is a byte, int, int8, int16, int32, or
// int64. Returns nil and [ErrInvType] if the value's type is not a supported
// numeric type.
func CreateInt64(name string, val any, _ ...Option) (*Int64, error) {
	v, err := createInt64(val, defaultOptions)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewInt64(name, v), nil
}

// createInt64 casts the value to int64. Returns the int64 and nil error if the
// value is a byte, int, int8, int16, int32, or int64. Returns 0 and
// [ErrInvType] if the value is not a supported integer type.
func createInt64(val any, _ Options) (int64, error) {
	switch v := val.(type) {
	case byte:
		return int64(v), nil
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	}
	return 0, ErrInvType
}

// ParseInt64 parses string representation of the 64-bit integer tag.
func ParseInt64(name, v string, opts ...Option) (*Int64, error) {
	def := defaultOptions
	for _, opt := range opts {
		opt(&def)
	}
	val, err := strconv.ParseInt(v, def.intBase, 64)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, ErrInvFormat)
	}
	return NewInt64(name, val), nil
}

// strValueInt64 converts int64 to its string representation.
func strValueInt64(v int64) string { return strconv.FormatInt(v, 10) }
