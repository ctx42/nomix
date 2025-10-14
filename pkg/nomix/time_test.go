// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/check"
	"github.com/ctx42/testing/pkg/must"
)

func Test_NewTime(t *testing.T) {
	// --- When ---
	tag := NewTime("name", time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC))

	// --- Then ---
	assert.SameType(t, &Time{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Exact(t, time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC), tag.value)
	assert.Equal(t, KindTime, tag.kind)
	assert.NotNil(t, tag.stringer)

	tim := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, "2000-01-01T00:00:00Z", tag.stringer(tim))
}

func Test_ParseTime_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		str  string
		opts []Option
		exp  time.Time
	}{
		{
			"RFC3339",
			"2022-03-04T05:06:07Z",
			nil,
			time.Date(2022, 3, 4, 5, 6, 7, 0, time.UTC),
		},
		{
			"date",
			"2022-03-04",
			[]Option{WithTimeFormat("2006-01-02")},
			time.Date(2022, 3, 4, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range tt {
		t.Run(tc.str, func(t *testing.T) {
			t.Parallel()

			// --- When ---
			have, err := ParseTime("name", tc.str, tc.opts...)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.exp, have.TagValue())
		})
	}
}

func Test_ParseTime(t *testing.T) {
	t.Run("error - invalid time format", func(t *testing.T) {
		// --- When ---
		have, err := ParseTime("name", "2022-03-04")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, have)
	})
}

func Test_parseTime(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		opts := &Options{
			timeFormat: time.RFC3339,
		}

		// --- When ---
		have, err := parseTime("2000-01-02T03:04:05Z", opts)

		// --- Then ---
		assert.NoError(t, err)
		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		exactTime := check.WithTypeChecker(time.Time{}, check.Exact)
		assert.Equal(t, want, have, exactTime)
	})

	t.Run("parse in location", func(t *testing.T) {
		// --- Given ---
		WAW := must.Value(time.LoadLocation("Europe/Warsaw"))
		opts := &Options{
			timeFormat: "2006-01-02",
			location:   WAW,
		}

		// --- When ---
		have, err := parseTime("2000-01-02", opts)

		// --- Then ---
		assert.NoError(t, err)
		want := time.Date(2000, 1, 2, 0, 0, 0, 0, WAW)
		exactTime := check.WithTypeChecker(time.Time{}, check.Exact)
		assert.Equal(t, want, have, exactTime)
	})

	t.Run("recognize zero value time", func(t *testing.T) {
		// --- Given ---
		opts := &Options{
			timeFormat: "2006-01-02",
			zeroTime:   []string{"0000-00-00"},
		}

		// --- When ---
		have, err := parseTime("0000-00-00", opts)

		// --- Then ---
		assert.NoError(t, err)
		assert.Zero(t, have)
	})

	t.Run("use UTC if the location is an empty string", func(t *testing.T) {
		// --- Given ---
		opts := &Options{
			timeFormat: "2006-01-02",
			location:   time.FixedZone("", 120),
		}

		// --- When ---
		have, err := parseTime("2000-01-02", opts)

		// --- Then ---
		assert.NoError(t, err)
		want := time.Date(2000, 1, 1, 23, 58, 0, 0, time.UTC)
		exactTime := check.WithTypeChecker(time.Time{}, check.Exact)
		assert.Equal(t, want, have, exactTime)
	})

	t.Run("error - time format isn't set", func(t *testing.T) {
		// --- When ---
		have, err := parseTime("2000-01-02T03:04:05Z", nil)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Zero(t, have)
	})

	t.Run("error - invalid time format", func(t *testing.T) {
		// --- Given ---
		opts := &Options{
			timeFormat: time.RFC3339,
		}

		// --- When ---
		have, err := parseTime("2000-01-02", opts)

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Zero(t, have)
	})

	t.Run("error - invalid time format with a loc option", func(t *testing.T) {
		// --- Given ---
		opts := &Options{
			timeFormat: time.RFC3339,
			location:   time.UTC,
		}

		// --- When ---
		have, err := parseTime("2000-01-02", opts)

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Zero(t, have)
	})
}
