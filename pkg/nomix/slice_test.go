// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_array_TagName(t *testing.T) {
	tag := &slice[int]{name: "name"}

	// --- When ---
	have := tag.TagName()

	// --- Then ---
	assert.Equal(t, "name", have)
}

func Test_array_TagKind(t *testing.T) {
	tag := &slice[int]{kind: KindIntSlice}

	// --- When ---
	have := tag.TagKind()

	// --- Then ---
	assert.Equal(t, KindIntSlice, have)
}

func Test_array_TagValue(t *testing.T) {
	tag := &slice[int]{value: []int{42, 44}}

	// --- When ---
	have := tag.TagValue()

	// --- Then ---
	assert.Equal(t, []int{42, 44}, have)
}

func Test_array_Set(t *testing.T) {
	// --- Given ---
	tag := &slice[int]{value: []int{42, 44}}

	// --- When ---
	tag.Set([]int{52, 54})

	// --- Then ---
	assert.Equal(t, []int{52, 54}, tag.value)
}

func Test_array_TagSet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tag := &slice[int]{value: []int{42, 44}}

		// --- When ---
		err := tag.TagSet([]int{52, 54})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []int{52, 54}, tag.value)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- Given ---
		tag := &slice[int]{value: []int{42, 44}}

		// --- When ---
		err := tag.TagSet(true)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Equal(t, []int{42, 44}, tag.value)
	})
}

func Test_array_TagEqual(t *testing.T) {
	tt := []struct {
		testN string

		tag0 TagValueComparer
		tag1 Tag
		want bool
	}{
		{
			"same",
			&slice[int]{value: []int{42, 44}},
			&slice[int]{value: []int{42, 44}},
			true,
		},
		{
			"diff values",
			&slice[int]{value: []int{42, 44}},
			&slice[int]{value: []int{52, 45}},
			false,
		},
		{
			"diff value",
			&slice[int]{value: []int{42, 44}},
			&slice[int]{value: []int{42, 45}},
			false,
		},
		{
			"diff lengths",
			&slice[int]{value: []int{42, 44}},
			&slice[int]{value: []int{42}},
			false,
		},
		{
			"diff name",
			&slice[int]{name: "name0", value: []int{42, 44}},
			&slice[int]{name: "name1", value: []int{42, 44}},
			true,
		},
		{
			"diff kind",
			&slice[int]{value: []int{42, 44}},
			&slice[string]{value: []string{"abc"}},
			false,
		},
		{
			"other nil",
			&slice[int]{value: []int{42, 44}},
			nil,
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := tc.tag0.TagEqual(tc.tag1)

			// --- Then ---
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_array_TagSame(t *testing.T) {
	tt := []struct {
		testN string

		tag0 TagComparer
		tag1 Tag
		want bool
	}{
		{
			"same",
			&slice[int]{name: "name", value: []int{42, 44}},
			&slice[int]{name: "name", value: []int{42, 44}},
			true,
		},
		{
			"diff values",
			&slice[int]{name: "name", value: []int{42, 44}},
			&slice[int]{name: "name", value: []int{52, 54}},
			false,
		},
		{
			"diff value",
			&slice[int]{name: "name", value: []int{42, 44}},
			&slice[int]{name: "name", value: []int{42, 54}},
			false,
		},
		{
			"diff lengths",
			&slice[int]{value: []int{42, 44}},
			&slice[int]{value: []int{42}},
			false,
		},
		{
			"diff name",
			&slice[int]{name: "name0", value: []int{42, 44}},
			&slice[int]{name: "name1", value: []int{42, 44}},
			false,
		},
		{
			"diff kind",
			&slice[int]{name: "name", value: []int{42, 44}},
			&slice[string]{name: "name", value: []string{"abc"}},
			false,
		},
		{
			"other nil",
			&slice[int]{name: "name", value: []int{42, 44}},
			nil,
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := tc.tag0.TagSame(tc.tag1)

			// --- Then ---
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_array_String(t *testing.T) {
	t.Run("nil value", func(t *testing.T) {
		// --- Given ---
		tag := &slice[int]{
			value:    nil,
			stringer: func(v []int) string { return fmt.Sprint(v) },
		}

		// --- When ---
		have := tag.String()

		// --- Then ---
		assert.Equal(t, "[]", have)
	})

	t.Run("single value", func(t *testing.T) {
		// --- Given ---
		tag := &slice[int]{
			value:    []int{42},
			stringer: func(v []int) string { return fmt.Sprint(v) },
		}

		// --- When ---
		have := tag.String()

		// --- Then ---
		assert.Equal(t, "[42]", have)
	})

	t.Run("two values", func(t *testing.T) {
		// --- Given ---
		tag := &slice[int]{
			value:    []int{42, 44},
			stringer: func(v []int) string { return fmt.Sprint(v) },
		}

		// --- When ---
		have := tag.String()

		// --- Then ---
		assert.Equal(t, "[42 44]", have)
	})
}
