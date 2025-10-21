// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/nomix/pkg/nomix"
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
	assert.Equal(t, nomix.KindBoolSlice, tag.TagKind())

	tag, err = have.TagParse("name", "[true, false]")
	assert.ErrorIs(t, nomix.ErrNotImpl, err)
	assert.Nil(t, tag)
}

func Test_NewBoolSlice(t *testing.T) {
	// --- When ---
	tag := NewBoolSlice("name", true, false)

	// --- Then ---
	assert.SameType(t, &BoolSlice{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []bool{true, false}, tag.TagValue())
	assert.Equal(t, nomix.KindBoolSlice, tag.TagKind())
	assert.Equal(t, "[true, false]", tag.String())
}

func Test_CreateBoolSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := CreateBoolSlice("name", []bool{true, false})

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &BoolSlice{}, have)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, []bool{true, false}, have.TagValue())
		assert.Equal(t, nomix.KindBoolSlice, have.TagKind())
		assert.Equal(t, "[true, false]", have.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := CreateBoolSlice("name", 42)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := CreateBoolSlice("name", nil)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.Nil(t, have)
	})
}
