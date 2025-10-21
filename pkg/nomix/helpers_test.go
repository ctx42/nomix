// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"
	"time"
	"unsafe"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/check"
	"github.com/ctx42/testing/pkg/must"
)

func Test_GetTag(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tagA := TstTag(t, "A", KindInt, 1)
		set := NewTagSet()
		set.TagSet(tagA)

		// --- Then ---
		have, err := GetTag[*TagMock](set, "A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Same(t, tagA, have)
	})

	t.Run("error - wrong tag type", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()
		set.TagSet(TstTag(t, "A", KindInt, 1))

		// --- Then ---
		have, err := GetTag[int](set, "A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Equal(t, 0, have)
	})

	t.Run("error - missing tag", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()
		set.TagSet()

		// --- Then ---
		have, err := GetTag[*TagMock](set, "A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_GetTagValue(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tagA := TstTag(t, "A", KindInt, 1)
		set := NewTagSet()
		set.TagSet(tagA)

		// --- Then ---
		have, err := GetTagValue[int](set, "A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, 1, have)
	})

	t.Run("error - wrong tag type", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()
		set.TagSet(TstTag(t, "A", KindInt, 1))

		// --- Then ---
		have, err := GetTagValue[float64](set, "A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Equal(t, 0.0, have)
	})

	t.Run("error - missing tag", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()
		set.TagSet()

		// --- Then ---
		have, err := GetTagValue[int](set, "A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Equal(t, 0, have)
	})
}

