// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/check"

	"github.com/ctx42/nomix/pkg/nomix"
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
	assert.Equal(t, nomix.KindTime, tag.TagKind())

	tag, err = have.TagParse("name", "2000-01-02T03:04:05Z")
	assert.NoError(t, err)
	assert.SameType(t, &Time{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Exact(t, tim, tag.TagValue())
	assert.Equal(t, nomix.KindTime, tag.TagKind())
}

func Test_NewTime(t *testing.T) {
	// --- When ---
	have := NewTime("name", time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC))

	// --- Then ---
	assert.SameType(t, &Time{}, have)
	assert.Equal(t, "name", have.TagName())
	assert.Exact(t, "2000-01-02T03:04:05Z", have.TagValue())
	assert.Equal(t, nomix.KindTime, have.TagKind())
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
		assert.Equal(t, "name", have.TagName())
		assert.Exact(t, "2000-01-02T03:04:05Z", have.TagValue())
		assert.Equal(t, nomix.KindTime, have.TagKind())
		assert.Exact(t, "2000-01-02T03:04:05Z", have.String())
	})

	t.Run("options passed", func(t *testing.T) {
		// --- Given ---
		opt := nomix.WithTimeFormat("2006-01-02")

		// --- When ---
		have, err := CreateTime("name", "2000-01-02", opt)

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &Time{}, have)
		assert.Equal(t, "name", have.TagName())
		assert.Exact(t, "2000-01-02T00:00:00Z", have.TagValue())
		assert.Equal(t, nomix.KindTime, have.TagKind())
		assert.Exact(t, "2000-01-02T00:00:00Z", have.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := CreateTime("name", 42)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, have)
	})
}

func Test_ParseTime_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		str  string
		opts []nomix.Option
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
			[]nomix.Option{nomix.WithTimeFormat("2006-01-02")},
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
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		opt := nomix.WithTimeFormat(time.RFC3339)

		// --- When ---
		have, err := ParseTime("name", "2000-01-02T03:04:05Z", opt)

		// --- Then ---
		assert.NoError(t, err)
		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		exactTime := check.WithTypeChecker(time.Time{}, check.Exact)
		assert.Equal(t, want, have.TagValue(), exactTime)
	})

	t.Run("recognize zero value time", func(t *testing.T) {
		// --- Given ---
		opts := []nomix.Option{
			nomix.WithTimeFormat("2006-01-02"),
			nomix.WithZeroTime("0000-00-00"),
		}

		// --- When ---
		have, err := ParseTime("name", "0000-00-00", opts...)

		// --- Then ---
		assert.NoError(t, err)
		assert.Zero(t, have.TagValue())
	})

	t.Run("error - invalid time format", func(t *testing.T) {
		// --- When ---
		have, err := ParseTime("name", "2022-03-04")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, nomix.ErrInvFormat, err)
		assert.Nil(t, have)
	})

	t.Run("error - time format isn't set", func(t *testing.T) {
		// --- Given ---
		opt := nomix.WithTimeFormat("")

		// --- When ---
		have, err := ParseTime("name", "2000-01-02T03:04:05Z", opt)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.Nil(t, have)
	})
}

func Test_sqlValueTime(t *testing.T) {
	// --- When ---
	have, err := sqlValueTime(time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC))

	// --- Then ---
	assert.NoError(t, err)
	assert.Exact(t, time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC), have)
}
