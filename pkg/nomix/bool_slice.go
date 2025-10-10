// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"strconv"
)

// BoolSlice is a tag for a slice of bool values.
type BoolSlice = slice[bool]

// NewBoolSlice returns a new instance of [BoolSlice].
func NewBoolSlice(name string, v ...bool) *BoolSlice {
	return &slice[bool]{
		name:     name,
		value:    v,
		kind:     KindBoolSlice,
		stringer: boolSliceToString,
	}
}

// boolSliceToString converts a bool slice to its string representation.
func boolSliceToString(v []bool) string {
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += strconv.FormatBool(val)
	}
	return ret + "]"
}
