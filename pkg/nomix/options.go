// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac.
// SPDX-License-Identifier: MIT

package nomix

import (
	"time"
)

// Option represents an option function.
type Option func(*Options)

// Options represents a set of options used when creating sets ([NewTagSet],
// [NewMetaSet]) and when creating or parsing tag values (for example
// [CreateTime], [ParseTime], and [CreateTimeSlice]).
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
	// When set, [CreateTime] and [CreateTimeSlice] accept times represented
	// as strings, and [ParseTime] parses strings with this layout.
	TimeFormat string

	// Location to parse the time format.
	//
	// When set, [ParseTime] uses [time.ParseInLocation] instead of
	// [time.Parse] to parse strings.
	Location *time.Location

	// Deprecated: LocationAsString has no effect and will be removed.
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

// WithTimeFormat is an option setting the layout used to parse time strings.
func WithTimeFormat(format string) Option {
	return func(opts *Options) { opts.TimeFormat = format }
}

// WithTimeLoc is an option setting the location for parsed time strings.
func WithTimeLoc(loc *time.Location) Option {
	return func(opts *Options) { opts.Location = loc }
}

// WithLocString sets [Options.LocationAsString].
//
// Deprecated: WithLocString has no effect and will be removed.
func WithLocString(opts *Options) { opts.LocationAsString = true }

// WithZeroTime is an option setting the strings parsed as the zero time.
func WithZeroTime(zero ...string) Option {
	return func(opts *Options) { opts.zeroTime = zero }
}

// WithRadixHEX sets base to hexadecimal when parsing integers.
func WithRadixHEX(opts *Options) { opts.Radix = 16 }
