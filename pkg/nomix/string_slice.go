// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
)

// StringSlice is a tag for a slice of strings.
type StringSlice = Slice[string]

// stringSliceSpec defines the [KindSpec] for [StringSlice] type.
var stringSliceSpec = KindSpec{
	knd: KindStringSlice,
	tcr: CreateFunc(CreateStringSlice),
	tpr: func(name string, val string, opts ...Option) (Tag, error) {
		return nil, ErrNotImpl
	},
}

// StringSliceSpec returns a [KindSpec] for [StringSlice] type.
func StringSliceSpec() KindSpec { return stringSliceSpec }

// NewStringSlice returns a new instance of [StringSlice].
func NewStringSlice(name string, val ...string) *StringSlice {
	return NewSlice(name, val, KindStringSlice, strValueStringSlice, nil)
}

// CreateStringSlice casts the value to []string. Returns the [StringSlice]
// instance with the given name and nil error on success. Returns nil and
// [ErrInvType] if the value is not the []string type.
func CreateStringSlice(name string, val any, _ ...Option) (*StringSlice, error) {
	v, err := createStringSlice(val, defaultOptions)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewStringSlice(name, v...), nil
}

// createStringSlice casts the value to []string. Returns the []string and nil
// error on success. Returns nil and [ErrInvType] if the value is not the
// []string type.
func createStringSlice(val any, _ Options) ([]string, error) {
	if v, ok := val.([]string); ok {
		return v, nil
	}
	return nil, ErrInvType
}

// strValueStringSlice converts a string slice to its string representation.
func strValueStringSlice(v []string) string {
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += `"` + val + `"`
	}
	return ret + "]"
}
