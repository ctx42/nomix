// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_BoolSliceSpec(t *testing.T) {
	// --- When ---
	have := BoolSliceSpec()

	// --- Then ---
	tag, err := have.TagCreate("name", []bool{true, false})
	assert.NoError(t, err)
	assert.SameType(t, &BoolSlice{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []bool{true, false}, tag.TagValue())
	assert.Equal(t, KindBoolSlice, tag.TagKind())

	tag, err = have.TagParse("name", "[true, false]")
	assert.ErrorIs(t, ErrNotImpl, err)
	assert.Nil(t, tag)
}

func Test_NewBoolSlice(t *testing.T) {
	// --- When ---
	tag := NewBoolSlice("name", true, false)

	// --- Then ---
	assert.SameType(t, &BoolSlice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []bool{true, false}, tag.value)
	assert.Equal(t, KindBoolSlice, tag.kind)
	assert.Equal(t, "[true, false]", tag.String())
}

func Test_CreateBoolSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		tag, err := CreateBoolSlice("name", []bool{true, false})

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &BoolSlice{}, tag)
		assert.Equal(t, "name", tag.name)
		assert.Equal(t, []bool{true, false}, tag.value)
		assert.Equal(t, KindBoolSlice, tag.kind)
		assert.Equal(t, "[true, false]", tag.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		tag, err := CreateBoolSlice("name", 42)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, tag)
	})
}

func Test_createBoolSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := createBoolSlice([]bool{true, false}, Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []bool{true, false}, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := createBoolSlice(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := createBoolSlice(42, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})
}
