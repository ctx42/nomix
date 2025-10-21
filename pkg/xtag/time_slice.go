// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"fmt"
	"time"

	"github.com/ctx42/nomix/pkg/nomix"
)

// TimeSlice is a tag for a slice of [time.Time] values.
type TimeSlice = nomix.Slice[time.Time]

// timeSliceSpec defines the [nomix.Spec] for [TimeSlice] type.
var timeSliceSpec = nomix.NewSpec(
	nomix.KindTimeSlice,
	nomix.TagCreateFunc(CreateTimeSlice),
	nomix.TagParserNotImpl,
)

// TimeSliceSpec returns a [nomix.Spec] for [TimeSlice] type.
func TimeSliceSpec() nomix.Spec { return timeSliceSpec }

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
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += `"` + val.Format(time.RFC3339Nano) + `"`
	}
	return ret + "]"
}
