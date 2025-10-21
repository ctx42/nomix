// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/nomix/pkg/nomix"
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
	assert.Equal(t, nomix.KindByteSlice, tag.TagKind())

	tag, err = have.TagParse("name", "[42, 44]")
	assert.ErrorIs(t, nomix.ErrNotImpl, err)
	assert.Nil(t, tag)
}

func Test_NewByteSlice(t *testing.T) {
	// --- When ---
	have := NewByteSlice("name", 42, 44)

	// --- Then ---
	assert.SameType(t, &ByteSlice{}, have)
	assert.Equal(t, "name", have.TagName())
	assert.Equal(t, []byte{42, 44}, have.TagValue())
	assert.Equal(t, nomix.KindByteSlice, have.TagKind())
	assert.Equal(t, "[42, 44]", have.String())
}

func Test_CreateByteSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := CreateByteSlice("name", []byte{42, 44})

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &ByteSlice{}, have)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, []byte{42, 44}, have.TagValue())
		assert.Equal(t, nomix.KindByteSlice, have.TagKind())
		assert.Equal(t, "[42, 44]", have.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := CreateByteSlice("name", 42)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := CreateByteSlice("name", nil)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.Nil(t, have)
	})
}
