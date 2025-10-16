package nomix

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewCreators(t *testing.T) {
	// --- When ---
	reg := NewCreators()

	// --- Then ---
	assert.NotNil(t, reg.m)
	assert.Len(t, 0, reg.m)
}

func Test_Creators_Register(t *testing.T) {
	t.Run("register new", func(t *testing.T) {
		// --- Given ---
		reg := NewCreators()

		var called bool
		tcr := func(name string, value any, _ ...Option) (Tag, error) {
			called = true
			return nil, nil
		}

		// --- When ---
		have := reg.Register(42i+44, CreateFunc(tcr))

		// --- Then ---
		assert.Nil(t, have)

		_, err := reg.Create("name", 42i+44)
		assert.NoError(t, err)
		assert.True(t, called)
	})

	t.Run("register existing", func(t *testing.T) {
		// --- Given ---
		tcrInt := CreateFunc(CreateInt)
		reg := NewCreators()
		reg.Register(42, tcrInt)

		var called bool
		tcr := func(name string, value any, _ ...Option) (Tag, error) {
			called = true
			return nil, nil
		}

		// --- When ---
		have := reg.Register(42, CreateFunc(tcr))

		// --- Then ---
		assert.Same(t, tcrInt, have)

		_, err := reg.Create("name", 42)
		assert.NoError(t, err)
		assert.True(t, called)
	})
}

func Test_Creators_Create_tabular(t *testing.T) {
	tt := []struct {
		testN string

		argValue  any
		kind      TagKind
		wantValue any
	}{
		{"byte", byte(42), KindInt64, int64(42)},
		{"int", int(42), KindInt, int(42)},
		{"int8", int8(42), KindInt64, int64(42)},
		{"int16", int16(42), KindInt64, int64(42)},
		{"int32", int32(42), KindInt64, int64(42)},
		{"int64", int64(42), KindInt64, int64(42)},

		{"float32", float32(42), KindFloat64, float64(42)},
		{"float64", float64(42), KindFloat64, float64(42)},

		{"bool", true, KindBool, true},
		{"string", "abc", KindString, "abc"},
		{"json", json.RawMessage(`{"A": 1}`), KindJSON, []byte(`{"A": 1}`)},
		{
			"time",
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
			KindTime,
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
		},

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

func Test_Creators_Create(t *testing.T) {
	t.Run("error - not registered type", func(t *testing.T) {
		// --- Given ---
		reg := NewCreators()

		// --- When ---
		have, err := reg.Create("name", 42i+44)

		// --- Then ---
		assert.ErrorIs(t, ErrNoCreator, err)
		assert.Nil(t, have)
	})
}
