// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/nomix/pkg/nomix"
)

func Test_Int64SliceSpec(t *testing.T) {
	// --- When ---
	have := Int64SliceSpec()

	// --- Then ---
	tag, err := have.TagCreate("name", []int64{42, 44})
	assert.NoError(t, err)
	assert.SameType(t, &Int64Slice{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []int64{42, 44}, tag.TagValue())
	assert.Equal(t, nomix.KindInt64Slice, tag.TagKind())

	tag, err = have.TagParse("name", "[42, 44]")
	assert.ErrorIs(t, nomix.ErrNotImpl, err)
	assert.Nil(t, tag)
}

func Test_NewInt64Slice(t *testing.T) {
	// --- When ---
	tag := NewInt64Slice("name", 42, 44)

	// --- Then ---
	assert.SameType(t, &Int64Slice{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []int64{42, 44}, tag.TagValue())
	assert.Equal(t, nomix.KindInt64Slice, tag.TagKind())
	assert.Equal(t, "[42, 44]", tag.String())
}

func Test_CreateInt64Slice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		tag, err := CreateInt64Slice("name", []int64{42, 44})

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &Int64Slice{}, tag)
		assert.Equal(t, "name", tag.TagName())
		assert.Equal(t, []int64{42, 44}, tag.TagValue())
		assert.Equal(t, nomix.KindInt64Slice, tag.TagKind())
		assert.Equal(t, "[42, 44]", tag.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		tag, err := CreateInt64Slice("name", "abc")

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, tag)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := CreateInt64Slice("name", nil)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.Nil(t, have)
	})
}

func Test_CreateInt64Slice_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want []int64
	}{
		{"[]int", []int{42, 44}, []int64{42, 44}},
		{"[]int8", []int8{42, 44}, []int64{42, 44}},
		{"[]int16", []int16{42, 44}, []int64{42, 44}},
		{"[]int32", []int32{42, 44}, []int64{42, 44}},
		{"[]int64", []int64{42, 44}, []int64{42, 44}},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := CreateInt64Slice("name", tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have.TagValue())
		})
	}
}

func Test_strValueInt64Slice(t *testing.T) {
	// --- Given ---
	i64s := []int64{42, 44}

	// --- When ---
	have := strValueInt64Slice(i64s)

	// --- Then ---
	assert.Equal(t, "[42, 44]", have)
}
