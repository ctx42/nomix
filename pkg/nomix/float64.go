// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"strconv"
)

// Float64 is a tag for a single float64 value.
type Float64 = single[float64]

// NewFloat64 returns a new instance of [Float64].
func NewFloat64(name string, v float64) *Float64 {
	return &single[float64]{
		name:     name,
		value:    v,
		kind:     KindFloat64,
		stringer: float64ToString,
	}
}

// ParseFloat64 parses string representation of the 64-bit floating point tag.
func ParseFloat64(name, v string, _ ...Option) (*Float64, error) {
	val, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, ErrInvFormat)
	}
	return NewFloat64(name, val), nil
}

// float64ToString converts float64 to its string representation.
func float64ToString(v float64) string {
	return strconv.FormatFloat(v, 'g', -1, 64)
}

// asFloat64 casts the value to float64. Returns the float64 and nil error if
// the value is a byte, int, int8, int16, int32, int64, float32, or float64.
// Returns 0.0 and [ErrInvType] if the value is not a supported numeric type.
//
// NOTE: For int64 values outside ±2^53 range, the result is undefined.
func asFloat64(val any, _ Options) (float64, error) {
	switch v := val.(type) {
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case byte:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	}
	return 0, ErrInvType
}
