// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewInt64(t *testing.T) {
	// --- When ---
	tag := NewInt64("name", 42)

	// --- Then ---
	assert.SameType(t, &Int64{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, int64(42), tag.value)
	assert.Equal(t, KindInt64, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, "44", tag.stringer(44))
}

func Test_ParseInt64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		str  string
		opts []Option
		exp  int64
	}{
		{"negative", "-1", nil, -1},
		{"zero", "0", nil, 0},
		{"positive", "1", nil, 1},
		{"hex", "AA", []Option{WithBaseHEX}, 170},
		{"negative hex", "-AA", []Option{WithBaseHEX}, -170},
	}

	for _, tc := range tt {
		t.Run(tc.str, func(t *testing.T) {
			t.Parallel()

			// --- When ---
			tag, err := ParseInt64("name", tc.str, tc.opts...)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.exp, tag.TagValue())
		})
	}
}

func Test_ParseInt64_error(t *testing.T) {
	t.Run("error - not supported string value", func(t *testing.T) {
		// --- When ---
		tag, err := ParseInt64("name", "bad")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, tag)
	})

	t.Run("error - hex without option", func(t *testing.T) {
		// --- When ---
		tag, err := ParseInt64("name", "AA")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, tag)
	})
}

func Test_asInt64(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := asInt64(42, Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, int64(42), have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asInt64("abc", Options{})

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("nil value", func(t *testing.T) {
		// --- When ---
		have, err := asInt64(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Equal(t, int64(0), have)
	})
}

func Test_asInt64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want int64
	}{
		{"int", 42, 42},
		{"byte", byte(42), 42},
		{"int8", int8(42), 42},
		{"int16", int16(42), 42},
		{"int32", int32(42), 42},
		{"int64", int64(42), 42},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := asInt64(tc.have, Options{})

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}
