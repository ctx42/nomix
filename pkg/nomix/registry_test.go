// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/must"
)

func Test_NewRegistry(t *testing.T) {
	// --- When ---
	reg := NewRegistry()

	// --- Then ---
	assert.NotNil(t, reg.kinds)
	assert.Len(t, 0, reg.kinds)
	assert.NotNil(t, reg.specs)
	assert.Len(t, 0, reg.specs)
}

func Test_Registry_Register(t *testing.T) {
	t.Run("register new", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		// --- When ---
		err := reg.Register(stringSpec)

		// --- Then ---
		assert.NoError(t, err)
		assert.HasKey(t, stringSpec.knd, reg.kinds)
		assert.Len(t, 0, reg.specs)
	})

	t.Run("register existing", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		_ = reg.Register(stringSpec)

		// --- When ---
		err := reg.Register(stringSpec)

		// --- Then ---
		wMsg := "KindSpec for KindString(2) already registered"
		assert.ErrorEqual(t, wMsg, err)
		assert.Len(t, 1, reg.kinds)
		assert.Len(t, 0, reg.specs)
	})
}

func Test_Registry_Associate(t *testing.T) {
	t.Run("associate with existing kind", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		must.Nil(reg.Register(intSpec))

		// --- When ---
		have, err := reg.Associate(42, KindInt)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, TagKind(0), have)
		assert.Len(t, 1, reg.specs)
		assert.Len(t, 1, reg.kinds)
	})

	t.Run("associate with a different kind", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		must.Nil(reg.Register(intSpec))
		must.Nil(reg.Register(int64Spec))
		must.Value(reg.Associate(42, KindInt))

		// --- When ---
		have, err := reg.Associate(42, KindInt64)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, KindInt, have)
		assert.Len(t, 1, reg.specs)
		assert.Len(t, 2, reg.kinds)
	})

	t.Run("error - associate unknown kind", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		// --- When ---
		have, err := reg.Associate(42, KindInt64)

		// --- Then ---
		assert.ErrorEqual(t, "no spec for KindInt64(4)", err)
		assert.Equal(t, TagKind(0), have)
		assert.Len(t, 0, reg.specs)
		assert.Len(t, 0, reg.kinds)
	})
}

func Test_Register_SpecForType(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- When ---
		have := SpecForType(42)

		// --- Then ---
		assert.False(t, have.IsZero())
		assert.Equal(t, intSpec, have)
	})

	t.Run("not existing", func(t *testing.T) {
		// --- When ---
		have := SpecForType(42i + 44)

		// --- Then ---
		assert.True(t, have.IsZero())
	})
}

func Test_Register_SpecForKind(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- When ---
		have := SpecForKind(KindInt)

		// --- Then ---
		assert.False(t, have.IsZero())
		assert.Equal(t, intSpec, have)
	})

	t.Run("not existing", func(t *testing.T) {
		// --- When ---
		have := SpecForKind(0)

		// --- Then ---
		assert.True(t, have.IsZero())
	})
}

func Test_Registry_Create(t *testing.T) {
	t.Run("error - not registered type", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		// --- When ---
		have, err := reg.Create("name", 42i+44)

		// --- Then ---
		assert.ErrorIs(t, ErrNoCreator, err)
		assert.Nil(t, have)
	})
}

func Test_Registry_Create_tabular(t *testing.T) {
	tt := []struct {
		testN string

		argValue  any
		kind      TagKind
		wantValue any
	}{
		{"int", int(42), KindInt, int(42)},
		{"int8", int8(42), KindInt64, int64(42)},
		{"int16", int16(42), KindInt64, int64(42)},
		{"int32", int32(42), KindInt64, int64(42)},
		{"int64", int64(42), KindInt64, int64(42)},
		{"float32", float32(42), KindFloat64, float64(42)},
		{"float64", float64(42), KindFloat64, float64(42)},

		{"bool", true, KindBool, true},
		{"string", "abc", KindString, "abc"},
		{
			"time",
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
			KindTime,
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
		},
		{"json", json.RawMessage(`{"A": 1}`), KindJSON, []byte(`{"A": 1}`)},

		{"byte slice", []byte{42}, KindByteSlice, []byte{42}},
		{"int slice", []int{42}, KindIntSlice, []int{42}},
		{"int8 slice", []int8{42}, KindInt64Slice, []int64{42}},
		{"int16 slice", []int16{42}, KindInt64Slice, []int64{42}},
		{"int32 slice", []int32{42}, KindInt64Slice, []int64{42}},
		{"int64 slice", []int64{42}, KindInt64Slice, []int64{42}},
		{"float64 slice", []float64{42}, KindFloat64Slice, []float64{42}},
		{"float32 slice", []float32{42}, KindFloat64Slice, []float64{42}},

		{"bool slice", []bool{true}, KindBoolSlice, []bool{true}},
		{"string slice", []string{"abc"}, KindStringSlice, []string{"abc"}},
		{
			"time slice",
			[]time.Time{time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)},
			KindTimeSlice,
			[]time.Time{time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)},
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			tag, err := CreateTag(tc.testN, tc.argValue)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.testN, tag.TagName())
			assert.Equal(t, tc.kind, tag.TagKind())
			assert.Equal(t, tc.wantValue, tag.TagValue())
		})
	}
}
