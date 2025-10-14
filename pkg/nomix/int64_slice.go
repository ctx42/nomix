// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"strconv"
)

// Int64Slice is a tag for a slice of int64 values.
type Int64Slice = slice[int64]

// NewInt64Slice returns a new instance of [Int64Slice].
func NewInt64Slice(name string, v ...int64) *Int64Slice {
	return &slice[int64]{
		name:     name,
		value:    v,
		kind:     KindInt64Slice,
		stringer: int64SliceToString,
	}
}

// int64SliceToString converts an int64 slice to its string representation.
func int64SliceToString(v []int64) string {
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += strconv.FormatInt(val, 10)
	}
	return ret + "]"
}

// convertableToInt64 lists types that can be upgraded to int64 without loss of
// precision.
type convertableToInt64 interface {
	int | int8 | int16 | int32 | int64
}

// asInt64Slice upgrades a slice of values that can be upgraded to []int64
// without loss of precision.
func toInt64Slice[T convertableToInt64](v []T, _ Options) []int64 {
	upgraded := make([]int64, len(v))
	for i, val := range v {
		upgraded[i] = int64(val)
	}
	return upgraded
}

// asInt64Slice casts the value to []int64. Returns the slice and nil error if
// the value is a []int, []int8, []int16, []int32, or []int64. Returns 0 and
// [ErrInvType] if the value is not a supported integer type.
func asInt64Slice(val any, opts Options) ([]int64, error) {
	switch v := val.(type) {
	case []int:
		return toInt64Slice(v, opts), nil
	case []int8:
		return toInt64Slice(v, opts), nil
	case []int16:
		return toInt64Slice(v, opts), nil
	case []int32:
		return toInt64Slice(v, opts), nil
	case []int64:
		return v, nil
	}
	return nil, ErrInvType
}
