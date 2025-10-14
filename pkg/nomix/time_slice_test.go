// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewTimeSlice(t *testing.T) {
	// --- When ---
	tim0 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
	tim1 := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)
	tag := NewTimeSlice("name", []time.Time{tim0, tim1})

	// --- Then ---
	assert.SameType(t, &TimeSlice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []time.Time{tim0, tim1}, tag.value)
	assert.Equal(t, KindTimeSlice, tag.kind)
	assert.NotNil(t, tag.stringer)
	have := tag.stringer([]time.Time{tim0, tim1})
	assert.Equal(t, `["2000-01-02T03:04:05Z", "2001-01-02T03:04:05Z"]`, have)
}

func Test_asTimeSlice(t *testing.T) {
	t.Run("error - not supported type", func(t *testing.T) {
		// --- When ---
		have, err := asTimeSlice(42, nil)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})

	t.Run("error - string type is not allowed by default", func(t *testing.T) {
		// --- When ---
		have, err := asTimeSlice("2000-01-02T03:04:05Z", nil)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})

	t.Run("error - parsing string time", func(t *testing.T) {
		// --- Given ---
		opts := &Options{timeFormat: time.RFC3339}

		// --- When ---
		have, err := asTimeSlice([]string{"abc"}, opts)

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, have)
	})
}

func Test_asTimeSlice_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want []time.Time
	}{
		{
			"[]time.Time",
			[]time.Time{
				time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
				time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC),
			},
			[]time.Time{
				time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
				time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC),
			},
		},
		{
			"[]string",
			[]string{"2000-01-02T03:04:05Z", "2001-01-02T03:04:05Z"},
			[]time.Time{
				time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
				time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			opt := &Options{timeFormat: time.RFC3339}

			// --- When ---
			have, err := asTimeSlice(tc.have, opt)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}
