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

// convertableToFloat64 lists types that can be upgraded to float64 without
// loss of precision.
type convertableToFloat64 interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// toFloat64Slice upgrades a slice of values that can be upgraded to
// []float64 without loss of precision.
//
// NOTE: For int64 values outside ±2^53 range, the result is undefined.
func toFloat64Slice[T convertableToFloat64](v []T, _ Options) []float64 {
	upgraded := make([]float64, len(v))
	for i, val := range v {
		upgraded[i] = float64(val)
	}
	return upgraded
}

// asFloat64Slice casts the value to []float64. Returns the slice and nil error
// if the value is a []int, []int8, []int16, []int32, []int64, []float32, or
// []float64. Returns 0.0 and [ErrInvType] if the value is not a supported
// numeric type.
func asFloat64Slice(val any, opts Options) ([]float64, error) {
	switch v := val.(type) {
	case []int:
		return toFloat64Slice(v, opts), nil
	case []int8:
		return toFloat64Slice(v, opts), nil
	case []int16:
		return toFloat64Slice(v, opts), nil
	case []int32:
		return toFloat64Slice(v, opts), nil
	case []int64:
		return toFloat64Slice(v, opts), nil
	case []float32:
		return toFloat64Slice(v, opts), nil
	case []float64:
		return v, nil
	}
	return nil, ErrInvType
}
