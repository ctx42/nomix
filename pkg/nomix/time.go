// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"slices"
	"time"
)

// Time is a tag for a single time.Time value.
type Time = single[time.Time]

// NewTime returns a new instance of [Time].
func NewTime(name string, v time.Time) *Time {
	return &single[time.Time]{
		name:     name,
		value:    v,
		kind:     KindTime,
		stringer: timeToString,
	}
}

// CreateTime casts the value to [time.Time], or when the value is a string, it
// parses it as [time.RFC3339Nano] time. Returns the [Time] instance with the
// given name and nil error on success. Returns nil and error if the value's
// type is not a supported type or the value is not a valid time representation.
func CreateTime(name string, val any, _ ...Option) (*Time, error) {
	v, err := createTime(val, defaultOptions)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewTime(name, v), nil
}

// ParseTime parses string representation of the time tag. The time is always
// returned as UTC.
func ParseTime(name, v string, opts ...Option) (*Time, error) {
	def := defaultOptions
	for _, opt := range opts {
		opt(&def)
	}
	val, err := parseTime(v, def)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewTime(name, val), nil
}

// parseTime parses a string representation of [time.Time]. Returns the time
// and nil error if the value is a valid time representation. Returns zero value
// time and [ErrInvFormat] if the value is not a valid time representation.
//
// To support string zero time values, use the [WithZeroTime] option.
func parseTime(val string, opts Options) (time.Time, error) {
	if opts.timeFormat == "" {
		return time.Time{}, ErrInvType
	}
	if slices.Contains(opts.zeroTime, val) {
		return time.Time{}, nil
	}
	var tim time.Time
	var err error
	if opts.location != nil {
		tim, err = time.ParseInLocation(opts.timeFormat, val, opts.location)
	} else {
		tim, err = time.Parse(opts.timeFormat, val)
	}
	if err != nil {
		return time.Time{}, ErrInvFormat
	}
	if tim.Location().String() == "" {
		tim = tim.UTC()
	}
	return tim, nil
}

// timeToString converts [time.Time] to its string representation.
func timeToString(v time.Time) string { return v.Format(time.RFC3339Nano) }

// createTime casts the value to [time.Time] or when the value is a string
// parses it as time but only when [Options.timeFormat] is set. Returns the
// time and nil error on success. Returns zero value time and error if the
// value's type is not a supported type or the value is not a valid time
// representation.
func createTime(val any, opts Options) (time.Time, error) {
	switch v := val.(type) {
	case time.Time:
		return v, nil
	case string:
		if opts.timeFormat != "" {
			return parseTime(v, opts)
		}
	}
	return time.Time{}, ErrInvType
}
