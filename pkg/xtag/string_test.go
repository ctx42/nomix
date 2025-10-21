// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/nomix/pkg/nomix"
)

func Test_StringSpec(t *testing.T) {
	// --- When ---
	have := StringSpec()

	// --- Then ---
	tag, err := have.TagCreate("name", "abc")
	assert.NoError(t, err)
	assert.SameType(t, &String{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, "abc", tag.TagValue())
	assert.Equal(t, nomix.KindString, tag.TagKind())

	tag, err = have.TagParse("name", "abc")
	assert.NoError(t, err)
	assert.SameType(t, &String{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, "abc", tag.TagValue())
	assert.Equal(t, nomix.KindString, tag.TagKind())
}

func Test_NewString(t *testing.T) {
	// --- When ---
	have := NewString("name", "abc")

	// --- Then ---
	assert.SameType(t, &String{}, have)
	assert.Equal(t, "name", have.TagName())
	assert.Equal(t, "abc", have.TagValue())
	assert.Equal(t, nomix.KindString, have.TagKind())
	assert.Equal(t, "abc", have.String())

	val, err := have.Value()
	assert.NoError(t, err)
	assert.Equal(t, "abc", val)
}

func Test_CreateString(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := CreateString("name", "abc")

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &String{}, have)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, "abc", have.TagValue())
		assert.Equal(t, nomix.KindString, have.TagKind())
		assert.Equal(t, "abc", have.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := CreateString("name", 42)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := CreateString("name", nil)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.Empty(t, have)
	})
}

func Test_strValueString(t *testing.T) {
	// --- When ---
	have := strValueString("abc")

	// --- Then ---
	assert.Equal(t, "abc", have)
}
