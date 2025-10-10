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
