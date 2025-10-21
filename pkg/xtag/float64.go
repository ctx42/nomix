// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"fmt"
	"strconv"

	"github.com/ctx42/nomix/pkg/nomix"
)

// Float64 is a tag representing a single float64 value.
type Float64 = nomix.Single[float64]

// float64Spec defines the [nomix.Spec] for [Float64] type.
var float64Spec = nomix.NewSpec(
	nomix.KindFloat64,
	nomix.TagCreateFunc(CreateFloat64),
	nomix.TagParseFunc(ParseFloat64),
)

// Float64Spec returns a [nomix.Spec] for [Float64] type.
func Float64Spec() nomix.Spec { return float64Spec }

// NewFloat64 returns a new instance of [Float64].
func NewFloat64(name string, val float64) *Float64 {
	return nomix.NewSingle(name, val, nomix.KindFloat64, float64ToString, nil)
}

// CreateFloat64 casts the value to float64. Returns the [Float64] instance
// with the given name and nil error if the value is a byte, int, int8, int16,
// int32, int64, float32, or float64. Returns nil and [nomix.ErrInvType] if the
// value's type is not a supported numeric type.
//
// NOTE: For int64 values outside ±2^53 range, the result is undefined.
func CreateFloat64(name string, val any, _ ...nomix.Option) (*Float64, error) {
	v, err := nomix.CreateFloat64(val)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewFloat64(name, v), nil
}

// ParseFloat64 parses string representation of the 64-bit floating point tag.
func ParseFloat64(name, v string, _ ...nomix.Option) (*Float64, error) {
	val, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, nomix.ErrInvFormat)
	}
	return NewFloat64(name, val), nil
}

// float64ToString converts float64 to its string representation.
func float64ToString(v float64) string {
	return strconv.FormatFloat(v, 'g', -1, 64)
}
