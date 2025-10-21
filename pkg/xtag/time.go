// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/ctx42/nomix/pkg/nomix"
)

// Time is a tag representing a single [time.Time] value.
type Time = nomix.Single[time.Time]

// timeSpec defines the [nomix.Spec] for [Time] type.
var timeSpec = nomix.NewSpec(
	nomix.KindTime,
	nomix.TagCreateFunc(CreateTime),
	nomix.TagParseFunc(ParseTime),
)

// TimeSpec returns a [nomix.Spec] for [Time] type.
func TimeSpec() nomix.Spec { return timeSpec }

// NewTime returns a new instance of [Time].
func NewTime(name string, val time.Time) *Time {
	return nomix.NewSingle(
		name,
		val,
		nomix.KindTime,
		strValueTime,
		sqlValueTime,
	)
}

// CreateTime casts the value to [time.Time], or when the value is a string, it
// parses it as [time.RFC3339Nano] time. Returns the [Time] instance with the
// given name and nil error on success. Returns nil and error if the value's
// type is not a supported type or the value is not a valid time representation.
func CreateTime(name string, val any, opts ...nomix.Option) (*Time, error) {
	def := nomix.NewOptions(opts...)
	v, err := nomix.CreateTime(val, def)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewTime(name, v), nil
}

// ParseTime parses a string representation of [time.Time]. Returns the time
// and nil error if the value is a valid time representation. Returns zero
// value time and [nomix.ErrInvFormat] if the value is not a valid time
// representation.
//
// To support string zero time values, use the [WithZeroTime] option.
func ParseTime(name, val string, opts ...nomix.Option) (*Time, error) {
	def := nomix.NewOptions(opts...)
	v, err := nomix.ParseTime(val, def)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewTime(name, v), nil
}

// strValueTime converts [time.Time] to its string representation.
func strValueTime(v time.Time) string { return v.Format(time.RFC3339Nano) }

// sqlValueTime returns the value as is. Never returns an error.
func sqlValueTime(v time.Time) (driver.Value, error) { return v, nil }
