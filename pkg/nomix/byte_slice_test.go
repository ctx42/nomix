// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_ByteSliceSpec(t *testing.T) {
	// --- When ---
	have := ByteSliceSpec()

	// --- Then ---
	tag, err := have.TagCreate("name", []byte{42, 44})
	assert.NoError(t, err)
	assert.SameType(t, &ByteSlice{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []byte{42, 44}, tag.TagValue())
	assert.Equal(t, KindByteSlice, tag.TagKind())

	tag, err = have.TagParse("name", "[42, 44]")
	assert.ErrorIs(t, ErrNotImpl, err)
	assert.Nil(t, tag)
}

func Test_NewByteSlice(t *testing.T) {
	// --- When ---
	tag := NewByteSlice("name", 42, 44)

	// --- Then ---
	assert.SameType(t, &ByteSlice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []byte{42, 44}, tag.value)
	assert.Equal(t, KindByteSlice, tag.kind)
	assert.Equal(t, "[42, 44]", tag.String())
}

func Test_CreateByteSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		tag, err := CreateByteSlice("name", []byte{42, 44})

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &ByteSlice{}, tag)
		assert.Equal(t, "name", tag.name)
		assert.Equal(t, []byte{42, 44}, tag.value)
		assert.Equal(t, KindByteSlice, tag.kind)
		assert.Equal(t, "[42, 44]", tag.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		tag, err := CreateByteSlice("name", 42)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, tag)
	})
}

func Test_createByteSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := createByteSlice([]byte{0, 1, 2}, Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []byte{0, 1, 2}, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := createByteSlice(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := createByteSlice(42, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})
}
