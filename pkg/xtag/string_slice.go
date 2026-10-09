// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac.
// SPDX-License-Identifier: MIT

package xtag

import (
	"fmt"
	"strconv"

	"github.com/ctx42/nomix/pkg/nomix"
)

// StringSlice is a tag for a slice of strings.
//
// Its SQL value is the []string itself, so storing it requires a database
// driver that accepts array arguments.
type StringSlice = nomix.Slice[string]

// stringSliceSpec defines the [nomix.KindSpec] for [StringSlice] type.
var stringSliceSpec = nomix.NewKindSpec(
	nomix.KindStringSlice,
	nomix.TagCreateFunc(CreateStringSlice),
	nomix.TagParserNotImpl,
)

// StringSliceSpec returns a [nomix.KindSpec] for [StringSlice] type.
func StringSliceSpec() nomix.KindSpec { return stringSliceSpec }

// NewStringSlice returns a new instance of [StringSlice].
func NewStringSlice(name string, val ...string) *StringSlice {
	return nomix.NewSlice(
		name,
		val,
		nomix.KindStringSlice,
		strValueStringSlice,
		nil,
	)
}

// CreateStringSlice casts the value to []string. Returns the [StringSlice]
// instance with the given name and nil error on success. Returns nil and
// [nomix.ErrInvType] if the value is not the []string type.
func CreateStringSlice(
	name string,
	val any,
	_ ...nomix.Option,
) (*StringSlice, error) {

	v, err := createStringSlice(val)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewStringSlice(name, v...), nil
}

// createStringSlice casts the value to []string. Returns the []string and nil
// error on success. Returns nil and [nomix.ErrInvType] if the value is not
// the []string type.
func createStringSlice(val any) ([]string, error) {
	if v, ok := val.([]string); ok {
		return v, nil
	}
	return nil, nomix.ErrInvType
}

// strValueStringSlice converts a string slice to its string representation.
func strValueStringSlice(v []string) string {
	return formatSlice(v, strconv.Quote)
}
