// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac.
// SPDX-License-Identifier: MIT

package xtag

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ctx42/nomix/pkg/nomix"
)

// TimeSlice is a tag for a slice of [time.Time] values.
//
// Its SQL value is the []time.Time itself, so storing it requires a database
// driver that accepts array arguments.
type TimeSlice = nomix.Slice[time.Time]

// timeSliceSpec defines the [nomix.KindSpec] for [TimeSlice] type.
var timeSliceSpec = nomix.NewKindSpec(
	nomix.KindTimeSlice,
	nomix.TagCreateFunc(CreateTimeSlice),
	nomix.TagParserNotImpl,
)

// TimeSliceSpec returns a [nomix.KindSpec] for [TimeSlice] type.
func TimeSliceSpec() nomix.KindSpec { return timeSliceSpec }

// NewTimeSlice returns a new instance of [TimeSlice].
func NewTimeSlice(name string, val ...time.Time) *TimeSlice {
	return nomix.NewSlice(
		name,
		val,
		nomix.KindTimeSlice,
		strValueTimeSlice,
		nil,
	)
}

// CreateTimeSlice casts the value to []time.Time, or when the value is a
// []string, it parses its elements as [time.RFC3339Nano] time. Returns the
// [TimeSlice] instance with the given name and nil error on success. Returns
// nil and error if the value's type is not a supported type or the value is
// not a valid time representation.
func CreateTimeSlice(
	name string,
	val any,
	opts ...nomix.Option,
) (*TimeSlice, error) {

	def := nomix.NewOptions(opts...)
	v, err := nomix.CreateTimeSlice(val, def)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewTimeSlice(name, v...), nil
}

// strValueTimeSlice converts a [time.Time] slice to its string representation.
func strValueTimeSlice(v []time.Time) string {
	return formatSlice(v, quoteTime)
}

// quoteTime formats the time as a quoted [time.RFC3339Nano] string.
func quoteTime(v time.Time) string { return strconv.Quote(strValueTime(v)) }
