// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
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
	assert.Equal(t, KindFloat64Slice, tag.TagKind())

	tag, err = have.TagParse("name", "[42.1, 44.2]")
	assert.ErrorIs(t, ErrNotImpl, err)
	assert.Nil(t, tag)
}

func Test_NewFloat64Slice(t *testing.T) {
	// --- When ---
	tag := NewFloat64Slice("name", 42.1, 44.2)

	// --- Then ---
	assert.SameType(t, &Float64Slice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []float64{42.1, 44.2}, tag.value)
	assert.Equal(t, KindFloat64Slice, tag.kind)
	assert.Equal(t, "[42.1, 44.2]", tag.String())
}

func Test_CreateFloat64Slice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		tag, err := CreateFloat64Slice("name", []float64{42.1, 44.2})

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &Float64Slice{}, tag)
		assert.Equal(t, "name", tag.name)
		assert.Equal(t, []float64{42.1, 44.2}, tag.value)
		assert.Equal(t, KindFloat64Slice, tag.kind)
		assert.Equal(t, "[42.1, 44.2]", tag.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		tag, err := CreateFloat64Slice("name", "abc")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, tag)
	})
}

func Test_createFloat64Slice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := createFloat64Slice([]float64{42, 44}, Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []float64{42, 44}, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := createFloat64Slice("abc", Options{})

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := createFloat64Slice(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})
}

func Test_createFloat64Slice_success_tabular(t *testing.T) {
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
			have, err := createFloat64Slice(tc.have, Options{})

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}
