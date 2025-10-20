// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/verax/pkg/verax"
)

func Test_NewSingle(t *testing.T) {
	// --- Given ---
	sqlValuer := func(v int) (driver.Value, error) { return v, nil }

	// --- When ---
	have := NewSingle[int]("name", 42, KindInt, strconv.Itoa, sqlValuer)

	// --- Then ---
	assert.Equal(t, "name", have.name)
	assert.Equal(t, 42, have.value)
	assert.Equal(t, KindInt, have.kind)
	assert.Same(t, strconv.Itoa, have.strValuer)
	assert.Same(t, sqlValuer, have.sqlValuer)
}

func Test_Single_TagName(t *testing.T) {
	tag := &Single[int]{name: "name"}

	// --- When ---
	have := tag.TagName()

	// --- Then ---
	assert.Equal(t, "name", have)
}

func Test_Single_TagKind(t *testing.T) {
	tag := &Single[int]{kind: KindInt}

	// --- When ---
	have := tag.TagKind()

	// --- Then ---
	assert.Equal(t, KindInt, have)
}

func Test_Single_TagValue(t *testing.T) {
	tag := &Single[int]{value: 42}

	// --- When ---
	have := tag.TagValue()

	// --- Then ---
	assert.Equal(t, 42, have)
}

func Test_Single_TagSet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tag := &Single[int]{value: 42}

		// --- When ---
		err := tag.TagSet(44)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, 44, tag.value)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- Given ---
		tag := &Single[int]{value: 42}

		// --- When ---
		err := tag.TagSet(true)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Equal(t, 42, tag.value)
	})
}

func Test_Single_Get(t *testing.T) {
	tag := &Single[int]{value: 42}

	// --- When ---
	have := tag.Get()

	// --- Then ---
	assert.Equal(t, 42, have)
}

func Test_Single_Set(t *testing.T) {
	// --- Given ---
	tag := &Single[int]{value: 42}

	// --- When ---
	tag.Set(44)

	// --- Then ---
	assert.Equal(t, 44, tag.value)
}

func Test_Single_Value(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		// --- Given ---
		tag := &Single[int]{value: 42}

		// --- When ---
		have, err := tag.Value()

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, 42, have)
	})

	t.Run("custom", func(t *testing.T) {
		// --- Given ---
		sqlValuer := func(v int) (driver.Value, error) {
			return float64(v), nil
		}
		tag := &Single[int]{value: 42, sqlValuer: sqlValuer}

		// --- When ---
		have, err := tag.Value()

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, float64(42), have)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		sqlValuer := func(v int) (driver.Value, error) {
			return 0, errors.New("test error")
		}
		tag := &Single[int]{value: 42, sqlValuer: sqlValuer}

		// --- When ---
		have, err := tag.Value()

		// --- Then ---
		assert.ErrorEqual(t, "test error", err)
		assert.Equal(t, 0, have)
	})
}

func Test_Single_TagEqual(t *testing.T) {
	tt := []struct {
		testN string

		tag0 TagValueComparer
		tag1 Tag
		want bool
	}{
		{
			"same",
			&Single[int]{value: 42},
			&Single[int]{value: 42},
			true,
		},
		{
			"diff value",
			&Single[int]{value: 42},
			&Single[int]{value: 44},
			false,
		},
		{
			"diff name",
			&Single[int]{name: "name0", value: 42},
			&Single[int]{name: "name1", value: 42},
			true,
		},
		{
			"diff kind",
			&Single[int]{value: 42},
			&Single[string]{value: "abc"},
			false,
		},
		{
			"other nil",
			&Single[int]{value: 42},
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

func Test_Single_TagSame(t *testing.T) {
	tt := []struct {
		testN string

		tag0 TagComparer
		tag1 Tag
		want bool
	}{
		{
			"same",
			&Single[int]{name: "name", value: 42},
			&Single[int]{name: "name", value: 42},
			true,
		},
		{
			"diff value",
			&Single[int]{name: "name", value: 42},
			&Single[int]{name: "name", value: 44},
			false,
		},
		{
			"diff name",
			&Single[int]{name: "name0", value: 42},
			&Single[int]{name: "name1", value: 42},
			false,
		},
		{
			"diff kind",
			&Single[int]{name: "name", value: 42},
			&Single[string]{name: "name", value: "abc"},
			false,
		},
		{
			"other nil",
			&Single[int]{name: "name", value: 42},
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

func Test_Single_String(t *testing.T) {
	// --- Given ---
	tag := &Single[int]{value: 42, strValuer: strconv.Itoa}

	// --- When ---
	have := tag.String()

	// --- Then ---
	assert.Equal(t, "42", have)
}

func Test_Single_ValidateWith(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tag := &Single[int]{name: "name", value: 42}

		// --- When ---
		err := tag.ValidateWith(verax.Max(42))

		// --- Then ---
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		tag := &Single[int]{name: "name", value: 44}

		// --- When ---
		err := tag.ValidateWith(verax.Max(42))

		// --- Then ---
		assert.ErrorEqual(t, "name: must be no greater than 42", err)
	})
}
