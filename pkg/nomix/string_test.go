// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewString(t *testing.T) {
	// --- When ---
	tag := NewString("name", "abc")

	// --- Then ---
	assert.SameType(t, &String{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, "abc", tag.value)
	assert.Equal(t, KindString, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, "44", tag.stringer("44"))
}

func Test_ParseString_tabular(t *testing.T) {
	tt := []struct {
		testN string

		str string
	}{
		{"string", "abc"},
	}

	for _, tc := range tt {
		t.Run(tc.str, func(t *testing.T) {
			// --- When ---
			tag, err := ParseString("name", tc.str)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, "abc", tag.TagValue())
		})
	}
}

func Test_asString(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := asString("abc", Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "abc", have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asString(42, Options{})

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("nil value", func(t *testing.T) {
		// --- When ---
		have, err := asString(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Empty(t, have)
	})
}
