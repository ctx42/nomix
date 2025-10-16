// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewKindSpec(t *testing.T) {
	// --- Given ---
	create := CreateFunc(CreateInt)
	parse := ParseFunc(ParseInt)

	// --- When ---
	have := NewKindSpec(KindInt, create, parse)

	// --- Then ---
	assert.Equal(t, KindInt, have.knd)
	assert.Same(t, create, have.tcr)
	assert.Same(t, parse, have.tpr)
}

func Test_KindSpec_TagCreate(t *testing.T) {
	// --- Given ---
	ks := NewKindSpec(KindInt, CreateFunc(CreateInt), ParseFunc(ParseInt))

	// --- When ---
	have, err := ks.TagCreate("name", 42)

	// --- Then ---
	assert.NoError(t, err)
	assert.SameType(t, &Int{}, have)
	assert.Equal(t, "name", have.TagName())
	assert.Equal(t, 42, have.TagValue())
	assert.Equal(t, KindInt, have.TagKind())
}

func Test_KindSpec_TagParse(t *testing.T) {
	// --- Given ---
	ks := NewKindSpec(KindInt, CreateFunc(CreateInt), ParseFunc(ParseInt))

	// --- When ---
	have, err := ks.TagParse("name", "42")

	// --- Then ---
	assert.NoError(t, err)
	assert.SameType(t, &Int{}, have)
	assert.Equal(t, "name", have.TagName())
	assert.Equal(t, 42, have.TagValue())
	assert.Equal(t, KindInt, have.TagKind())
}

func Test_KindSpec_IsZero_tabular(t *testing.T) {
	tt := []struct {
		testN string

		knd  TagKind
		ctr  TagCreateFunc
		tpr  TagParseFunc
		want bool
	}{
		{"zero", 0, nil, nil, true},
		{"kind set", 1, nil, nil, false},
		{"crt set", 0, CreateFunc(CreateInt), nil, false},
		{"prs set", 0, nil, ParseFunc(ParseInt), false},
		{"all set", 1, CreateFunc(CreateInt), ParseFunc(ParseInt), false},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			ks := NewKindSpec(tc.knd, tc.ctr, tc.tpr)

			// --- When ---
			have := ks.IsZero()

			// --- Then ---
			assert.Equal(t, tc.want, have)
		})
	}
}
