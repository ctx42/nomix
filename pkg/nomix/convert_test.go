package nomix

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_asString(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := asString("abc", nil)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "abc", have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asString(42, nil)

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})
}

func Test_asStringSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := asStringSlice([]string{"abc", "xyz"}, nil)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []string{"abc", "xyz"}, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asStringSlice(42, nil)

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})
}

func Test_asInt64(t *testing.T) {
	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asInt64("abc", nil)

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})
}

func Test_asInt64_success_tabular(t *testing.T) {
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
			have, err := asInt64(tc.have, nil)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_toInt64Slice(t *testing.T) {
	// --- Given ---
	i32s := []int32{42, 44}

	// --- When ---
	have := toInt64Slice(i32s, nil)

	// --- Then ---
	assert.Equal(t, []int64{42, 44}, have)
}

func Test_asInt64Slice(t *testing.T) {
	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asInt64Slice("abc", nil)

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})
}

func Test_asInt64Slice_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want []int64
	}{
		{"[]int", []int{42, 44}, []int64{42, 44}},
		{"[]int8", []int8{42, 44}, []int64{42, 44}},
		{"[]int16", []int16{42, 44}, []int64{42, 44}},
		{"[]int32", []int32{42, 44}, []int64{42, 44}},
		{"[]int64", []int64{42, 44}, []int64{42, 44}},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := asInt64Slice(tc.have, nil)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_asFloat64(t *testing.T) {
	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asFloat64("abc", nil)

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})
}

func Test_asFloat64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want float64
	}{
		{"int", 42, 42},
		{"byte", byte(42), 42},
		{"int8", int8(42), 42},
		{"int16", int16(42), 42},
		{"int32", int32(42), 42},
		{"int64", int64(42), 42},
		{"float32", float32(42), 42},
		{"float64", 42.0, 42.0},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := asFloat64(tc.have, nil)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_asFloat64Slice(t *testing.T) {
	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asFloat64Slice("abc", nil)

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
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
			have, err := asFloat64Slice(tc.have, nil)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_asTime(t *testing.T) {
	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := asTime(42, nil)

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})

	t.Run("error - by default sting is not supported", func(t *testing.T) {
		// --- When ---
		have, err := asTime("2000-01-02T03:04:05Z", nil)

		// --- Then ---
		assert.ErrorIs(t, err, ErrInvType)
		assert.Empty(t, have)
	})
}

func Test_asTime_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want time.Time
	}{
		{
			"time.Time",
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
		},
		{
			"string",
			"2000-01-02T03:04:05Z",
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			opts := &Options{timeFormat: time.RFC3339}

			// --- When ---
			have, err := asTime(tc.have, opts)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_asTimeSlice(t *testing.T) {
	t.Run("error - not supported type", func(t *testing.T) {
		// --- When ---
		have, err := asTimeSlice(42, nil)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})

	t.Run("error - string type is not allowed by default", func(t *testing.T) {
		// --- When ---
		have, err := asTimeSlice("2000-01-02T03:04:05Z", nil)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})

	t.Run("error - parsing string time", func(t *testing.T) {
		// --- Given ---
		opts := &Options{timeFormat: time.RFC3339}

		// --- When ---
		have, err := asTimeSlice([]string{"abc"}, opts)

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, have)
	})
}

func Test_asTimeSlice_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want []time.Time
	}{
		{
			"[]time.Time",
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
			[]string{"2000-01-02T03:04:05Z", "2001-01-02T03:04:05Z"},
			[]time.Time{
				time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
				time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			opt := &Options{timeFormat: time.RFC3339}

			// --- When ---
			have, err := asTimeSlice(tc.have, opt)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}
