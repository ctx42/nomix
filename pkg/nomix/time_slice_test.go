// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_TimeSliceSpec(t *testing.T) {
	// --- When ---
	have := TimeSliceSpec()

	// --- Then ---
	tim0 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
	tim1 := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)
	tag, err := have.TagCreate("name", []time.Time{tim0, tim1})
	assert.NoError(t, err)
	assert.SameType(t, &TimeSlice{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []time.Time{tim0, tim1}, tag.TagValue())
	assert.Equal(t, KindTimeSlice, tag.TagKind())

	data := `["2000-01-02T03:04:05Z", "2001-01-02T03:04:05Z"]`
	tag, err = have.TagParse("name", data)
	assert.ErrorIs(t, ErrNotImpl, err)
	assert.Nil(t, tag)
}

func Test_NewTimeSlice(t *testing.T) {
	// --- Given ---
	tim0 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
	tim1 := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)

	// --- When ---
	tag := NewTimeSlice("name", tim0, tim1)

	// --- Then ---
	assert.SameType(t, &TimeSlice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []time.Time{tim0, tim1}, tag.value)
	assert.Equal(t, KindTimeSlice, tag.kind)
	have := tag.String()
	assert.Equal(t, `["2000-01-02T03:04:05Z", "2001-01-02T03:04:05Z"]`, have)
}

func Test_CreateTimeSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tim0 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		tim1 := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		tag, err := CreateTimeSlice("name", []time.Time{tim0, tim1})

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &TimeSlice{}, tag)
		assert.Equal(t, "name", tag.name)
		assert.Equal(t, []time.Time{tim0, tim1}, tag.value)
		assert.Equal(t, KindTimeSlice, tag.kind)
		have := tag.String()
		want := `["2000-01-02T03:04:05Z", "2001-01-02T03:04:05Z"]`
		assert.Equal(t, want, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		tag, err := CreateTimeSlice("name", 42)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, tag)
	})
}

func Test_createTimeSlice(t *testing.T) {
	t.Run("error - not supported type", func(t *testing.T) {
		// --- When ---
		have, err := createTimeSlice(42, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})

	t.Run("error - string type is not allowed by default", func(t *testing.T) {
		// --- When ---
		have, err := createTimeSlice("2000-01-02T03:04:05Z", Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})

	t.Run("error - parsing string time", func(t *testing.T) {
		// --- Given ---
		opts := Options{timeFormat: time.RFC3339}

		// --- When ---
		have, err := createTimeSlice([]string{"abc"}, opts)

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := createTimeSlice(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})
}

func Test_createTimeSlice_success_tabular(t *testing.T) {
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
			opt := Options{timeFormat: time.RFC3339}

			// --- When ---
			have, err := createTimeSlice(tc.have, opt)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}
