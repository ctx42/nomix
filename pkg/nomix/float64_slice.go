// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"strconv"
)

// Float64Slice is a tag for a slice of float64 values.
type Float64Slice = slice[float64]

// NewFloat64Slice returns a new instance of [Float64Slice].
func NewFloat64Slice(name string, v ...float64) *Float64Slice {
	return &slice[float64]{
		name:     name,
		value:    v,
		kind:     KindFloat64Slice,
		stringer: float64SliceToString,
	}
}

// float64SliceToString converts a float64 slice to its string representation.
func float64SliceToString(v []float64) string {
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += strconv.FormatFloat(val, 'g', -1, 64)
	}
	return ret + "]"
}
