// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"database/sql/driver"
	"fmt"
	"strconv"

	"github.com/ctx42/nomix/pkg/nomix"
)

// Int is a tag representing a single integer value.
type Int = nomix.Single[int]

// intSpec defines the [nomix.Spec] for [Int] type.
var intSpec = nomix.NewSpec(
	nomix.KindInt,
	nomix.TagCreateFunc(CreateInt),
	nomix.TagParseFunc(ParseInt),
)

// IntSpec returns a [nomix.Spec] for [Int] type.
func IntSpec() nomix.Spec { return intSpec }

// NewInt returns a new instance of [Int].
func NewInt(name string, val int) *Int {
	return nomix.NewSingle(name, val, nomix.KindInt, strconv.Itoa, sqlValueInt)
}

// CreateInt casts the value to int. Returns the [Int] instance with the given
// type and nil error on success. Returns nil and [nomix.ErrInvType] if the
// value is not the int type.
func CreateInt(name string, val any, _ ...nomix.Option) (*Int, error) {
	if v, ok := val.(int); ok {
		return NewInt(name, v), nil
	}
	return nil, fmt.Errorf("%s: %w", name, nomix.ErrInvType)
}

// ParseInt parses string representation of the integer tag.
func ParseInt(name, val string, opts ...nomix.Option) (*Int, error) {
	def := nomix.NewOptions(opts...)
	v, err := strconv.ParseInt(val, def.Radix, 0)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, nomix.ErrInvFormat)
	}
	return NewInt(name, int(v)), nil
}

// sqlValueInt converts int to its int64 representation. Never returns an error.
func sqlValueInt(val int) (driver.Value, error) { return int64(val), nil }
