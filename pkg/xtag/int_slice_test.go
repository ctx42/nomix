// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/nomix/pkg/nomix"
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
	assert.Equal(t, nomix.KindIntSlice, tag.TagKind())

	tag, err = have.TagParse("name", "[42, 44]")
	assert.ErrorIs(t, nomix.ErrNotImpl, err)
	assert.Nil(t, tag)
}

func Test_NewIntSlice(t *testing.T) {
	// --- When ---
	tag := NewIntSlice("name", 42, 44)

	// --- Then ---
	assert.SameType(t, &IntSlice{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []int{42, 44}, tag.TagValue())
	assert.Equal(t, nomix.KindIntSlice, tag.TagKind())
	assert.Equal(t, "[42, 44]", tag.String())
}

func Test_CreateIntSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		tag, err := CreateIntSlice("name", []int{42, 44})

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &IntSlice{}, tag)
		assert.Equal(t, "name", tag.TagName())
		assert.Equal(t, []int{42, 44}, tag.TagValue())
		assert.Equal(t, nomix.KindIntSlice, tag.TagKind())
		assert.Equal(t, "[42, 44]", tag.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		tag, err := CreateIntSlice("name", "abc")

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, tag)
	})
}

func Test_createIntSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := createIntSlice([]int{42, 44}, nomix.Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []int{42, 44}, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := createIntSlice("abc", nomix.Options{})

		// --- Then ---
		assert.ErrorIs(t, err, nomix.ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := createIntSlice(nil, nomix.Options{})

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.Nil(t, have)
	})
}
