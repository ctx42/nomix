// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac.
// SPDX-License-Identifier: MIT

package xtag

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/nomix/pkg/nomix"
)

func Test_Int64Spec(t *testing.T) {
	// --- When ---
	have := Int64Spec()

	// --- Then ---
	tag, err := have.TagCreate("name", 42)
	assert.NoError(t, err)
	assert.SameType(t, &Int64{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, int64(42), tag.TagValue())
	assert.Equal(t, nomix.KindInt64, tag.TagKind())

	tag, err = have.TagParse("name", "42")
	assert.NoError(t, err)
	assert.SameType(t, &Int64{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, int64(42), tag.TagValue())
	assert.Equal(t, nomix.KindInt64, tag.TagKind())
}

func Test_NewInt64(t *testing.T) {
	// --- When ---
	have := NewInt64("name", 42)

	// --- Then ---
	assert.SameType(t, &Int64{}, have)
	assert.Equal(t, "name", have.TagName())
	assert.Equal(t, int64(42), have.TagValue())
	assert.Equal(t, nomix.KindInt64, have.TagKind())
	assert.Equal(t, "42", have.String())

	val, err := have.Value()
	assert.NoError(t, err)
	assert.Equal(t, int64(42), val)
}

func Test_CreateInt64(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := CreateInt64("name", 42)

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &Int64{}, have)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, int64(42), have.TagValue())
		assert.Equal(t, nomix.KindInt64, have.TagKind())
		assert.Equal(t, "42", have.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := CreateInt64("name", "abc")

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := CreateInt64("name", nil)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.Nil(t, have)
	})
}

func Test_ParseInt64(t *testing.T) {
	t.Run("error - not supported string value", func(t *testing.T) {
		// --- When ---
		have, err := ParseInt64("name", "bad")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, nomix.ErrInvFormat, err)
		assert.Nil(t, have)
	})

	t.Run("error - hex without option", func(t *testing.T) {
		// --- When ---
		have, err := ParseInt64("name", "AA")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, nomix.ErrInvFormat, err)
		assert.Nil(t, have)
	})
}

func Test_ParseInt64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		str  string
		opts []nomix.Option
		want int64
	}{
		{"negative", "-1", nil, -1},
		{"zero", "0", nil, 0},
		{"positive", "1", nil, 1},
		{"hex", "AA", []nomix.Option{nomix.WithRadixHEX}, 170},
		{"negative hex", "-AA", []nomix.Option{nomix.WithRadixHEX}, -170},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			t.Parallel()

			// --- When ---
			have, err := ParseInt64("name", tc.str, tc.opts...)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have.TagValue())
		})
	}
}

func Test_strValueInt64(t *testing.T) {
	// --- When ---
	have := strValueInt64(42)

	// --- Then ---
	assert.Equal(t, "42", have)
}
