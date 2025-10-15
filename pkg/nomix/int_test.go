// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewInt(t *testing.T) {
	// --- When ---
	tag := NewInt("name", 42)

	// --- Then ---
	assert.SameType(t, &Int{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, 42, tag.value)
	assert.Equal(t, KindInt, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, "44", tag.stringer(44))
}

func Test_CreateInt(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		tag, err := CreateInt("name", 42)

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &Int{}, tag)
		assert.Equal(t, "name", tag.name)
		assert.Equal(t, 42, tag.value)
		assert.Equal(t, KindInt, tag.kind)
		assert.NotNil(t, tag.stringer)
		assert.Equal(t, "44", tag.stringer(44))
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		tag, err := CreateInt("name", "abc")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, tag)
	})
}

func Test_ParseInt_success_tabular(t *testing.T) {
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
			// --- When ---
			tag, err := ParseInt("name", tc.str, tc.opts...)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.exp, tag.value)
		})
	}
}

func Test_ParseInt(t *testing.T) {
	t.Run("error - not supported string value", func(t *testing.T) {
		// --- When ---
		tag, err := ParseInt("name", "bad")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, tag)
	})

	t.Run("error - hex without option", func(t *testing.T) {
		// --- When ---
		tag, err := ParseInt("name", "AA")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, tag)
	})
}

func Test_createInt(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := createInt(42, Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, 42, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := createInt("abc", Options{})

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := createInt(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Equal(t, 0, have)
	})
}
