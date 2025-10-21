// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"fmt"
	"strconv"

	"github.com/ctx42/nomix/pkg/nomix"
)

// IntSlice is a tag for a slice of int values.
type IntSlice = nomix.Slice[int]

// intSliceSpec defines the [nomix.Spec] for [IntSlice] type.
var intSliceSpec = nomix.NewSpec(
	nomix.KindIntSlice,
	nomix.TagCreateFunc(CreateIntSlice),
	nomix.TagParserNotImpl,
)

// IntSliceSpec returns a [nomix.Spec] for [IntSlice] type.
func IntSliceSpec() nomix.Spec { return intSliceSpec }

// NewIntSlice returns a new instance of [IntSlice].
func NewIntSlice(name string, val ...int) *IntSlice {
	return nomix.NewSlice(name, val, nomix.KindIntSlice, strValueIntSlice, nil)
}

// CreateIntSlice casts the value to []int. Returns the [IntSlice] instance
// with the given name and nil error on success. Returns nil and [ErrInvType]
// if the value is not []int type.
func CreateIntSlice(name string, val any, _ ...nomix.Option) (*IntSlice, error) {
	v, err := createIntSlice(val, nomix.Options{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewIntSlice(name, v...), nil
}

// createIntSlice casts the value to []int. Returns the []int and nil error on
// success. Returns nil and [ErrInvType] if the value is not the []int.
func createIntSlice(val any, _ nomix.Options) ([]int, error) {
	if v, ok := val.([]int); ok {
		return v, nil
	}
	return nil, nomix.ErrInvType
}

// strValueIntSlice converts an int slice to its string representation.
func strValueIntSlice(v []int) string {
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += strconv.Itoa(val)
	}
	return ret + "]"
}
