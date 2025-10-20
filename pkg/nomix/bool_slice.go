// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"strconv"
)

// BoolSlice is a tag for a slice of bool values.
type BoolSlice = Slice[bool]

// boolSliceSpec defines the [KindSpec] for [BoolSlice] type.
var boolSliceSpec = KindSpec{
	knd: KindBoolSlice,
	tcr: CreateFunc(CreateBoolSlice),
	tpr: func(name string, val string, opts ...Option) (Tag, error) {
		return nil, ErrNotImpl
	},
}

// BoolSliceSpec returns a [KindSpec] for [BoolSlice] type.
func BoolSliceSpec() KindSpec { return boolSliceSpec }

// NewBoolSlice returns a new instance of [BoolSlice].
func NewBoolSlice(name string, val ...bool) *BoolSlice {
	return NewSlice(name, val, KindBoolSlice, strValueBoolSlice, nil)
}

// CreateBoolSlice casts the value to []bool. Returns the [BoolSlice] instance
// with the given name and nil error on success. Returns nil and [ErrInvType]
// if the value is the []bool type.
func CreateBoolSlice(name string, val any, _ ...Option) (*BoolSlice, error) {
	v, err := createBoolSlice(val, Options{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewBoolSlice(name, v...), nil
}

// createBoolSlice casts the value to []bool. Returns the []bool and nil error
// on success. Returns nil and [ErrInvType] if the value is not the []bool type.
func createBoolSlice(val any, _ Options) ([]bool, error) {
	if v, ok := val.([]bool); ok {
		return v, nil
	}
	return nil, ErrInvType
}

// strValueBoolSlice converts a bool slice to its string representation.
func strValueBoolSlice(v []bool) string {
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += strconv.FormatBool(val)
	}
	return ret + "]"
}
