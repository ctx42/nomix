// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/nomix/pkg/nomix"
)

func Test_Float64Spec(t *testing.T) {
	// --- When ---
	have := Float64Spec()

	// --- Then ---
	tag, err := have.TagCreate("name", 42)
	assert.NoError(t, err)
	assert.SameType(t, &Float64{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, float64(42), tag.TagValue())
	assert.Equal(t, nomix.KindFloat64, tag.TagKind())

	tag, err = have.TagParse("name", "42")
	assert.NoError(t, err)
	assert.SameType(t, &Float64{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, float64(42), tag.TagValue())
	assert.Equal(t, nomix.KindFloat64, tag.TagKind())
}

func Test_NewFloat64(t *testing.T) {
	// --- When ---
	have := NewFloat64("name", 4.2)

	// --- Then ---
	assert.Equal(t, "name", have.TagName())
	assert.Equal(t, 4.2, have.TagValue())
	assert.Equal(t, nomix.KindFloat64, have.TagKind())
	assert.Equal(t, "4.2", have.String())
}

func Test_CreateFloat64(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := CreateFloat64("name", 4.2)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, 4.2, have.TagValue())
		assert.Equal(t, nomix.KindFloat64, have.TagKind())
		assert.Equal(t, "4.2", have.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := CreateFloat64("name", "abc")

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := CreateFloat64("name", nil)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, have)
	})
}

func Test_ParseFloat64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		str string
		exp float64
	}{
		{"negative", "-1.1", -1.1},
		{"zero", "0", 0.0},
		{"positive", "1.1", 1.1},
		{"no fraction", "4", 4.0},
	}

	for _, tc := range tt {
		t.Run(tc.str, func(t *testing.T) {
			// --- When ---
			have, err := ParseFloat64("name", tc.str)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.exp, have.TagValue())
		})
	}
}

func Test_ParseFloat64(t *testing.T) {
	t.Run("error - not supported string value", func(t *testing.T) {
		// --- When ---
		have, err := ParseFloat64("name", "bad")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, nomix.ErrInvFormat, err)
		assert.Nil(t, have)
	})
}
