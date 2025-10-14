// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewInt64Slice(t *testing.T) {
	// --- When ---
	tag := NewInt64Slice("name", 42, 44)

	// --- Then ---
	assert.SameType(t, &Int64Slice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []int64{42, 44}, tag.value)
	assert.Equal(t, KindInt64Slice, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, "[42, 44]", tag.stringer([]int64{42, 44}))
}

func Test_toInt64Slice(t *testing.T) {
	// --- Given ---
	i32s := []int32{42, 44}

	// --- When ---
	have := toInt64Slice(i32s, Options{})

	// --- Then ---
	assert.Equal(t, []int64{42, 44}, have)
}

func Test_asInt64Slice(t *testing.T) {
	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asInt64Slice("abc", Options{})

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})
}

func Test_asInt64Slice_success_tabular(t *testing.T) {
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
			have, err := asInt64Slice(tc.have, Options{})

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}
