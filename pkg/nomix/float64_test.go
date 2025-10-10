// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewFloat64(t *testing.T) {
	// --- When ---
	tag := NewFloat64("name", 4.2)

	// --- Then ---
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, 4.2, tag.value)
	assert.Equal(t, KindFloat64, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, "4.4", tag.stringer(4.4))
}

func Test_ParseFloat64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		str string
		exp float64
	}{
		{"negative", "-1.1", -1.1},
		{"zero", "0", 0.0},
		{"positive", "1.1", 1.1},
		{"no fraction", "4", 4.0},
	}

	for _, tc := range tt {
		t.Run(tc.str, func(t *testing.T) {
			t.Parallel()

			// --- When ---
			tag, err := ParseFloat64("name", tc.str)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.exp, tag.TagValue())
		})
	}
}

func Test_ParseFloat64(t *testing.T) {
	t.Run("error - not supported string value", func(t *testing.T) {
		// --- When ---
		tag, err := ParseFloat64("name", "bad")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, tag)
	})
}
