// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewByteSlice(t *testing.T) {
	// --- When ---
	tag := NewByteSlice("name", 42, 44)

	// --- Then ---
	assert.SameType(t, &ByteSlice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []byte{42, 44}, tag.value)
	assert.Equal(t, KindByteSlice, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, "[42, 44]", tag.stringer([]byte{42, 44}))
}

func Test_asByteSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := asByteSlice([]byte{0, 1, 2}, Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []byte{0, 1, 2}, have)
	})

	t.Run("nil value", func(t *testing.T) {
		// --- When ---
		have, err := asByteSlice(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asByteSlice(42, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})
}
