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

func Test_asFloat64(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := asFloat64(42, Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, 42.0, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asFloat64("abc", Options{})

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("nil value", func(t *testing.T) {
		// --- When ---
		have, err := asFloat64(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Equal(t, 0.0, have)
	})
}

func Test_asFloat64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want float64
	}{
		{"int", 42, 42},
		{"byte", byte(42), 42},
		{"int8", int8(42), 42},
		{"int16", int16(42), 42},
		{"int32", int32(42), 42},
		{"int64", int64(42), 42},
		{"float32", float32(42), 42},
		{"float64", 42.0, 42.0},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := asFloat64(tc.have, Options{})

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}
