// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"fmt"
	"strconv"

	"github.com/ctx42/nomix/pkg/nomix"
)

// BoolSlice is a tag representing multiple boolean values.
type BoolSlice = nomix.Slice[bool]

// boolSliceSpec defines the [nomix.Spec] for [BoolSlice] type.
var boolSliceSpec = nomix.NewSpec(
	nomix.KindBoolSlice,
	nomix.TagCreateFunc(CreateBoolSlice),
	nomix.TagParserNotImpl,
)

// BoolSliceSpec returns a [nomix.Spec] for [BoolSlice] type.
func BoolSliceSpec() nomix.Spec { return boolSliceSpec }

// NewBoolSlice returns a new instance of [BoolSlice].
func NewBoolSlice(name string, val ...bool) *BoolSlice {
	return nomix.NewSlice(
		name,
		val,
		nomix.KindBoolSlice,
		strValueBoolSlice,
		nil,
	)
}

// CreateBoolSlice casts the value to []bool. Returns the [BoolSlice] instance
// with the given name and nil error on success. Returns nil and
// [nomix.ErrInvType] if the value is the []bool type.
func CreateBoolSlice(
	name string,
	val any,
	_ ...nomix.Option,
) (*BoolSlice, error) {

	if v, ok := val.([]bool); ok {
		return NewBoolSlice(name, v...), nil
	}
	return nil, fmt.Errorf("%s: %w", name, nomix.ErrInvType)
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
