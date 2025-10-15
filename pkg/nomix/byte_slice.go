// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"strconv"
)

// ByteSlice is a tag for a slice of bytes.
type ByteSlice = slice[byte]

// NewByteSlice returns a new instance of [ByteSlice].
func NewByteSlice(name string, val ...byte) *ByteSlice {
	return &slice[byte]{
		name:     name,
		value:    val,
		kind:     KindByteSlice,
		stringer: byteSliceToString,
	}
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

// byteSliceToString converts a byte slice to its string representation.
func byteSliceToString(v []byte) string {
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += strconv.Itoa(int(val))
	}
	return ret + "]"
}

// createByteSlice casts the value to []byte. Returns the []byte and nil error
// on success. Returns nil and [ErrInvType] if the value is not the []byte type.
func createByteSlice(val any, _ Options) ([]byte, error) {
	if v, ok := val.([]byte); ok {
		return v, nil
	}
	return nil, ErrInvType
}
