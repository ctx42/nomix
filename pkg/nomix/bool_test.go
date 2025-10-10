// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewBool(t *testing.T) {
	// --- When ---
	tag := NewBool("name", true)

	// --- Then ---
	assert.SameType(t, &Bool{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, true, tag.value)
	assert.Equal(t, KindBool, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, "false", tag.stringer(false))
}

func Test_ParseBool_success_tabular(t *testing.T) {
	tt := []struct {
		str string
		exp bool
	}{
		{"1", true},
		{"t", true},
		{"T", true},
		{"true", true},
		{"TRUE", true},
		{"True", true},
		{"0", false},
		{"f", false},
		{"F", false},
		{"false", false},
		{"FALSE", false},
		{"False", false},
	}

	for _, tc := range tt {
		t.Run(tc.str, func(t *testing.T) {
			t.Parallel()

			// --- When ---
			tag, err := ParseBool("name", tc.str)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.exp, tag.TagValue())
		})
	}
}

func Test_ParseBool_error(t *testing.T) {
	t.Run("error - not supported string value", func(t *testing.T) {
		// --- When ---
		tag, err := ParseBool("name", "bad")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, tag)
	})
}
