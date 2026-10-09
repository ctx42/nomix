// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac.
// SPDX-License-Identifier: MIT

package nomix

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/verax/pkg/verax"
)

func Test_NewSlice(t *testing.T) {
	// --- Given ---
	strValuer := func(v []int) string { return fmt.Sprint(v) }
	sqlValuer := func(v []int) (driver.Value, error) { return v, nil }

	// --- When ---
	have := NewSlice[int]("name", []int{42, 44}, KindInt, strValuer, sqlValuer)

	// --- Then ---
	assert.Equal(t, "name", have.name)
	assert.Equal(t, []int{42, 44}, have.value)
	assert.Equal(t, KindInt, have.kind)
	assert.Same(t, strValuer, have.strValuer)
	assert.Same(t, sqlValuer, have.sqlValuer)
}

func Test_Slice_TagName(t *testing.T) {
	// --- Given ---
	tag := &Slice[int]{name: "name"}

	// --- When ---
	have := tag.TagName()

	// --- Then ---
	assert.Equal(t, "name", have)
}

func Test_Slice_TagKind(t *testing.T) {
	// --- Given ---
	tag := &Slice[int]{kind: KindIntSlice}

	// --- When ---
	have := tag.TagKind()

	// --- Then ---
	assert.Equal(t, KindIntSlice, have)
}

func Test_Slice_TagValue(t *testing.T) {
	// --- Given ---
	tag := &Slice[int]{value: []int{42, 44}}

	// --- When ---
	have := tag.TagValue()

	// --- Then ---
	assert.Equal(t, []int{42, 44}, have)
}

func Test_Slice_TagSet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tag := &Slice[int]{value: []int{42, 44}}

		// --- When ---
		err := tag.TagSet([]int{52, 54})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []int{52, 54}, tag.value)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- Given ---
		tag := &Slice[int]{name: "name", value: []int{42, 44}}

		// --- When ---
		err := tag.TagSet(true)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Equal(t, []int{42, 44}, tag.value)
	})
}

func Test_Slice_Get(t *testing.T) {
	// --- Given ---
	tag := &Slice[int]{value: []int{42, 44}}

	// --- When ---
	have := tag.Get()

	// --- Then ---
	assert.Equal(t, []int{42, 44}, have)
}

func Test_Slice_Set(t *testing.T) {
	// --- Given ---
	tag := &Slice[int]{value: []int{42, 44}}

	// --- When ---
	tag.Set([]int{52, 54})

	// --- Then ---
	assert.Equal(t, []int{52, 54}, tag.value)
}

func Test_Slice_Value(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		// --- Given ---
		tag := &Slice[int]{value: []int{42, 44}}

		// --- When ---
		have, err := tag.Value()

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []int{42, 44}, have)
	})

	t.Run("custom", func(t *testing.T) {
		// --- Given ---
		sqlValuer := func(vs []int) (driver.Value, error) {
			for i := range vs {
				vs[i] += 1
			}
			return vs, nil
		}
		tag := &Slice[int]{value: []int{42, 44}, sqlValuer: sqlValuer}

		// --- When ---
		have, err := tag.Value()

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []int{43, 45}, have)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		sqlValuer := func(v []int) (driver.Value, error) {
			return nil, errors.New("test error")
		}
		tag := &Slice[int]{value: []int{42, 44}, sqlValuer: sqlValuer}

		// --- When ---
		have, err := tag.Value()

		// --- Then ---
		assert.ErrorEqual(t, "test error", err)
		assert.Nil(t, have)
	})
}

