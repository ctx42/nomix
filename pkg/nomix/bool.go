// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// Bool is a tag for a single bool value.
type Bool = Single[bool]

// boolSpec defines the [KindSpec] for [Bool] type.
var boolSpec = KindSpec{
	knd: KindBool,
	tcr: CreateFunc(CreateBool),
	tpr: ParseFunc(ParseBool),
}

// BoolSpec returns a [KindSpec] for [Bool] type.
func BoolSpec() KindSpec { return boolSpec }

// NewBool returns a new instance of [Bool].
func NewBool(name string, val bool) *Bool {
	return NewSingle(name, val, KindBool, strconv.FormatBool, sqlValueBool)
}

// CreateBool casts the given value to bool. Returns the [Bool] instance with
// the given name and nil error on success. Returns nil and [ErrInvType] if the
// value is not the bool type.
func CreateBool(name string, val any, _ ...Option) (*Bool, error) {
	v, err := createBool(val, defaultOptions)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewBool(name, v), nil
}

// createBool casts the given value to bool. Returns the bool and nil error on
// success. Returns false and [ErrInvType] if the value is not the bool type.
func createBool(val any, _ Options) (bool, error) {
	if v, ok := val.(bool); ok {
		return v, nil
	}
	return false, ErrInvType
}

// ParseBool parses string representation of the boolean tag.
func ParseBool(name, val string, _ ...Option) (*Bool, error) {
	vv, err := strconv.ParseBool(val)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, ErrInvFormat)
	}
	return NewBool(name, vv), nil
}

// sqlValueBool converts bool to int64 according to [KindBool] base type. Never
// returns an error.
func sqlValueBool(val bool) (driver.Value, error) {
	if val {
		return int64(1), nil
	}
	return int64(0), nil
}
