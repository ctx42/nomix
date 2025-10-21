// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/nomix/pkg/nomix"
)

func Test_BoolSpec(t *testing.T) {
	// --- When ---
	have := BoolSpec()

	// --- Then ---
	tag, err := have.TagCreate("name", true)
	assert.NoError(t, err)
	assert.SameType(t, &Bool{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, true, tag.TagValue())
	assert.Equal(t, nomix.KindBool, tag.TagKind())

	tag, err = have.TagParse("name", "true")
	assert.NoError(t, err)
	assert.SameType(t, &Bool{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, true, tag.TagValue())
	assert.Equal(t, nomix.KindBool, tag.TagKind())
}

func Test_NewBool(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		// --- When ---
		have := NewBool("name", true)

		// --- Then ---
		assert.SameType(t, &Bool{}, have)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, true, have.TagValue())
		assert.Equal(t, nomix.KindBool, have.TagKind())
		assert.Equal(t, "true", have.String())

		val, err := have.Value()
		assert.NoError(t, err)
		assert.Equal(t, int64(1), val)
	})

	t.Run("false", func(t *testing.T) {
		// --- When ---
		have := NewBool("name", false)

		// --- Then ---
		assert.SameType(t, &Bool{}, have)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, false, have.TagValue())
		assert.Equal(t, nomix.KindBool, have.TagKind())
		assert.Equal(t, "false", have.String())

		val, err := have.Value()
		assert.NoError(t, err)
		assert.Equal(t, int64(0), val)
	})
}

func Test_CreateBool(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := CreateBool("name", true)

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &Bool{}, have)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, true, have.TagValue())
		assert.Equal(t, nomix.KindBool, have.TagKind())
		assert.Equal(t, "true", have.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := CreateBool("name", 42)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := CreateBool("name", nil)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.Nil(t, have)
	})
}

func Test_ParseBool_success_tabular(t *testing.T) {
	tt := []struct {
		str string
		exp bool
	}{
		{"1", true},
		{"t", true},
		{"T", true},
		{"true", true},
		{"TRUE", true},
		{"True", true},
		{"0", false},
		{"f", false},
		{"F", false},
		{"false", false},
		{"FALSE", false},
		{"False", false},
	}

	for _, tc := range tt {
		t.Run(tc.str, func(t *testing.T) {
			// --- When ---
			have, err := ParseBool("name", tc.str)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.exp, have.TagValue())
		})
	}
}

func Test_ParseBool(t *testing.T) {
	t.Run("error - not supported string value", func(t *testing.T) {
		// --- When ---
		have, err := ParseBool("name", "bad")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, nomix.ErrInvFormat, err)
		assert.Nil(t, have)
	})
}

func Test_sqlValueBool_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have bool
		want any
	}{
		{"true", true, int64(1)},
		{"false", false, int64(0)},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := sqlValueBool(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}
