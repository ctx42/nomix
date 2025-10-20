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

func Test_TimeSpec(t *testing.T) {
	// --- When ---
	have := TimeSpec()

	// --- Then ---
	tim := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
	tag, err := have.TagCreate("name", tim)
	assert.NoError(t, err)
	assert.SameType(t, &Time{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Exact(t, tim, tag.TagValue())
	assert.Equal(t, KindTime, tag.TagKind())

	tag, err = have.TagParse("name", "2000-01-02T03:04:05Z")
	assert.NoError(t, err)
	assert.SameType(t, &Time{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Exact(t, tim, tag.TagValue())
	assert.Equal(t, KindTime, tag.TagKind())
}

func Test_NewTime(t *testing.T) {
	// --- When ---
	have := NewTime("name", time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC))

	// --- Then ---
	assert.SameType(t, &Time{}, have)
	assert.Equal(t, "name", have.name)
	assert.Exact(t, time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC), have.value)
	assert.Equal(t, KindTime, have.kind)
	assert.Exact(t, "2000-01-02T03:04:05Z", have.String())

	val, err := have.Value()
	assert.NoError(t, err)
	assert.Exact(t, time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC), val)
}

func Test_CreateTime(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tim := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		have, err := CreateTime("name", tim)

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &Time{}, have)
		assert.Equal(t, "name", have.name)
		assert.Exact(t, time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC), have.value)
		assert.Equal(t, KindTime, have.kind)
		assert.Exact(t, "2000-01-02T03:04:05Z", have.String())
	})

	t.Run("options passed", func(t *testing.T) {
		// --- Given ---
		opt := WithTimeFormat("2006-01-02")

		// --- When ---
		have, err := CreateTime("name", "2000-01-02", opt)

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &Time{}, have)
		assert.Equal(t, "name", have.name)
		assert.Exact(t, time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC), have.value)
		assert.Equal(t, KindTime, have.kind)
		assert.Exact(t, "2000-01-02T00:00:00Z", have.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := CreateTime("name", 42)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, have)
	})
}

func Test_createTime(t *testing.T) {
	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := createTime(42, Options{})

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("error - by default sting is not supported", func(t *testing.T) {
		// --- When ---
		have, err := createTime("2000-01-02T03:04:05Z", Options{})

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := createTime(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Zero(t, have)
	})
}

func Test_createTime_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want time.Time
	}{
		{
			"time.Time",
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
		},
		{
			"string",
			"2000-01-02T03:04:05Z",
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			opts := Options{timeFormat: time.RFC3339}

			// --- When ---
			have, err := createTime(tc.have, opts)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
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
			"format option passed to parser",
			"2022-03-04",
			[]Option{WithTimeFormat("2006-01-02")},
			time.Date(2022, 3, 4, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range tt {
		t.Run(tc.str, func(t *testing.T) {
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
		opts := Options{timeFormat: time.RFC3339}

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
		opts := Options{
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
		opts := Options{
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
		opts := Options{
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
		have, err := parseTime("2000-01-02T03:04:05Z", Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Zero(t, have)
	})

	t.Run("error - invalid time format", func(t *testing.T) {
		// --- Given ---
		opts := Options{timeFormat: time.RFC3339}

		// --- When ---
		have, err := parseTime("2000-01-02", opts)

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Zero(t, have)
	})

	t.Run("error - invalid time format with a loc option", func(t *testing.T) {
		// --- Given ---
		opts := Options{
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

func Test_sqlValueTime(t *testing.T) {
	// --- When ---
	have, err := sqlValueTime(time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC))

	// --- Then ---
	assert.NoError(t, err)
	assert.Exact(t, time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC), have)
}
