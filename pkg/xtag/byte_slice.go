// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"fmt"
	"strconv"

	"github.com/ctx42/nomix/pkg/nomix"
)

// ByteSlice is a tag for a slice of bytes.
type ByteSlice = nomix.Slice[byte]

// byteSliceSpec defines the [nomix.Spec] for [ByteSlice] type.
var byteSliceSpec = nomix.NewSpec(
	nomix.KindByteSlice,
	nomix.TagCreateFunc(CreateByteSlice),
	nomix.TagParserNotImpl,
)

// ByteSliceSpec returns a [nomix.Spec] for [ByteSlice] type.
func ByteSliceSpec() nomix.Spec { return byteSliceSpec }

// NewByteSlice returns a new instance of [ByteSlice].
func NewByteSlice(name string, val ...byte) *ByteSlice {
	return nomix.NewSlice(
		name,
		val,
		nomix.KindByteSlice,
		strValueByteSlice,
		nil,
	)
}

// CreateByteSlice casts the value to []byte. Returns the [ByteSlice] instance
// with the given name and nil error on success. Returns nil and
// [nomix.ErrInvType] if the value is not the []byte type.
func CreateByteSlice(
	name string,
	val any,
	_ ...nomix.Option,
) (*ByteSlice, error) {

	if v, ok := val.([]byte); ok {
		return NewByteSlice(name, v...), nil
	}
	return nil, fmt.Errorf("%s: %w", name, nomix.ErrInvType)
}

// strValueByteSlice converts a byte slice to its string representation.
func strValueByteSlice(v []byte) string {
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += strconv.Itoa(int(val))
	}
	return ret + "]"
}
