// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"strconv"
)

// IntSlice is a tag for a slice of int values.
type IntSlice = slice[int]

// intSliceSpec defines the [KindSpec] for [IntSlice] type.
var intSliceSpec = KindSpec{
	knd: KindIntSlice,
	tcr: CreateFunc(CreateIntSlice),
	tpr: func(name string, val string, opts ...Option) (Tag, error) {
		return nil, ErrNotImpl
	},
}

// IntSliceSpec returns a [KindSpec] for [IntSlice] type.
func IntSliceSpec() KindSpec { return intSliceSpec }

// NewIntSlice returns a new instance of [IntSlice].
func NewIntSlice(name string, v ...int) *IntSlice {
	return &slice[int]{
		name:     name,
		value:    v,
		kind:     KindIntSlice,
		stringer: intSliceToString,
	}
}

// CreateIntSlice casts the value to []int. Returns the [IntSlice] instance
// with the given name and nil error on success. Returns nil and [ErrInvType]
// if the value is not []int type.
func CreateIntSlice(name string, val any, _ ...Option) (*IntSlice, error) {
	v, err := createIntSlice(val, defaultOptions)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewIntSlice(name, v...), nil
}

// createIntSlice casts the value to []int. Returns the []int and nil error on
// success. Returns nil and [ErrInvType] if the value is not the []int.
func createIntSlice(val any, _ Options) ([]int, error) {
	if v, ok := val.([]int); ok {
		return v, nil
	}
	return nil, ErrInvType
}

// intSliceToString converts an int slice to its string representation.
func intSliceToString(v []int) string {
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += strconv.Itoa(val)
	}
	return ret + "]"
}
