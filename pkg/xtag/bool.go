// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"database/sql/driver"
	"fmt"
	"strconv"

	"github.com/ctx42/nomix/pkg/nomix"
)

// Bool is a tag representing a single boolean value.
type Bool = nomix.Single[bool]

// boolSpec defines the [nomix.Spec] for [Bool] type.
var boolSpec = nomix.NewSpec(
	nomix.KindBool,
	nomix.TagCreateFunc(CreateBool),
	nomix.TagParseFunc(ParseBool),
)

// BoolSpec returns a [nomix.Spec] for [Bool] type.
func BoolSpec() nomix.Spec { return boolSpec }

// NewBool returns a new instance of [Bool].
func NewBool(name string, val bool) *Bool {
	return nomix.NewSingle(
		name,
		val,
		nomix.KindBool,
		strconv.FormatBool,
		sqlValueBool,
	)
}

// CreateBool casts the given value to bool. Returns the [Bool] instance with
// the given name and nil error on success. Returns nil and [nomix.ErrInvType]
// if the value is not the bool type.
func CreateBool(name string, val any, _ ...nomix.Option) (*Bool, error) {
	if v, ok := val.(bool); ok {
		return NewBool(name, v), nil
	}
	return nil, fmt.Errorf("%s: %w", name, nomix.ErrInvType)
}

// ParseBool parses string representation of the boolean tag.
func ParseBool(name, val string, _ ...nomix.Option) (*Bool, error) {
	vv, err := strconv.ParseBool(val)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, nomix.ErrInvFormat)
	}
	return NewBool(name, vv), nil
}

// sqlValueBool converts bool to int64 according to [nomix.KindBool] base type.
// Never returns an error.
func sqlValueBool(val bool) (driver.Value, error) {
	if val {
		return int64(1), nil
	}
	return int64(0), nil
}
