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
