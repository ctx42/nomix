// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewSpec(t *testing.T) {
	// --- When ---
	have := NewSpec(KindInt, TstIntCreate, TstIntParse)

	// --- Then ---
	assert.Equal(t, KindInt, have.knd)
	assert.Same(t, TstIntCreate, have.tcr)
	assert.Same(t, TstIntParse, have.tpr)
}

func Test_Spec_TagKind(t *testing.T) {
	// --- Given ---
	spec := NewSpec(KindInt, TstIntCreate, TstIntParse)

	// --- Given ---
	have := spec.TagKind()

	// --- Then ---
	assert.Equal(t, KindInt, have)
}

func Test_Spec_TagCreate(t *testing.T) {
	// --- Given ---
	spec := NewSpec(KindInt, TstIntCreate, TstIntParse)

	// --- When ---
	have, err := spec.TagCreate("name", 42)

	// --- Then ---
	assert.NoError(t, err)
	assert.Equal(t, "name", have.TagName())
	assert.Equal(t, 42, have.TagValue())
	assert.Equal(t, KindInt, have.TagKind())
}

func Test_Spec_TagParse(t *testing.T) {
	// --- Given ---
	spec := NewSpec(KindInt, TstIntCreate, TstIntParse)

	// --- When ---
	have, err := spec.TagParse("name", "42")

	// --- Then ---
	assert.NoError(t, err)
	assert.Equal(t, "name", have.TagName())
	assert.Equal(t, 42, have.TagValue())
	assert.Equal(t, KindInt, have.TagKind())
}

func Test_Spec_IsZero_tabular(t *testing.T) {
	tt := []struct {
		testN string

		knd  Kind
		ctr  CreateFunc
		tpr  ParseFunc
		want bool
	}{
		{"zero", 0, nil, nil, true},
		{"kind set", 1, nil, nil, false},
		{"crt set", 0, TstIntCreate, nil, false},
		{"prs set", 0, nil, TstIntParse, false},
		{"all set", 1, TstIntCreate, TstIntParse, false},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			spec := NewSpec(tc.knd, tc.ctr, tc.tpr)

			// --- When ---
			have := spec.IsZero()

			// --- Then ---
			assert.Equal(t, tc.want, have)
		})
	}
}
