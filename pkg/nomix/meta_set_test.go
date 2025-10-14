// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/check"
)

func Test_NewMetaSet(t *testing.T) {
	t.Run("no options", func(t *testing.T) {
		// --- When ---
		have := NewMetaSet()

		// --- Then ---
		assert.NotNil(t, have.m)
		assert.Len(t, 0, have.m)
	})

	t.Run("with the initial map", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"A": 1, "B": nil}

		// --- When ---
		have := NewMetaSet(WithMeta(m))

		// --- Then ---
		assert.Equal(t, map[string]any{"A": 1}, have.m)
		assert.Same(t, m, have.m)
	})
}

func Test_MetaSet_MetaGet(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 1}}

		// --- When ---
		have := set.MetaGet("A")

		// --- Then ---
		assert.Equal(t, 1, have)
	})

	t.Run("not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have := set.MetaGet("A")

		// --- Then ---
		assert.Nil(t, have)
	})
}

func Test_MetaSet_MetaSet(t *testing.T) {
	t.Run("set", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		set.MetaSet("A", 1)

		// --- Then ---
		assert.Equal(t, map[string]any{"A": 1}, set.m)
	})

	t.Run("set existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 1}}

		// --- When ---
		set.MetaSet("A", 2)

		// --- Then ---
		assert.Equal(t, map[string]any{"A": 2}, set.m)
	})

	t.Run("set nil is ignored", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 1}}

		// --- When ---
		set.MetaSet("A", nil)

		// --- Then ---
		assert.Equal(t, map[string]any{"A": 1}, set.m)
	})
}

func Test_MetaSet_MetaDelete(t *testing.T) {
	t.Run("delete not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 1, "B": 2}}

		// --- When ---
		set.MetaDelete("D")

		// --- Then ---
		assert.Equal(t, map[string]any{"A": 1, "B": 2}, set.m)
	})

	t.Run("delete existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 1, "B": 2}}

		// --- When ---
		set.MetaDelete("A")

		// --- Then ---
		assert.Equal(t, map[string]any{"B": 2}, set.m)
	})
}

func Test_MetaSet_MetaCount(t *testing.T) {
	t.Run("empty set", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have := set.MetaCount()

		// --- Then ---
		assert.Equal(t, 0, have)
	})

	t.Run("set with keys", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 1, "B": 2}}

		// --- When ---
		have := set.MetaCount()

		// --- Then ---
		assert.Equal(t, 2, have)
	})
}

func Test_MetaSet_MetaGetAll(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have := set.MetaGetAll()

		// --- Then ---
		assert.NotNil(t, have)
		assert.Len(t, 0, have)
	})

	t.Run("not empty", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 1, "B": 2}}

		// --- When ---
		have := set.MetaGetAll()

		// --- Then ---
		assert.Equal(t, map[string]any{"A": 1, "B": 2}, have)
	})
}

func Test_MetaSet_MetaDeleteAll(t *testing.T) {
	// --- Given ---
	set := MetaSet{m: map[string]any{"A": 1, "B": 2}}

	// --- When ---
	set.MetaDeleteAll()

	// --- Then ---
	assert.Len(t, 0, set.m)
}

func Test_MetaSet_MetaGetString(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": "abc"}}

		// --- When ---
		have, err := set.MetaGetString("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "abc", have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have, err := set.MetaGetString("A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Equal(t, "", have)
	})

	t.Run("error - not string type", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 1}}

		// --- When ---
		have, err := set.MetaGetString("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Equal(t, "", have)
	})
}

