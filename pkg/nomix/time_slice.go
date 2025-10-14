// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
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

// asTimeSlice casts the value to []time.Time or when the value is a []string
// parses its values as time but only when [Options.timeFormat] is set. Returns
// the slice and nil error on success. Returns nil and [ErrInvType] if the
// value is not a supported type.
func asTimeSlice(val any, opts Options) ([]time.Time, error) {
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
