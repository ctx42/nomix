// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"strconv"
)

// ByteSlice is a tag for a slice of bytes.
type ByteSlice = Slice[byte]

// byteSliceSpec defines the [KindSpec] for [ByteSlice] type.
var byteSliceSpec = KindSpec{
	knd: KindByteSlice,
	tcr: CreateFunc(CreateByteSlice),
	tpr: func(name string, val string, opts ...Option) (Tag, error) {
		return nil, ErrNotImpl
	},
}

// ByteSliceSpec returns a [KindSpec] for [ByteSlice] type.
func ByteSliceSpec() KindSpec { return byteSliceSpec }

// NewByteSlice returns a new instance of [ByteSlice].
func NewByteSlice(name string, val ...byte) *ByteSlice {
	return NewSlice(name, val, KindByteSlice, strValueByteSlice, nil)
}

// CreateByteSlice casts the value to []byte. Returns the [ByteSlice] instance
// with the given name and nil error on success. Returns nil and [ErrInvType]
// if the value is not the []byte type.
func CreateByteSlice(name string, val any, _ ...Option) (*ByteSlice, error) {
	v, err := createByteSlice(val, defaultOptions)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewByteSlice(name, v...), nil
}

// createByteSlice casts the value to []byte. Returns the []byte and nil error
// on success. Returns nil and [ErrInvType] if the value is not the []byte type.
func createByteSlice(val any, _ Options) ([]byte, error) {
	if v, ok := val.([]byte); ok {
		return v, nil
	}
	return nil, ErrInvType
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
