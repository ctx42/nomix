// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/nomix/pkg/nomix"
)

func Test_RegisterAll_tabular(t *testing.T) {
	tt := []struct {
		testN string

		argValue  any
		kind      nomix.Kind
		wantValue any
	}{
		{"int", int(42), nomix.KindInt, int(42)},
		{"int8", int8(42), nomix.KindInt64, int64(42)},
		{"int16", int16(42), nomix.KindInt64, int64(42)},
		{"int32", int32(42), nomix.KindInt64, int64(42)},
		{"int64", int64(42), nomix.KindInt64, int64(42)},
		{"float32", float32(42), nomix.KindFloat64, float64(42)},
		{"float64", float64(42), nomix.KindFloat64, float64(42)},
		{"bool", true, nomix.KindBool, true},
		{"string", "abc", nomix.KindString, "abc"},
		{
			"time",
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
			nomix.KindTime,
			time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
		},
		{
			"json",
			json.RawMessage(`{"A": 1}`),
			nomix.KindJSON,
			[]byte(`{"A": 1}`),
		},

		{"byte slice", []byte{42}, nomix.KindByteSlice, []byte{42}},
		{"int slice", []int{42}, nomix.KindIntSlice, []int{42}},
		{"int8 slice", []int8{42}, nomix.KindInt64Slice, []int64{42}},
		{"int16 slice", []int16{42}, nomix.KindInt64Slice, []int64{42}},
		{"int32 slice", []int32{42}, nomix.KindInt64Slice, []int64{42}},
		{"int64 slice", []int64{42}, nomix.KindInt64Slice, []int64{42}},
		{"float64 slice", []float64{42}, nomix.KindFloat64Slice, []float64{42}},
		{"float32 slice", []float32{42}, nomix.KindFloat64Slice, []float64{42}},
		{"bool slice", []bool{true}, nomix.KindBoolSlice, []bool{true}},
		{"string slice", []string{"abc"}, nomix.KindStringSlice, []string{"abc"}},
		{
			"time slice",
			[]time.Time{time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)},
			nomix.KindTimeSlice,
			[]time.Time{time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)},
		},
	}

	reg := nomix.NewRegistry()
	RegisterAll(reg)

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			tag, err := reg.Create(tc.testN, tc.argValue)
			assert.NoError(t, err)
			assert.Equal(t, tc.testN, tag.TagName())
			assert.Equal(t, tc.kind, tag.TagKind())
			assert.Equal(t, tc.wantValue, tag.TagValue())
		})
	}
}
