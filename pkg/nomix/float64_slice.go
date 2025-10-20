// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"strconv"
)

// Float64Slice is a tag for a slice of float64 values.
type Float64Slice = Slice[float64]

// float64SliceSpec defines the [KindSpec] for [Float64Slice] type.
var float64SliceSpec = KindSpec{
	knd: KindFloat64Slice,
	tcr: CreateFunc(CreateFloat64Slice),
	tpr: func(name string, val string, opts ...Option) (Tag, error) {
		return nil, ErrNotImpl
	},
}

// Float64SliceSpec returns a [KindSpec] for [Float64Slice] type.
func Float64SliceSpec() KindSpec { return float64SliceSpec }

// NewFloat64Slice returns a new instance of [Float64Slice].
func NewFloat64Slice(name string, val ...float64) *Float64Slice {
	return NewSlice(name, val, KindFloat64Slice, strValueFloat64Slice, nil)
}

// CreateFloat64Slice casts the value to []float64. Returns the [Float64Slice]
// instance with the given name and nil error if the value is a []int, []int8,
// []int16, []int32, []int64, []float32, or []float64. Returns nil and
// [ErrInvType] if the value's type is not a supported numeric slice type.
func CreateFloat64Slice(name string, val any, _ ...Option) (*Float64Slice, error) {
	v, err := createFloat64Slice(val, Options{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewFloat64Slice(name, v...), nil
}

// strValueFloat64Slice converts a float64 slice to its string representation.
func strValueFloat64Slice(v []float64) string {
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
// TODO(rz): Return an error when above happens.
func toFloat64Slice[T convertableToFloat64](v []T, _ Options) []float64 {
	upgraded := make([]float64, len(v))
	for i, val := range v {
		upgraded[i] = float64(val)
	}
	return upgraded
}

// createFloat64Slice casts the value to []float64. Returns the slice and nil
// error if the value is a []int, []int8, []int16, []int32, []int64, []float32,
// or []float64. Returns nil and [ErrInvType] if the value's type is not a
// supported numeric slice type.
//
// NOTE: For int64 values outside ±2^53 range, the result is undefined.
// TODO(rz): Return an error when above happens.
func createFloat64Slice(val any, opts Options) ([]float64, error) {
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
