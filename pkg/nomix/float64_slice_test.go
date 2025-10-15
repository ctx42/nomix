// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewFloat64Slice(t *testing.T) {
	// --- When ---
	tag := NewFloat64Slice("name", 42.1, 44.2)

	// --- Then ---
	assert.SameType(t, &Float64Slice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []float64{42.1, 44.2}, tag.value)
	assert.Equal(t, KindFloat64Slice, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, "[42.1, 44.2]", tag.stringer([]float64{42.1, 44.2}))
}

func Test_asFloat64Slice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := asFloat64Slice([]float64{42, 44}, Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []float64{42, 44}, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asFloat64Slice("abc", Options{})

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("nil value", func(t *testing.T) {
		// --- When ---
		have, err := asFloat64Slice(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})
}

func Test_asFloat64Slice_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want []float64
	}{
		{"[]int", []int{42, 44}, []float64{42, 44}},
		{"[]int8", []int8{42, 44}, []float64{42, 44}},
		{"[]int16", []int16{42, 44}, []float64{42, 44}},
		{"[]int32", []int32{42, 44}, []float64{42, 44}},
		{"[]int64", []int64{42, 44}, []float64{42, 44}},
		{"[]float32", []float32{42, 44}, []float64{42, 44}},
		{"[]float64", []float64{42, 44}, []float64{42, 44}},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := asFloat64Slice(tc.have, Options{})

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}
