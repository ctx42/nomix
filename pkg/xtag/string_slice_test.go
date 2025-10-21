// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/nomix/pkg/nomix"
)

func Test_StringSliceSpec(t *testing.T) {
	// --- When ---
	have := StringSliceSpec()

	// --- Then ---
	tag, err := have.TagCreate("name", []string{"abc", "xyz"})
	assert.NoError(t, err)
	assert.SameType(t, &StringSlice{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []string{"abc", "xyz"}, tag.TagValue())
	assert.Equal(t, nomix.KindStringSlice, tag.TagKind())

	tag, err = have.TagParse("name", `["abc", "xyz"]`)
	assert.ErrorIs(t, nomix.ErrNotImpl, err)
	assert.Nil(t, tag)
}

func Test_NewStringSlice(t *testing.T) {
	// --- When ---
	tag := NewStringSlice("name", "abc", "xyz")

	// --- Then ---
	assert.SameType(t, &StringSlice{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []string{"abc", "xyz"}, tag.TagValue())
	assert.Equal(t, nomix.KindStringSlice, tag.TagKind())
	assert.Equal(t, `["abc", "xyz"]`, tag.String())
}

func Test_CreateStringSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		tag, err := CreateStringSlice("name", []string{"abc", "xyz"})

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &StringSlice{}, tag)
		assert.Equal(t, "name", tag.TagName())
		assert.Equal(t, []string{"abc", "xyz"}, tag.TagValue())
		assert.Equal(t, nomix.KindStringSlice, tag.TagKind())
		assert.Equal(t, `["abc", "xyz"]`, tag.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		tag, err := CreateStringSlice("name", 42)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, tag)
	})
}

func Test_createStringSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := createStringSlice([]string{"abc", "xyz"}, nomix.Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []string{"abc", "xyz"}, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := createStringSlice(42, nomix.Options{})

		// --- Then ---
		assert.ErrorIs(t, err, nomix.ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := createStringSlice(nil, nomix.Options{})

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.Nil(t, have)
	})
}
