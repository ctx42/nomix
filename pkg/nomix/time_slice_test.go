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