func Test_MetaSet_MetaGetStringSlice(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": []string{"abc", "xyz"}}}

		// --- When ---
		have, err := set.MetaGetStringSlice("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []string{"abc", "xyz"}, have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have, err := set.MetaGetStringSlice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - not a slice of strings", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 1}}

		// --- When ---
		have, err := set.MetaGetStringSlice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_MetaSet_MetaGetInt64_existing_tabular(t *testing.T) {
	tt := []struct {
		testN string

		set  map[string]any
		want int64
	}{
		{"int", map[string]any{"A": 1}, 1},
		{"int8", map[string]any{"A": int8(1)}, 1},
		{"int16", map[string]any{"A": int16(1)}, 1},
		{"int32", map[string]any{"A": int32(1)}, 1},
		{"int64", map[string]any{"A": int64(1)}, 1},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			set := MetaSet{m: tc.set}

			// --- When ---
			have, err := set.MetaGetInt64("A")

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_MetaSet_MetaGetInt64(t *testing.T) {
	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have, err := set.MetaGetInt64("A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Equal(t, int64(0), have)
	})

	t.Run("error - not supported type", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": "abc"}}

		// --- When ---
		have, err := set.MetaGetInt64("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Equal(t, int64(0), have)
	})
}

func Test_MetaSet_MetaGetInt64Slice_existing_tabular(t *testing.T) {
	tt := []struct {
		testN string

		set  map[string]any
		want []int64
	}{
		{"int", map[string]any{"A": []int{42, 44}}, []int64{42, 44}},
		{"int8", map[string]any{"A": []int8{42, 44}}, []int64{42, 44}},
		{"int16", map[string]any{"A": []int16{42, 44}}, []int64{42, 44}},
		{"int32", map[string]any{"A": []int32{42, 44}}, []int64{42, 44}},
		{"int64", map[string]any{"A": []int64{42, 44}}, []int64{42, 44}},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			set := MetaSet{m: tc.set}

			// --- When ---
			have, err := set.MetaGetInt64Slice("A")

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_MetaSet_MetaGetInt64Slice(t *testing.T) {
	t.Run("not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have, err := set.MetaGetInt64Slice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - not supported type", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": "abc"}}

		// --- When ---
		have, err := set.MetaGetInt64Slice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_MetaSet_MetaGetFloat64_existing_tabular(t *testing.T) {
	tt := []struct {
		testN string

		set  map[string]any
		want float64
	}{
		{"int", map[string]any{"A": 1}, 1},
		{"int8", map[string]any{"A": int8(1)}, 1},
		{"int16", map[string]any{"A": int16(1)}, 1},
		{"int32", map[string]any{"A": int32(1)}, 1},
		{"int64", map[string]any{"A": int64(1)}, 1},
		{"float32", map[string]any{"A": float32(1)}, 1},
		{"float64", map[string]any{"A": float64(1)}, 1},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			set := MetaSet{m: tc.set}

			// --- When ---
			have, err := set.MetaGetFloat64("A")

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_MetaSet_MetaGetFloat64(t *testing.T) {
	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have, err := set.MetaGetFloat64("A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Equal(t, 0.0, have)
	})

	t.Run("error - not supported type", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": "abc"}}

		// --- When ---
		have, err := set.MetaGetFloat64("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Equal(t, 0.0, have)
	})
}

func Test_MetaSet_MetaGetFloat64Slice_existing_tabular(t *testing.T) {
	tt := []struct {
		testN string

		set  map[string]any
		want []float64
	}{
		{"int", map[string]any{"A": []int{42, 44}}, []float64{42, 44}},
		{"int8", map[string]any{"A": []int8{42, 44}}, []float64{42, 44}},
		{"int16", map[string]any{"A": []int16{42, 44}}, []float64{42, 44}},
		{"int32", map[string]any{"A": []int32{42, 44}}, []float64{42, 44}},
		{"int64", map[string]any{"A": []int64{42, 44}}, []float64{42, 44}},
		{"float32", map[string]any{"A": []float32{42, 44}}, []float64{42, 44}},
		{"float64", map[string]any{"A": []float64{42, 44}}, []float64{42, 44}},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			set := MetaSet{m: tc.set}

			// --- When ---
			have, err := set.MetaGetFloat64Slice("A")

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_MetaSet_MetaGetFloat64Slice(t *testing.T) {
	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have, err := set.MetaGetFloat64Slice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - not supported type", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": "abc"}}

		// --- When ---
		have, err := set.MetaGetFloat64Slice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_MetaSet_MetaGetTime(t *testing.T) {
	t.Run("time instance", func(t *testing.T) {
		// --- Given ---
		tim := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		set := MetaSet{m: map[string]any{"A": tim}}

		// --- When ---
		have, err := set.MetaGetTime("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Exact(t, tim, have)
	})

	t.Run("error - string time is not allowed by default", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": "abc"}}

		// --- When ---
		have, err := set.MetaGetTime("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Zero(t, have)
	})

	t.Run("allow string time", func(t *testing.T) {
		// --- Given ---
		opt := WithTimeFormat(time.RFC3339)
		set := MetaSet{m: map[string]any{"A": "2000-01-02T03:04:05Z"}}

		// --- When ---
		have, err := set.MetaGetTime("A", opt)

		// --- Then ---
		assert.NoError(t, err)
		want := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
		exactTime := check.WithTypeChecker(time.Time{}, check.Exact)
		assert.Exact(t, want, have, exactTime)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have, err := set.MetaGetTime("A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Zero(t, have)
	})

	t.Run("error - not time type", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 1}}

		// --- When ---
		have, err := set.MetaGetTime("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Zero(t, have)
	})
}

func Test_MetaSet_MetaGetTimeSlice(t *testing.T) {
	tim0 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
	tim1 := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)

	t.Run("time intance", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": []time.Time{tim0, tim1}}}

		// --- When ---
		have, err := set.MetaGetTimeSlice("A")

		// --- Then ---
		assert.NoError(t, err)
		exactTime := check.WithTypeChecker(time.Time{}, check.Exact)
		assert.Equal(t, []time.Time{tim0, tim1}, have, exactTime)
	})

	t.Run("error - string type is not allowed by default", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": "abc"}}

		// --- When ---
		have, err := set.MetaGetTimeSlice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})

	t.Run("string time", func(t *testing.T) {
		// --- Given ---
		tim0Str := "2000-01-02T03:04:05Z"
		tim1Str := "2001-01-02T03:04:05Z"
		set := MetaSet{m: map[string]any{"A": []string{tim0Str, tim1Str}}}
		optTimFmt := WithTimeFormat(time.RFC3339)

		// --- When ---
		have, err := set.MetaGetTimeSlice("A", optTimFmt)

		// --- Then ---
		assert.NoError(t, err)
		exactTime := check.WithTypeChecker(time.Time{}, check.Exact)
		assert.Equal(t, []time.Time{tim0, tim1}, have, exactTime)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have, err := set.MetaGetTimeSlice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - not supported type", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 1}}

		// --- When ---
		have, err := set.MetaGetTimeSlice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_MetaSet_MetaGetJSON(t *testing.T) {
	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have, err := set.MetaGetJOSN("A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - invalid", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": "!!!"}}

		// --- When ---
		have, err := set.MetaGetJOSN("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 42}}

		// --- When ---
		have, err := set.MetaGetJOSN("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})

	t.Run("[]byte", func(t *testing.T) {
		// --- Given ---
		data := []byte(`{"field": "value"}`)
		set := MetaSet{m: map[string]any{"A": data}}

		// --- When ---
		have, err := set.MetaGetJOSN("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, string(data), string(have))
	})

	t.Run("json.RawMessage", func(t *testing.T) {
		// --- Given ---
		data := json.RawMessage(`{"field": "value"}`)
		set := MetaSet{m: map[string]any{"A": data}}

		// --- When ---
		have, err := set.MetaGetJOSN("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, string(data), string(have))
	})

	t.Run("string", func(t *testing.T) {
		// --- Given ---
		data := `{"field": "value"}`
		set := MetaSet{m: map[string]any{"A": data}}

		// --- When ---
		have, err := set.MetaGetJOSN("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, data, string(have))
	})
}

func Test_MetaSet_MetaGetBool(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": true}}

		// --- When ---
		have, err := set.MetaGetBool("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.True(t, have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have, err := set.MetaGetBool("A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.False(t, have)
	})

	t.Run("error - not bool type", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 42}}

		// --- When ---
		have, err := set.MetaGetBool("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.False(t, have)
	})
}

func Test_MetaSet_MetaGetBoolSlice(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": []bool{true, false}}}

		// --- When ---
		have, err := set.MetaGetBoolSlice("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []bool{true, false}, have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have, err := set.MetaGetBoolSlice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - not a slice of booleans", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 1}}

		// --- When ---
		have, err := set.MetaGetBoolSlice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_MetaSet_MetaGetInt(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 1}}

		// --- When ---
		have, err := set.MetaGetInt("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, 1, have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have, err := set.MetaGetInt("A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Equal(t, 0, have)
	})

	t.Run("error - not int type", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": "abc"}}

		// --- When ---
		have, err := set.MetaGetInt("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Equal(t, 0, have)
	})
}

func Test_MetaSet_MetaGetIntSlice(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{
			m: map[string]any{
				"A": []int{42, 44},
			},
		}

		// --- When ---
		have, err := set.MetaGetIntSlice("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []int{42, 44}, have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{}}

		// --- When ---
		have, err := set.MetaGetIntSlice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - not slice of ints", func(t *testing.T) {
		// --- Given ---
		set := MetaSet{m: map[string]any{"A": 1}}

		// --- When ---
		have, err := set.MetaGetIntSlice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}
