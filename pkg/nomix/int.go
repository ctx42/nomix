// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// Int is a tag for a single int value.
type Int = Single[int]

// intSpec defines the [KindSpec] for [Int] type.
var intSpec = KindSpec{
	knd: KindInt,
	tcr: CreateFunc(CreateInt),
	tpr: ParseFunc(ParseInt),
}

// IntSpec returns a [KindSpec] for [Int] type.
func IntSpec() KindSpec { return intSpec }

// NewInt returns a new instance of [Int].
func NewInt(name string, val int) *Int {
	return NewSingle(name, val, KindInt, strconv.Itoa, sqlValueInt)
}

// CreateInt casts the value to int. Returns the [Int] instance with the given
// type and nil error on success. Returns nil and [ErrInvType] if the value is
// not the int type.
func CreateInt(name string, val any, _ ...Option) (*Int, error) {
	v, err := createInt(val, defaultOptions)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewInt(name, v), nil
}

// createInt casts the value to int. Returns the int and nil error on success.
// Returns 0 and [ErrInvType] if the value is not the int type.
func createInt(val any, _ Options) (int, error) {
	if v, ok := val.(int); ok {
		return v, nil
	}
	return 0, ErrInvType
}

// ParseInt parses string representation of the integer tag.
func ParseInt(name, v string, opts ...Option) (*Int, error) {
	def := defaultOptions
	for _, opt := range opts {
		opt(&def)
	}
	val, err := strconv.ParseInt(v, def.intBase, 0)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, ErrInvFormat)
	}
	return NewInt(name, int(val)), nil
}

// sqlValueInt converts int to its int64 representation. Never returns an error.
func sqlValueInt(val int) (driver.Value, error) { return int64(val), nil }
