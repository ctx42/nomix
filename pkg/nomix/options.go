// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"time"
)

// Option represents an option function.
type Option func(*Options)

// Options represent a set of options used by [MetaSet] and [TagSet].
type Options struct {
	// Initial map size.
	//
	// Used by [NewTagSet] and [NewMetaSet] to allocate the init map.
	Length int

	// Initial map.
	//
	// Set by [WithMeta] and [WithTags] functions to initialize a map.
	init any

	// Time format.
	//
	// When set, [MetaSet.MetaGetTime] will allow time to be represented as a
	// string.
	TimeFormat string

	// Location to parse format.
	//
	// When set [time.ParseInLocation] instead of [time.Parse] in the
	// [MetaSet.MetaGetTime] to parse strings.
	Location *time.Location

	// When set [MetaSet.MetaGetLoc] will allow timezone to be represented as a
	// sting.
	//
	// Example:
	//   Europe/Warsaw
	LocationAsString bool

	// String values considered as zero time value.
	//
	// List of strings representing zero time value.
	//
	// Examples:
	//   - 0001-01-01T00:00:00
	//   - 0000-00-00T00:00:00
	//   - 0000-00-00T00:00:00Z
	zeroTime []string

	// The base for integers when parsing.
	Radix int
}

// NewOptions returns a new [Options] instance with default values.
func NewOptions(opts ...Option) Options {
	o := Options{
		TimeFormat: time.RFC3339Nano,
		Radix:      10,
	}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

// WithLen is an option to set the default length for the map.
func WithLen(n int) Option {
	return func(opts *Options) { opts.Length = n }
}

// WithMeta is an option to set the initial map for the [MetaSet].
//
// The caller must not use the passed map after the call to this option. The
// [MetaSet] becomes its new owner.
func WithMeta(m map[string]any) Option {
	return func(opts *Options) { opts.init = m }
}

// WithTags is an option to set the initial map for the [TagSet].
//
// The caller must not use the passed map after the call to this option. The
// [TagSet] becomes its new owner.
func WithTags(m map[string]Tag) Option {
	return func(opts *Options) { opts.init = m }
}

// WithTimeFormat is the [MetaSet] option setting string time format.
func WithTimeFormat(format string) Option {
	return func(opts *Options) { opts.TimeFormat = format }
}

// WithTimeLoc is the [MetaSet] option setting location for parsed time strings.
func WithTimeLoc(loc *time.Location) Option {
	return func(opts *Options) { opts.Location = loc }
}

// WithLocString is the [MetaSet] option allowing string timezone names.
func WithLocString(opts *Options) { opts.LocationAsString = true }

// WithZeroTime is [MetaSet] option setting zero time values.
func WithZeroTime(zero ...string) Option {
	return func(opts *Options) { opts.zeroTime = zero }
}

// WithRadixHEX sets base to hexadecimal when parsing integers.
func WithRadixHEX(opts *Options) { opts.Radix = 16 }
