// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"strconv"
)

// IntSlice is a tag for a slice of int values.
type IntSlice = slice[int]

// NewIntSlice returns a new instance of [IntSlice].
func NewIntSlice(name string, v ...int) *IntSlice {
	return &slice[int]{
		name:     name,
		value:    v,
		kind:     KindIntSlice,
		stringer: intSliceToString,
	}
}

// intSliceToString converts an int slice to its string representation.
func intSliceToString(v []int) string {
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += strconv.Itoa(val)
	}
	return ret + "]"
}

// asIntSlice casts the value to []int. Returns the slice and nil error if the
// value is a []int. Returns nil and [ErrInvType] if the value is not []int.
func asIntSlice(val any, _ Options) ([]int, error) {
	if v, ok := val.([]int); ok {
		return v, nil
	}
	return nil, ErrInvType
}
