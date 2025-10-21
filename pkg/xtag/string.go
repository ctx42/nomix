// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"fmt"

	"github.com/ctx42/nomix/pkg/nomix"
)

// String is a tag for a slice of strings.
type String = nomix.Single[string]

// stringSpec defines the [nomix.Spec] for [String] type.
var stringSpec = nomix.NewSpec(
	nomix.KindString,
	nomix.TagCreateFunc(CreateString),
	func(name string, val string, opts ...nomix.Option) (nomix.Tag, error) {
		return NewString(name, val), nil
	},
)

// StringSpec returns a [nomix.Spec] for [String] type.
func StringSpec() nomix.Spec { return stringSpec }

// NewString returns a new instance of [String].
func NewString(name, val string) *String {
	return nomix.NewSingle(name, val, nomix.KindString, strValueString, nil)
}

// CreateString casts the value to a string. Returns the [String] instance with
// the given name and nil error on success. Returns nil and [nomix.ErrInvType]
// if the value is not the string type.
func CreateString(name string, val any, _ ...nomix.Option) (*String, error) {
	if v, ok := val.(string); ok {
		return NewString(name, v), nil
	}
	return nil, fmt.Errorf("%s: %w", name, nomix.ErrInvType)
}

// strValueString returns the string as is.
func strValueString(v string) string { return v }
