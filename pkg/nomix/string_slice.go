// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
)

// StringSlice is a tag for a slice of strings.
type StringSlice = slice[string]

// NewStringSlice returns a new instance of [StringSlice].
func NewStringSlice(name string, v ...string) *StringSlice {
	return &slice[string]{
		name:     name,
		value:    v,
		kind:     KindStringSlice,
		stringer: stringSliceToString,
	}
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

// stringSliceToString converts a string slice to its string representation.
func stringSliceToString(v []string) string {
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += `"` + val + `"`
	}
	return ret + "]"
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
