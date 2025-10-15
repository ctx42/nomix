// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"strconv"
)

// ByteSlice is a tag for a slice of bytes.
type ByteSlice = slice[byte]

// NewByteSlice returns a new instance of [ByteSlice].
func NewByteSlice(name string, v ...byte) *ByteSlice {
	return &slice[byte]{
		name:     name,
		value:    v,
		kind:     KindByteSlice,
		stringer: byteSliceToString,
	}
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

// asByteSlice casts the value to []byte. Returns the slice and nil error on
// success. Returns nil and [ErrInvType] if the value is not a supported type.
func asByteSlice(val any, _ Options) ([]byte, error) {
	if v, ok := val.([]byte); ok {
		return v, nil
	}
	return nil, ErrInvType
}
