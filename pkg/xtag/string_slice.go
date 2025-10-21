// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"fmt"

	"github.com/ctx42/nomix/pkg/nomix"
)

// StringSlice is a tag for a slice of strings.
type StringSlice = nomix.Slice[string]

// stringSliceSpec defines the [nomix.Spec] for [StringSlice] type.
var stringSliceSpec = nomix.NewSpec(
	nomix.KindStringSlice,
	nomix.TagCreateFunc(CreateStringSlice),
	nomix.TagParserNotImpl,
)

// StringSliceSpec returns a [nomix.Spec] for [StringSlice] type.
func StringSliceSpec() nomix.Spec { return stringSliceSpec }

// NewStringSlice returns a new instance of [StringSlice].
func NewStringSlice(name string, val ...string) *StringSlice {
	return nomix.NewSlice(name, val, nomix.KindStringSlice, strValueStringSlice, nil)
}

// CreateStringSlice casts the value to []string. Returns the [StringSlice]
// instance with the given name and nil error on success. Returns nil and
// [ErrInvType] if the value is not the []string type.
func CreateStringSlice(name string, val any, _ ...nomix.Option) (*StringSlice, error) {
	v, err := createStringSlice(val, nomix.Options{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewStringSlice(name, v...), nil
}

// createStringSlice casts the value to []string. Returns the []string and nil
// error on success. Returns nil and [ErrInvType] if the value is not the
// []string type.
func createStringSlice(val any, _ nomix.Options) ([]string, error) {
	if v, ok := val.([]string); ok {
		return v, nil
	}
	return nil, nomix.ErrInvType
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
