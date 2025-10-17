package nomix

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/verax/pkg/verax"
)

func Test_Define(t *testing.T) {
	t.Run("no rules", func(t *testing.T) {
		// --- When ---
		have := Define("name", stringSpec)

		// --- Then ---
		assert.Equal(t, "name", have.name)
		assert.Equal(t, stringSpec, have.spec)
		assert.Nil(t, have.rule)
	})

	t.Run("with rules", func(t *testing.T) {
		// --- Given ---
		rMin := verax.Min(42)
		rMax := verax.Max(42)

		// --- When ---
		have := Define("name", intSpec, rMin, rMax)

		// --- Then ---
		assert.Equal(t, "name", have.name)
		assert.Equal(t, intSpec, have.spec)
		assert.NotNil(t, have.rule)
	})
}

func Test_Definition_TagName(t *testing.T) {
	// --- Given ---
	def := &Definition{name: "name", spec: stringSpec}

	// --- When ---
	have := def.TagName()

	// --- Then ---
	assert.Equal(t, "name", have)
}

func Test_Definition_TagKind(t *testing.T) {
	// --- Given ---
	def := &Definition{name: "name", spec: stringSpec}

	// --- When ---
	have := def.TagKind()

	// --- Then ---
	assert.Equal(t, KindString, have)
}

func Test_Definition_TagCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		def := Define("name", stringSpec)

		// --- When ---
		have, err := def.TagCreate("value")

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &String{}, have)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, "value", have.TagValue())
		assert.Equal(t, KindString, have.TagKind())
	})

	t.Run("options are passed to the creator", func(t *testing.T) {
		// --- Given ---
		def := Define("name", timeSpec)

		// --- When ---
		have, err := def.TagCreate("2000-01-02", WithTimeFormat("2006-01-02"))

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &Time{}, have)
		assert.Equal(t, "name", have.TagName())
		wTim := time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, wTim, have.TagValue())
		assert.Equal(t, KindTime, have.TagKind())
	})

	t.Run("error - from creator", func(t *testing.T) {
		// --- Given ---
		def := Define("name", stringSpec)

		// --- When ---
		have, err := def.TagCreate(42)

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element type", err)
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})
}

func Test_Definition_TagParse(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		def := Define("name", intSpec)

		// --- When ---
		have, err := def.TagParse("42")

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &Int{}, have)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, 42, have.TagValue())
		assert.Equal(t, KindInt, have.TagKind())
	})

	t.Run("options are passed to the parser", func(t *testing.T) {
		// --- Given ---
		def := Define("name", intSpec)

		// --- When ---
		have, err := def.TagParse("AA", WithBaseHEX)

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &Int{}, have)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, 170, have.TagValue())
		assert.Equal(t, KindInt, have.TagKind())
	})

	t.Run("error - from creator", func(t *testing.T) {
		// --- Given ---
		def := Define("name", intSpec)

		// --- When ---
		have, err := def.TagParse("abc")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, have)
	})
}

func Test_Definition_Validate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		rMin := verax.Min(42)
		rMax := verax.Max(44)
		def := Define("name", int64Spec, rMin, rMax)

		// --- When ---
		err := def.Validate(42)

		// --- Then ---
		assert.NoError(t, err)
	})

	t.Run("success with a non-primary type", func(t *testing.T) {
		// --- Given ---
		rMin := verax.Min(42)
		rMax := verax.Max(44)
		def := Define("name", int64Spec, rMin, rMax)

		// --- When ---
		err := def.Validate(uint8(42))

		// --- Then ---
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		rMin := verax.Min(42)
		rMax := verax.Max(44)
		def := Define("name", int64Spec, rMin, rMax)

		// --- When ---
		errMin := def.Validate(40)
		errMax := def.Validate(88)

		// --- Then ---
		assert.ErrorEqual(t, "name: must be no less than 42", errMin)
		assert.ErrorEqual(t, "name: must be no greater than 44", errMax)
	})
}
