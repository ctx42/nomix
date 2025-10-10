// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
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
			tag, err := ParseTime("name", tc.str, tc.opts...)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.exp, tag.TagValue())
		})
	}
}

func Test_ParseTime(t *testing.T) {
	t.Run("error - not supported string value", func(t *testing.T) {
		// --- When ---
		tag, err := ParseTime("name", "bad")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, tag)
	})

	t.Run("error - invalid time format", func(t *testing.T) {
		// --- When ---
		tag, err := ParseTime("name", "2022-03-04")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, tag)
	})
}
