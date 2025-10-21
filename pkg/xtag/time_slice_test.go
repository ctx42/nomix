// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/nomix/pkg/nomix"
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
	assert.Equal(t, nomix.KindTimeSlice, tag.TagKind())

	data := `["2000-01-02T03:04:05Z", "2001-01-02T03:04:05Z"]`
	tag, err = have.TagParse("name", data)
	assert.ErrorIs(t, nomix.ErrNotImpl, err)
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
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []time.Time{tim0, tim1}, tag.TagValue())
	assert.Equal(t, nomix.KindTimeSlice, tag.TagKind())
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
		assert.Equal(t, "name", tag.TagName())
		assert.Equal(t, []time.Time{tim0, tim1}, tag.TagValue())
		assert.Equal(t, nomix.KindTimeSlice, tag.TagKind())
		have := tag.String()
		want := `["2000-01-02T03:04:05Z", "2001-01-02T03:04:05Z"]`
		assert.Equal(t, want, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		tag, err := CreateTimeSlice("name", 42)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, tag)
	})
}
