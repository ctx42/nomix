// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"strconv"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/verax/pkg/verax"
)

func Test_single_TagName(t *testing.T) {
	tag := &single[int]{name: "name"}

	// --- When ---
	have := tag.TagName()

	// --- Then ---
	assert.Equal(t, "name", have)
}

func Test_single_TagKind(t *testing.T) {
	tag := &single[int]{kind: KindInt}

	// --- When ---
	have := tag.TagKind()

	// --- Then ---
	assert.Equal(t, KindInt, have)
}

func Test_single_TagValue(t *testing.T) {
	tag := &single[int]{value: 42}

	// --- When ---
	have := tag.TagValue()

	// --- Then ---
	assert.Equal(t, 42, have)
}

func Test_single_Value(t *testing.T) {
	tag := &single[int]{value: 42}

	// --- When ---
	have := tag.Value()

	// --- Then ---
	assert.Equal(t, 42, have)
}

func Test_single_Set(t *testing.T) {
	// --- Given ---
	tag := &single[int]{value: 42}

	// --- When ---
	tag.Set(44)

	// --- Then ---
	assert.Equal(t, 44, tag.value)
}

func Test_single_TagSet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tag := &single[int]{value: 42}

		// --- When ---
		err := tag.TagSet(44)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, 44, tag.value)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- Given ---
		tag := &single[int]{value: 42}

		// --- When ---
		err := tag.TagSet(true)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Equal(t, 42, tag.value)
	})
}

func Test_single_TagEqual(t *testing.T) {
	tt := []struct {
		testN string

		tag0 TagValueComparer
		tag1 Tag
		want bool
	}{
		{
			"same",
			&single[int]{value: 42},
			&single[int]{value: 42},
			true,
		},
		{
			"diff value",
			&single[int]{value: 42},
			&single[int]{value: 44},
			false,
		},
		{
			"diff name",
			&single[int]{name: "name0", value: 42},
			&single[int]{name: "name1", value: 42},
			true,
		},
		{
			"diff kind",
			&single[int]{value: 42},
			&single[string]{value: "abc"},
			false,
		},
		{
			"other nil",
			&single[int]{value: 42},
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

func Test_single_TagSame(t *testing.T) {
	tt := []struct {
		testN string

		tag0 TagComparer
		tag1 Tag
		want bool
	}{
		{
			"same",
			&single[int]{name: "name", value: 42},
			&single[int]{name: "name", value: 42},
			true,
		},
		{
			"diff value",
			&single[int]{name: "name", value: 42},
			&single[int]{name: "name", value: 44},
			false,
		},
		{
			"diff name",
			&single[int]{name: "name0", value: 42},
			&single[int]{name: "name1", value: 42},
			false,
		},
		{
			"diff kind",
			&single[int]{name: "name", value: 42},
			&single[string]{name: "name", value: "abc"},
			false,
		},
		{
			"other nil",
			&single[int]{name: "name", value: 42},
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

func Test_single_String(t *testing.T) {
	// --- Given ---
	tag := &single[int]{value: 42, stringer: strconv.Itoa}

	// --- When ---
	have := tag.String()

	// --- Then ---
	assert.Equal(t, "42", have)
}

func Test_single_ValidateWith(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tag := &single[int]{name: "name", value: 42}

		// --- When ---
		err := tag.ValidateWith(verax.Max(42))

		// --- Then ---
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tag := &single[int]{name: "name", value: 44}

		// --- When ---
		err := tag.ValidateWith(verax.Max(42))

		// --- Then ---
		assert.ErrorEqual(t, "name: must be no greater than 42", err)
	})
}
