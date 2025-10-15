// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewStringSlice(t *testing.T) {
	// --- When ---
	tag := NewStringSlice("name", "abc", "xyz")

	// --- Then ---
	assert.SameType(t, &StringSlice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []string{"abc", "xyz"}, tag.value)
	assert.Equal(t, KindStringSlice, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, `["abc", "xyz"]`, tag.stringer([]string{"abc", "xyz"}))
}

func Test_asStringSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := asStringSlice([]string{"abc", "xyz"}, Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []string{"abc", "xyz"}, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asStringSlice(42, Options{})

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("nil value", func(t *testing.T) {
		// --- When ---
		have, err := asStringSlice(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})
}
