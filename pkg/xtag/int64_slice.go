// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"fmt"
	"strconv"

	"github.com/ctx42/nomix/pkg/nomix"
)

// Int64Slice is a tag representing multiple int64 values.
type Int64Slice = nomix.Slice[int64]

// int64SliceSpec defines the [nomix.Spec] for [Int64Slice] type.
var int64SliceSpec = nomix.NewSpec(
	nomix.KindInt64Slice,
	nomix.TagCreateFunc(CreateInt64Slice),
	nomix.TagParserNotImpl,
)

// Int64SliceSpec returns a [nomix.Spec] for [Int64Slice] type.
func Int64SliceSpec() nomix.Spec { return int64SliceSpec }

// NewInt64Slice returns a new instance of [Int64Slice].
func NewInt64Slice(name string, val ...int64) *Int64Slice {
	return nomix.NewSlice(
		name,
		val,
		nomix.KindInt64Slice,
		strValueInt64Slice,
		nil,
	)
}

// CreateInt64Slice casts the value to []int64. Returns the [Int64Slice]
// instance with the given name and nil error if the value is a []int, []int8,
// []int16, []int32, or []int64. Returns nil and [nomix.ErrInvType] if the
// value's type is not a supported numeric slice type.
func CreateInt64Slice(
	name string,
	val any,
	_ ...nomix.Option,
) (*Int64Slice, error) {

	v, err := nomix.CreateInt64Slice(val)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewInt64Slice(name, v...), nil
}

// strValueInt64Slice converts an int64 slice to its string representation.
func strValueInt64Slice(v []int64) string {
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += strconv.FormatInt(val, 10)
	}
	return ret + "]"
}
