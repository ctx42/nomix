// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/nomix/pkg/nomix"
)

func Test_Float64SliceSpec(t *testing.T) {
	// --- When ---
	have := Float64SliceSpec()

	// --- Then ---
	tag, err := have.TagCreate("name", []float64{42.1, 44.2})
	assert.NoError(t, err)
	assert.SameType(t, &Float64Slice{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []float64{42.1, 44.2}, tag.TagValue())
	assert.Equal(t, nomix.KindFloat64Slice, tag.TagKind())

	tag, err = have.TagParse("name", "[42.1, 44.2]")
	assert.ErrorIs(t, nomix.ErrNotImpl, err)
	assert.Nil(t, tag)
}

func Test_NewFloat64Slice(t *testing.T) {
	// --- When ---
	have := NewFloat64Slice("name", 42.1, 44.2)

	// --- Then ---
	assert.SameType(t, &Float64Slice{}, have)
	assert.Equal(t, "name", have.TagName())
	assert.Equal(t, []float64{42.1, 44.2}, have.TagValue())
	assert.Equal(t, nomix.KindFloat64Slice, have.TagKind())
	assert.Equal(t, "[42.1, 44.2]", have.String())
}

func Test_CreateFloat64Slice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := CreateFloat64Slice("name", []float64{42.1, 44.2})

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &Float64Slice{}, have)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, []float64{42.1, 44.2}, have.TagValue())
		assert.Equal(t, nomix.KindFloat64Slice, have.TagKind())
		assert.Equal(t, "[42.1, 44.2]", have.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := CreateFloat64Slice("name", "abc")

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := CreateFloat64Slice("name", nil)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.Nil(t, have)
	})
}