func Test_GetMetaValue(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		set := NewMetaSet()
		set.MetaSet("A", 1)

		// --- Then ---
		have, err := GetMetaValue[int](set, "A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, 1, have)
	})

	t.Run("error - wrong tag type", func(t *testing.T) {
		// --- Given ---
		set := NewMetaSet()
		set.MetaSet("A", 1)

		// --- Then ---
		have, err := GetMetaValue[float64](set, "A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Equal(t, 0.0, have)
	})

	t.Run("error - missing tag", func(t *testing.T) {
		// --- Given ---
		set := NewMetaSet()

		// --- Then ---
		have, err := GetMetaValue[int](set, "A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Equal(t, 0, have)
	})
}

func Test_CreateInt64_existing_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want int64
	}{
		{"int", 1, 1},
		{"byte", byte(1), 1},
		{"int8", int8(1), 1},
		{"int16", int16(1), 1},
		{"int32", int32(1), 1},
		{"int64", int64(1), 1},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := CreateInt64(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_CreateInt64Slice_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want []int64
	}{
		{"int", []int{42, 44}, []int64{42, 44}},
		{"byte", []byte{42, 44}, []int64{42, 44}},
		{"int8", []int8{42, 44}, []int64{42, 44}},
		{"int16", []int16{42, 44}, []int64{42, 44}},
		{"int32", []int32{42, 44}, []int64{42, 44}},
		{"int64", []int64{42, 44}, []int64{42, 44}},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := CreateInt64Slice(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_CreateFloat64_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want float64
	}{
		{"int", 1, 1},
		{"int8", int8(1), 1},
		{"int16", int16(1), 1},
		{"int32", int32(1), 1},
		{"int64", int64(1), 1},
		{"float32", float32(1), 1},
		{"float64", float64(1), 1},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := CreateFloat64(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}
func Test_CreateFloat64(t *testing.T) {
	t.Run("int64 out of range", func(t *testing.T) {
		// --- Given ---
		var val int64 = 1<<53 + 1

		// --- When ---
		have, err := CreateFloat64(val)

		// --- Then ---
		assert.ErrorContain(t, "int64 value out of range", err)
		assert.Equal(t, 0.0, have)
	})

	t.Run("int64 out of range", func(t *testing.T) {
		// --- Given ---
		var val int
		if unsafe.Sizeof(val) < 8 {
			t.Skip("int is not a 64-bit value")
		}
		val = 1<<53 + 1

		// --- When ---
		have, err := CreateFloat64(val)

		// --- Then ---
		assert.ErrorContain(t, "int value out of range", err)
		assert.Equal(t, 0.0, have)
	})
}

func Test_CreateFloat64Slice_success_tabular(t *testing.T) {
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
			have, err := CreateFloat64Slice(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_CreateFloat64Slice(t *testing.T) {
	t.Run("int64 out of range", func(t *testing.T) {
		// --- Given ---
		var val int64 = 1<<53 + 1

		// --- When ---
		have, err := CreateFloat64Slice([]int64{val})

		// --- Then ---
		assert.ErrorContain(t, "int64 value out of range", err)
		assert.Nil(t, have)
	})

	t.Run("int64 out of range", func(t *testing.T) {
		// --- Given ---
		var val int
		if unsafe.Sizeof(val) < 8 {
			t.Skip("int is not a 64-bit value")
		}
		val = 1<<53 + 1

		// --- When ---
		have, err := CreateFloat64Slice([]int{val})

		// --- Then ---
		assert.ErrorContain(t, "int value out of range", err)
		assert.Nil(t, have)
	})
}

func Test_CreateTime(t *testing.T) {
	t.Run("error - invalid type", func(t *testing.T) {
		// --- Given ---
		opts := Options{}

		// --- When ---
		have, err := CreateTime(42, opts)

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("error - parsing error", func(t *testing.T) {
		// --- Given ---
		opts := Options{TimeFormat: time.RFC3339}

		// --- When ---
		have, err := CreateTime("abc", opts)

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvFormat)
		assert.Empty(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- Given ---
		opts := Options{}

		// --- When ---
		have, err := CreateTime(nil, opts)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Zero(t, have)
	})
}

func Test_CreateTime_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		opts []Option
		want time.Time
	}{
		{
			"time.Time",
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
			[]Option{},
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
		},
		{
			"string",
			"2000-01-02T03:04:05Z",
			[]Option{WithTimeFormat(time.RFC3339)},
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
		},
		{
			"zero value time string",
			"0000-00-00T00:00:00",
			[]Option{WithZeroTime("0000-00-00T00:00:00")},
			time.Time{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			opts := Options{}
			for _, opt := range tc.opts {
				opt(&opts)
			}

			// --- When ---
			have, err := CreateTime(tc.have, opts)

			// --- Then ---
			assert.NoError(t, err)
			assert.Exact(t, tc.want, have)
		})
	}
}

func Test_CreateTimeSlice(t *testing.T) {
	t.Run("error - not supported type", func(t *testing.T) {
		// --- When ---
		have, err := CreateTimeSlice(42, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})

	t.Run("error - disallow string type", func(t *testing.T) {
		// --- Given ---
		times := []string{"2000-01-02T03:04:05Z"}

		// --- When ---
		have, err := CreateTimeSlice(times, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})

	t.Run("error - parsing string time", func(t *testing.T) {
		// --- Given ---
		opts := Options{TimeFormat: time.RFC3339}
		times := []string{"abc"}

		// --- When ---
		have, err := CreateTimeSlice(times, opts)

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := CreateTimeSlice(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})
}

func Test_CreateTimeSlice_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		format string
		loc    *time.Location
		zvt    []string
		have   any
		want   []time.Time
	}{
		{
			"[]time.Time",
			"",
			nil,
			nil,
			[]time.Time{
				time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
				time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC),
			},
			[]time.Time{
				time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
				time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC),
			},
		},
		{
			"[]string",
			time.RFC3339,
			nil,
			nil,
			[]string{"2000-01-02T03:04:05Z", "2001-01-02T03:04:05Z"},
			[]time.Time{
				time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
				time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC),
			},
		},
		{
			"[]string with zero value times",
			time.RFC3339,
			nil,
			[]string{"0000-00-00T00:00:00"},
			[]string{"0000-00-00T00:00:00", "2001-01-02T03:04:05Z"},
			[]time.Time{
				{},
				time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			opts := Options{TimeFormat: tc.format, Location: tc.loc}
			WithZeroTime(tc.zvt...)(&opts)

			// --- When ---
			have, err := CreateTimeSlice(tc.have, opts)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_ParseTime(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		opts := Options{TimeFormat: time.RFC3339}

		// --- When ---
		have, err := ParseTime("2000-01-02T03:04:05Z", opts)

		// --- Then ---
		assert.NoError(t, err)
		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		exactTime := check.WithTypeChecker(time.Time{}, check.Exact)
		assert.Equal(t, want, have, exactTime)
	})

	t.Run("parse in location", func(t *testing.T) {
		// --- Given ---
		WAW := must.Value(time.LoadLocation("Europe/Warsaw"))
		opts := Options{TimeFormat: "2006-01-02", Location: WAW}

		// --- When ---
		have, err := ParseTime("2000-01-02", opts)

		// --- Then ---
		assert.NoError(t, err)
		want := time.Date(2000, 1, 2, 0, 0, 0, 0, WAW)
		exactTime := check.WithTypeChecker(time.Time{}, check.Exact)
		assert.Equal(t, want, have, exactTime)
	})

	t.Run("use UTC if the location is an empty string", func(t *testing.T) {
		// --- Given ---
		zone := time.FixedZone("", 120)
		opts := Options{TimeFormat: "2006-01-02", Location: zone}

		// --- When ---
		have, err := ParseTime("2000-01-02", opts)

		// --- Then ---
		assert.NoError(t, err)
		want := time.Date(2000, 1, 1, 23, 58, 0, 0, time.UTC)
		exactTime := check.WithTypeChecker(time.Time{}, check.Exact)
		assert.Equal(t, want, have, exactTime)
	})

	t.Run("zero time", func(t *testing.T) {
		// --- Given ---
		opts := Options{
			TimeFormat: time.RFC3339,
			zeroTime:   []string{"0000-00-00T00:00:00Z"},
		}

		// --- When ---
		have, err := ParseTime("0000-00-00T00:00:00Z", opts)

		// --- Then ---
		assert.NoError(t, err)
		assert.Zero(t, have)
	})

	t.Run("error - time format isn't set", func(t *testing.T) {
		// --- Given ---
		opts := Options{}

		// --- When ---
		have, err := ParseTime("2000-01-02T03:04:05Z", opts)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Zero(t, have)
	})

	t.Run("error - invalid time format", func(t *testing.T) {
		// --- Given ---
		opts := Options{TimeFormat: time.RFC3339}

		// --- When ---
		have, err := ParseTime("2000-01-02", opts)

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Zero(t, have)
	})

	t.Run("error - invalid time format with a loc option", func(t *testing.T) {
		// --- Given ---
		opts := Options{TimeFormat: time.RFC3339}

		// --- When ---
		have, err := ParseTime("2000-01-02", opts)

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Zero(t, have)
	})
}

func Test_TagParserNotImpl(t *testing.T) {
	// --- When ---
	have, err := TagParserNotImpl("name", "value")

	// --- Then ---
	assert.ErrorIs(t, ErrNotImpl, err)
	assert.ErrorContain(t, "name: tag parser ", err)
	assert.Nil(t, have)
}
