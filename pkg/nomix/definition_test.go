package nomix

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Define(t *testing.T) {
	// --- When ---
	have := Define("name", stringSpec)

	// --- Then ---
	assert.Equal(t, "name", have.name)
	assert.Equal(t, stringSpec, have.spec)
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
