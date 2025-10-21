// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"fmt"
	"strconv"

	"github.com/ctx42/nomix/pkg/nomix"
)

// Float64Slice is a tag representing multiple boolean values.
type Float64Slice = nomix.Slice[float64]

// float64SliceSpec defines the [nomix.Spec] for [Float64Slice] type.
var float64SliceSpec = nomix.NewSpec(
	nomix.KindFloat64Slice,
	nomix.TagCreateFunc(CreateFloat64Slice),
	nomix.TagParserNotImpl,
)

// Float64SliceSpec returns a [nomix.Spec] for [Float64Slice] type.
func Float64SliceSpec() nomix.Spec { return float64SliceSpec }

// NewFloat64Slice returns a new instance of [Float64Slice].
func NewFloat64Slice(name string, val ...float64) *Float64Slice {
	return nomix.NewSlice(
		name,
		val,
		nomix.KindFloat64Slice,
		strValueFloat64Slice,
		nil,
	)
}

// CreateFloat64Slice casts the value to []float64. Returns the [Float64Slice]
// instance with the given name and nil error if the value is a []int, []int8,
// []int16, []int32, []int64, []float32, or []float64. Returns nil and
// [nomix.ErrInvType] if the value's type is not a supported numeric slice type.
//
// NOTE: For values outside ±2^53 range, the function will return an error.
func CreateFloat64Slice(
	name string,
	val any,
	_ ...nomix.Option,
) (*Float64Slice, error) {

	v, err := nomix.CreateFloat64Slice(val)
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
