// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_IntSliceSpec(t *testing.T) {
	// --- When ---
	have := IntSliceSpec()

	// --- Then ---
	tag, err := have.TagCreate("name", []int{42, 44})
	assert.NoError(t, err)
	assert.SameType(t, &IntSlice{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []int{42, 44}, tag.TagValue())
	assert.Equal(t, KindIntSlice, tag.TagKind())

	tag, err = have.TagParse("name", "[42, 44]")
	assert.ErrorIs(t, ErrNotImpl, err)
	assert.Nil(t, tag)
}

func Test_NewIntSlice(t *testing.T) {
	// --- When ---
	tag := NewIntSlice("name", 42, 44)

	// --- Then ---
	assert.SameType(t, &IntSlice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []int{42, 44}, tag.value)
	assert.Equal(t, KindIntSlice, tag.kind)
	assert.Equal(t, "[42, 44]", tag.String())
}

func Test_CreateIntSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		tag, err := CreateIntSlice("name", []int{42, 44})

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &IntSlice{}, tag)
		assert.Equal(t, "name", tag.name)
		assert.Equal(t, []int{42, 44}, tag.value)
		assert.Equal(t, KindIntSlice, tag.kind)
		assert.Equal(t, "[42, 44]", tag.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		tag, err := CreateIntSlice("name", "abc")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, tag)
	})
}

func Test_createIntSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := createIntSlice([]int{42, 44}, Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []int{42, 44}, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := createIntSlice("abc", Options{})

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := createIntSlice(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})
}
