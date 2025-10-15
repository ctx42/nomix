// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"time"
)

// TimeSlice is a tag for a slice of [time.Time] values.
type TimeSlice = slice[time.Time]

// NewTimeSlice returns a new instance of [TimeSlice].
func NewTimeSlice(name string, v []time.Time) *TimeSlice {
	return &slice[time.Time]{
		name:     name,
		value:    v,
		kind:     KindTimeSlice,
		stringer: timeSliceToString,
	}
}

// CreateTimeSlice casts the value to []time.Time, or when the value is a
// []string, it parses its elements as [time.RFC3339Nano] time. Returns the
// [TImeSlice] instance with the given name and nil error on success. Returns
// nil and error if the value's type is not a supported type or the value is
// not a valid time representation.
func CreateTimeSlice(name string, val any, _ ...Option) (*TimeSlice, error) {
	v, err := createTimeSlice(val, defaultOptions)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewTimeSlice(name, v), nil
}

// timeSliceToString converts a [time.Time] slice to its string representation.
func timeSliceToString(v []time.Time) string {
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += `"` + val.Format(time.RFC3339Nano) + `"`
	}
	return ret + "]"
}

// createTimeSlice casts the value to []time.Time, or when the value is a
// []string, it parses its elements as time but only when [Options.timeFormat]
// is set. Returns the []time.Time and nil error on success. Returns nil and
// error if the value's type is not a supported type or the value is not a
// valid time representation.
func createTimeSlice(val any, opts Options) ([]time.Time, error) {
	switch v := val.(type) {
	case []time.Time:
		return v, nil
	case []string:
		if opts.timeFormat != "" {
			times := make([]time.Time, len(v))
			for i, str := range v {
				var err error
				if times[i], err = parseTime(str, opts); err != nil {
					return nil, err
				}
			}
			return times, nil
		}
	}
	return nil, ErrInvType
}
