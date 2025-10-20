// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
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
	assert.Equal(t, KindString, tag.TagKind())

	tag, err = have.TagParse("name", "abc")
	assert.NoError(t, err)
	assert.SameType(t, &String{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, "abc", tag.TagValue())
	assert.Equal(t, KindString, tag.TagKind())
}

func Test_NewString(t *testing.T) {
	// --- When ---
	tag := NewString("name", "abc")

	// --- Then ---
	assert.SameType(t, &String{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, "abc", tag.value)
	assert.Equal(t, KindString, tag.kind)
	assert.Equal(t, "abc", tag.String())

	val, err := tag.Value()
	assert.NoError(t, err)
	assert.Equal(t, "abc", val)
}

func Test_CreateString(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		tag, err := CreateString("name", "abc")

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &String{}, tag)
		assert.Equal(t, "name", tag.name)
		assert.Equal(t, "abc", tag.value)
		assert.Equal(t, KindString, tag.kind)
		assert.Equal(t, "abc", tag.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		tag, err := CreateString("name", 42)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, tag)
	})
}

func Test_createString(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := createString("abc", Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "abc", have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := createString(42, Options{})

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := createString(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Empty(t, have)
	})
}

func Test_stringValueString(t *testing.T) {
	// --- When ---
	have := stringValueString("abc")

	// --- Then ---
	assert.Equal(t, "abc", have)
}
