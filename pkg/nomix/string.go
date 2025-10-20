// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
)

// String is a tag for a single byte value.
type String = Single[string]

// stringSpec defines the [KindSpec] for [String] type.
var stringSpec = KindSpec{
	knd: KindString,
	tcr: CreateFunc(CreateString),
	tpr: func(name string, val string, opts ...Option) (Tag, error) {
		return NewString(name, val), nil
	},
}

// StringSpec returns a [KindSpec] for [String] type.
func StringSpec() KindSpec { return stringSpec }

// NewString returns a new instance of [String].
func NewString(name, val string) *String {
	return NewSingle(name, val, KindString, strValueString, nil)
}

// CreateString casts the value to a string. Returns the [String] instance with
// the given name and nil error on success. Returns nil and [ErrInvType] if the
// value is not the string type.
func CreateString(name string, v any, _ ...Option) (*String, error) {
	vv, err := createString(v, defaultOptions)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewString(name, vv), nil
}

// createString casts the value to a string. Returns the string and nil error
// if the value is a string. Returns an empty string and [ErrInvType] if the
// value is not the string type.
func createString(val any, _ Options) (string, error) {
	if v, ok := val.(string); ok {
		return v, nil
	}
	return "", ErrInvType
}

// strValueString returns the string as is.
func strValueString(v string) string { return v }