func Test_Slice_TagEqual_tabular(t *testing.T) {
	now := time.Now()
	tim := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
	zone := time.FixedZone("X", 3600)

	tt := []struct {
		testN string

		tag0 ValueComparer
		tag1 Tag
		want bool
	}{
		{
			"same",
			&Slice[int]{value: []int{42, 44}},
			&Slice[int]{value: []int{42, 44}},
			true,
		},
		{
			"diff values",
			&Slice[int]{value: []int{42, 44}},
			&Slice[int]{value: []int{52, 45}},
			false,
		},
		{
			"diff value",
			&Slice[int]{value: []int{42, 44}},
			&Slice[int]{value: []int{42, 45}},
			false,
		},
		{
			"diff lengths",
			&Slice[int]{value: []int{42, 44}},
			&Slice[int]{value: []int{42}},
			false,
		},
		{
			"diff name",
			&Slice[int]{name: "name0", value: []int{42, 44}},
			&Slice[int]{name: "name1", value: []int{42, 44}},
			true,
		},
		{
			"diff type",
			&Slice[int]{value: []int{42, 44}},
			&Slice[string]{value: []string{"abc"}},
			false,
		},
		{
			"diff kind",
			&Slice[byte]{kind: KindJSON, value: []byte("1")},
			&Slice[byte]{kind: KindByteSlice, value: []byte("1")},
			false,
		},
		{
			"time without monotonic clock",
			&Slice[time.Time]{value: []time.Time{now}},
			&Slice[time.Time]{value: []time.Time{now.Round(0)}},
			true,
		},
		{
			"same time in diff zone",
			&Slice[time.Time]{value: []time.Time{tim}},
			&Slice[time.Time]{value: []time.Time{tim.In(zone)}},
			true,
		},
		{
			"diff time",
			&Slice[time.Time]{value: []time.Time{tim}},
			&Slice[time.Time]{value: []time.Time{tim.Add(time.Second)}},
			false,
		},
		{
			"other nil",
			&Slice[int]{value: []int{42, 44}},
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

func Test_Slice_TagSame_tabular(t *testing.T) {
	tt := []struct {
		testN string

		tag0 Comparer
		tag1 Tag
		want bool
	}{
		{
			"same",
			&Slice[int]{name: "name", value: []int{42, 44}},
			&Slice[int]{name: "name", value: []int{42, 44}},
			true,
		},
		{
			"diff values",
			&Slice[int]{name: "name", value: []int{42, 44}},
			&Slice[int]{name: "name", value: []int{52, 54}},
			false,
		},
		{
			"diff value",
			&Slice[int]{name: "name", value: []int{42, 44}},
			&Slice[int]{name: "name", value: []int{42, 54}},
			false,
		},
		{
			"diff lengths",
			&Slice[int]{value: []int{42, 44}},
			&Slice[int]{value: []int{42}},
			false,
		},
		{
			"diff name",
			&Slice[int]{name: "name0", value: []int{42, 44}},
			&Slice[int]{name: "name1", value: []int{42, 44}},
			false,
		},
		{
			"diff type",
			&Slice[int]{name: "name", value: []int{42, 44}},
			&Slice[string]{name: "name", value: []string{"abc"}},
			false,
		},
		{
			"diff kind",
			&Slice[byte]{name: "name", kind: KindJSON, value: []byte("1")},
			&Slice[byte]{name: "name", kind: KindByteSlice, value: []byte("1")},
			false,
		},
		{
			"other nil",
			&Slice[int]{name: "name", value: []int{42, 44}},
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

func Test_Slice_String(t *testing.T) {
	t.Run("nil value", func(t *testing.T) {
		// --- Given ---
		tag := &Slice[int]{
			value:     nil,
			strValuer: func(v []int) string { return fmt.Sprint(v) },
		}

		// --- When ---
		have := tag.String()

		// --- Then ---
		assert.Equal(t, "[]", have)
	})

	t.Run("single value", func(t *testing.T) {
		// --- Given ---
		tag := &Slice[int]{
			value:     []int{42},
			strValuer: func(v []int) string { return fmt.Sprint(v) },
		}

		// --- When ---
		have := tag.String()

		// --- Then ---
		assert.Equal(t, "[42]", have)
	})

	t.Run("two values", func(t *testing.T) {
		// --- Given ---
		tag := &Slice[int]{
			value:     []int{42, 44},
			strValuer: func(v []int) string { return fmt.Sprint(v) },
		}

		// --- When ---
		have := tag.String()

		// --- Then ---
		assert.Equal(t, "[42 44]", have)
	})

	t.Run("default", func(t *testing.T) {
		// --- Given ---
		tag := &Slice[int]{value: []int{42, 44}}

		// --- When ---
		have := tag.String()

		// --- Then ---
		assert.Equal(t, "[42 44]", have)
	})
}

func Test_Slice_ValidateWith(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		rule := verax.Each(verax.Max(44))
		tag := &Slice[int]{name: "name", value: []int{42, 44}}

		// --- When ---
		err := tag.ValidateWith(rule)

		// --- Then ---
		assert.NoError(t, err)
	})

	t.Run("error - slice", func(t *testing.T) {
		// --- Given ---
		rule := verax.Length(3, 3)
		tag := &Slice[int]{name: "name", value: []int{42, 44}}

		// --- When ---
		err := tag.ValidateWith(rule)

		// --- Then ---
		assert.ErrorEqual(t, "name: the length must be exactly 3", err)
	})

	t.Run("error - element", func(t *testing.T) {
		// --- Given ---
		rule := verax.Each(verax.Max(42))
		tag := &Slice[int]{name: "name", value: []int{42, 44}}

		// --- When ---
		err := tag.ValidateWith(rule)

		// --- Then ---
		assert.SameType(t, &verax.FieldErrors{}, err)
		assert.ErrorEqual(t, "name.1: must be less or equal to 42", err)
	})
}
