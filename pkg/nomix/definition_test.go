// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/verax/pkg/verax"
)

func Test_Define(t *testing.T) {
	t.Run("no rules", func(t *testing.T) {
		// --- Given ---
		spec := TstIntSpec()

		// --- When ---
		have := Define("name", spec)

		// --- Then ---
		assert.Equal(t, "name", have.name)
		assert.Equal(t, spec, have.spec)
		assert.Nil(t, have.rule)
	})

	t.Run("one rule", func(t *testing.T) {
		// --- Given ---
		spec := TstIntSpec()
		r0 := &TstRule{}

		// --- When ---
		have := Define("name", spec, r0)

		// --- Then ---
		assert.Equal(t, "name", have.name)
		assert.Equal(t, spec, have.spec)
		assert.Same(t, r0, have.rule)
	})

	t.Run("with rules", func(t *testing.T) {
		// --- Given ---
		spec := TstIntSpec()
		r0 := &TstRule{}
		r1 := &TstRule{}

		// --- When ---
		have := Define("name", spec, r0, r1)

		// --- Then ---
		assert.Equal(t, "name", have.name)
		assert.Equal(t, spec, have.spec)
		assert.Equal(t, verax.Set{r0, r1}, have.rule)
	})
}

func Test_Definition_TagName(t *testing.T) {
	// --- Given ---
	def := &Definition{name: "name", spec: TstIntSpec()}

	// --- When ---
	have := def.TagName()

	// --- Then ---
	assert.Equal(t, "name", have)
}

func Test_Definition_TagKind(t *testing.T) {
	// --- Given ---
	def := &Definition{name: "name", spec: TstIntSpec()}

	// --- When ---
	have := def.TagKind()

	// --- Then ---
	assert.Equal(t, KindInt, have)
}

func Test_Definition_TagRule(t *testing.T) {
	t.Run("nil when no rule", func(t *testing.T) {
		// --- Given ---
		def := &Definition{name: "name", spec: TstIntSpec()}

		// --- When ---
		have := def.TagRule()

		// --- Then ---
		assert.Nil(t, have)
	})

	t.Run("returns rule", func(t *testing.T) {
		// --- Given ---
		rule := &TstRule{}
		def := &Definition{name: "name", spec: TstIntSpec(), rule: rule}

		// --- When ---
		have := def.TagRule()

		// --- Then ---
		assert.Same(t, rule, have)
	})
}

func Test_Definition_TagCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		def := Define("name", TstIntSpec())

		// --- When ---
		have, err := def.TagCreate(42)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, 42, have.TagValue())
		assert.Equal(t, KindInt, have.TagKind())
	})

	t.Run("options are passed to the creator", func(t *testing.T) {
		// --- Given ---
		def := Define("name", TstIntSpec())

		// --- When ---
		have, err := def.TagCreate("AA", WithRadixHEX)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, 170, have.TagValue())
		assert.Equal(t, KindInt, have.TagKind())
	})

	t.Run("error - from creator", func(t *testing.T) {
		// --- Given ---
		def := Define("name", TstIntSpec())

		// --- When ---
		have, err := def.TagCreate(4.2)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})
}

func Test_Definition_TagParse(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		def := Define("name", TstIntSpec())

		// --- When ---
		have, err := def.TagParse("42")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, 42, have.TagValue())
		assert.Equal(t, KindInt, have.TagKind())
	})

	t.Run("options are passed to the parser", func(t *testing.T) {
		// --- Given ---
		def := Define("name", TstIntSpec())

		// --- When ---
		have, err := def.TagParse("AA", WithRadixHEX)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, 170, have.TagValue())
		assert.Equal(t, KindInt, have.TagKind())
	})

	t.Run("error - from creator", func(t *testing.T) {
		// --- Given ---
		def := Define("name", TstIntSpec())

		// --- When ---
		have, err := def.TagParse("abc")

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - from validator", func(t *testing.T) {
		// --- Given ---
		def := Define("name", TstIntSpec(), verax.Max(42))

		// --- When ---
		have, err := def.TagParse("44")

		// --- Then ---
		assert.SameType(t, &verax.FieldErrors{}, err)
		assert.ErrorEqual(t, "name: must be less or equal to 42", err)
		assert.Nil(t, have)
	})
}

func Test_Definition_Validate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		rMin := verax.Min(42)
		rMax := verax.Max(44)
		def := Define("name", TstIntSpec(), rMin, rMax)

		// --- When ---
		err := def.Validate(42)

		// --- Then ---
		assert.NoError(t, err)
	})

	t.Run("no validation rules", func(t *testing.T) {
		// --- Given ---
		def := Define("name", TstIntSpec())

		// --- When ---
		err := def.Validate(11)

		// --- Then ---
		assert.NoError(t, err)
	})

	t.Run("error - wrong type without validation rules", func(t *testing.T) {
		// --- Given ---
		def := Define("name", TstIntSpec())

		// --- When ---
		err := def.Validate(4.2)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
	})

	t.Run("error - single spec", func(t *testing.T) {
		// --- Given ---
		rMin := verax.Min(42)
		rMax := verax.Max(44)
		def := Define("name", TstIntSpec(), rMin, rMax)

		// --- When ---
		errMin := def.Validate(40)
		errMax := def.Validate(88)

		// --- Then ---
		assert.ErrorEqual(t, "name: must be greater or equal to 42", errMin)
		assert.ErrorEqual(t, "name: must be less or equal to 44", errMax)
	})
}
