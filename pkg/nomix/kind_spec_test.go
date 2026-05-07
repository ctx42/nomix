// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/must"
	"github.com/ctx42/verax/pkg/spec"
)

func Test_NewSpec(t *testing.T) {
	// --- When ---
	have := NewKindSpec(KindInt, TstIntCreate, TstIntParse)

	// --- Then ---
	assert.Equal(t, KindInt, have.knd)
	assert.Same(t, TstIntCreate, have.tcr)
	assert.Same(t, TstIntParse, have.tpr)
}

func Test_KindSpec_TagKind(t *testing.T) {
	// --- Given ---
	spec := NewKindSpec(KindInt, TstIntCreate, TstIntParse)

	// --- Given ---
	have := spec.TagKind()

	// --- Then ---
	assert.Equal(t, KindInt, have)
}

func Test_KindSpec_TagCreate(t *testing.T) {
	// --- Given ---
	spec := NewKindSpec(KindInt, TstIntCreate, TstIntParse)

	// --- When ---
	have, err := spec.TagCreate("name", 42)

	// --- Then ---
	assert.NoError(t, err)
	assert.Equal(t, "name", have.TagName())
	assert.Equal(t, 42, have.TagValue())
	assert.Equal(t, KindInt, have.TagKind())
}

func Test_KindSpec_TagParse(t *testing.T) {
	// --- Given ---
	spec := NewKindSpec(KindInt, TstIntCreate, TstIntParse)

	// --- When ---
	have, err := spec.TagParse("name", "42")

	// --- Then ---
	assert.NoError(t, err)
	assert.Equal(t, "name", have.TagName())
	assert.Equal(t, 42, have.TagValue())
	assert.Equal(t, KindInt, have.TagKind())
}

func Test_KindSpec_IsZero_tabular(t *testing.T) {
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
			spec := NewKindSpec(tc.knd, tc.ctr, tc.tpr)

			// --- When ---
			have := spec.IsZero()

			// --- Then ---
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_KindSpec_Spec(t *testing.T) {
	// --- Given ---
	spc := TstIntSpec()

	// --- When ---
	have, err := spc.Spec()

	// --- Then ---
	assert.NoError(t, err)
	assert.Equal(t, KindSpecName, have.Name)
	assert.Equal(t, int16(KindInt), have.Args[spec.ArgValue])
}

func Test_KindSpecFromSpec(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		must.Nil(reg.Register(TstIntSpec()))

		spc := spec.NewSpec(KindSpecName).SetArg(spec.ArgValue, int16(KindInt))

		// --- When ---
		have, err := KindSpecFromSpec(reg, spc)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, KindInt, have.knd)
		assert.Same(t, TstIntCreate, have.tcr)
		assert.Same(t, TstIntParse, have.tpr)
	})

	t.Run("error - invalid spec name", func(t *testing.T) {
		// --- Given ---
		spc := spec.NewSpec("other").SetArg(spec.ArgValue, int16(KindInt))

		// --- When ---
		have, err := KindSpecFromSpec(NewRegistry(), spc)

		// --- Then ---
		assert.SameType(t, &InternalError{}, err)
		wMsg := fmt.Sprintf(`%s: invalid spec name: "other"`, KindSpecName)
		assert.ErrorEqual(t, wMsg, err)
		assert.Equal(t, KindSpec{}, have)
	})

	t.Run("error - missing value argument", func(t *testing.T) {
		// --- Given ---
		spc := spec.NewSpec(KindSpecName)

		// --- When ---
		have, err := KindSpecFromSpec(NewRegistry(), spc)

		// --- Then ---
		assert.SameType(t, &InternalError{}, err)
		wMsg := "kind-spec: spec missing required argument: value"
		assert.ErrorEqual(t, wMsg, err)
		assert.Equal(t, KindSpec{}, have)
	})

	t.Run("error - value argument is of the wrong type", func(t *testing.T) {
		// --- Given ---
		spc := spec.NewSpec(KindSpecName).SetArg(spec.ArgValue, "not-int16")

		// --- When ---
		have, err := KindSpecFromSpec(NewRegistry(), spc)

		// --- Then ---
		assert.SameType(t, &InternalError{}, err)
		wMsg := `kind-spec: spec argument "value" must be int16, got string`
		assert.ErrorEqual(t, wMsg, err)
		assert.Equal(t, KindSpec{}, have)
	})

	t.Run("error - kind not registered", func(t *testing.T) {
		// --- Given ---
		spc := spec.NewSpec(KindSpecName).SetArg(spec.ArgValue, int16(KindInt))

		// --- When ---
		have, err := KindSpecFromSpec(NewRegistry(), spc)

		// --- Then ---
		assert.SameType(t, &InternalError{}, err)
		assert.ErrorEqual(t, `kind-spec: invalid spec kind: 516`, err)
		assert.Equal(t, KindSpec{}, have)
	})
}

func Test_KindSpec_Spec_KindSpecFromSpec_round_trip(t *testing.T) {
	// --- Given ---
	reg := NewRegistry()
	must.Nil(reg.Register(TstIntSpec()))

	// --- When ---
	s, err := TstIntSpec().Spec()

	// --- Then ---
	assert.NoError(t, err)

	// --- When ---
	have, err := KindSpecFromSpec(reg, s)

	// --- Then ---
	assert.NoError(t, err)
	assert.Equal(t, KindInt, have.knd)
	assert.Same(t, TstIntCreate, have.tcr)
	assert.Same(t, TstIntParse, have.tpr)
}
