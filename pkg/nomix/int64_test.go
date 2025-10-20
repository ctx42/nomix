// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
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
	assert.Equal(t, KindInt64, tag.TagKind())

	tag, err = have.TagParse("name", "42")
	assert.NoError(t, err)
	assert.SameType(t, &Int64{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, int64(42), tag.TagValue())
	assert.Equal(t, KindInt64, tag.TagKind())
}

func Test_NewInt64(t *testing.T) {
	// --- When ---
	tag := NewInt64("name", 42)

	// --- Then ---
	assert.SameType(t, &Int64{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, int64(42), tag.value)
	assert.Equal(t, KindInt64, tag.kind)
	assert.Equal(t, "42", tag.String())

	val, err := tag.Value()
	assert.NoError(t, err)
	assert.Equal(t, int64(42), val)
}

func Test_CreateInt64(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		tag, err := CreateInt64("name", 42)

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &Int64{}, tag)
		assert.Equal(t, "name", tag.name)
		assert.Equal(t, int64(42), tag.value)
		assert.Equal(t, KindInt64, tag.kind)
		assert.Equal(t, "42", tag.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		tag, err := CreateInt64("name", "abc")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, tag)
	})
}

func Test_createInt64(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := createInt64(42, Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, int64(42), have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := createInt64("abc", Options{})

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := createInt64(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Equal(t, int64(0), have)
	})
}

func Test_createInt64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want int64
	}{
		{"int", 42, 42},
		{"byte", byte(42), 42},
		{"int8", int8(42), 42},
		{"int16", int16(42), 42},
		{"int32", int32(42), 42},
		{"int64", int64(42), 42},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := createInt64(tc.have, Options{})

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_ParseInt64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		str  string
		opts []Option
		exp  int64
	}{
		{"negative", "-1", nil, -1},
		{"zero", "0", nil, 0},
		{"positive", "1", nil, 1},
		{"hex", "AA", []Option{WithBaseHEX}, 170},
		{"negative hex", "-AA", []Option{WithBaseHEX}, -170},
	}

	for _, tc := range tt {
		t.Run(tc.str, func(t *testing.T) {
			t.Parallel()

			// --- When ---
			tag, err := ParseInt64("name", tc.str, tc.opts...)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.exp, tag.TagValue())
		})
	}
}

func Test_ParseInt64_error(t *testing.T) {
	t.Run("error - not supported string value", func(t *testing.T) {
		// --- When ---
		tag, err := ParseInt64("name", "bad")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, tag)
	})

	t.Run("error - hex without option", func(t *testing.T) {
		// --- When ---
		tag, err := ParseInt64("name", "AA")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, tag)
	})
}

func Test_strValueInt64(t *testing.T) {
	// --- When ---
	have := strValueInt64(42)

	// --- Then ---
	assert.Equal(t, "42", have)
}
