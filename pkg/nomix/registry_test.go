// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/must"
)

func Test_GlobalRegistry(t *testing.T) {
	// --- When ---
	have := GlobalRegistry()

	// --- Then ---
	assert.Same(t, specs, have)
}

func Test_NewRegistry(t *testing.T) {
	// --- When ---
	reg := NewRegistry()

	// --- Then ---
	assert.NotNil(t, reg.kinds)
	assert.Len(t, 0, reg.kinds)
	assert.NotNil(t, reg.specs)
	assert.Len(t, 0, reg.specs)
}

func Test_Registry_Register(t *testing.T) {
	t.Run("register new", func(t *testing.T) {
		// --- Given ---
		spec := TstIntSpec()
		reg := NewRegistry()

		// --- When ---
		err := reg.Register(spec)

		// --- Then ---
		assert.NoError(t, err)
		assert.HasKey(t, KindInt, reg.kinds)
		assert.Len(t, 0, reg.specs)
	})

	t.Run("register existing", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		_ = reg.Register(TstIntSpec())

		// --- When ---
		err := reg.Register(TstIntSpec())

		// --- Then ---
		wMsg := "spec for KindInt(516) already registered"
		assert.ErrorEqual(t, wMsg, err)
		assert.Len(t, 1, reg.kinds)
		assert.Len(t, 0, reg.specs)
	})
}

func Test_Registry_Associate(t *testing.T) {
	t.Run("associate with existing kind", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		must.Nil(reg.Register(TstIntSpec()))

		// --- When ---
		have, err := reg.Associate(42, KindInt)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Kind(0), have)
		assert.Len(t, 1, reg.specs)
		assert.Len(t, 1, reg.kinds)
	})

	t.Run("associate with a different kind", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		must.Nil(reg.Register(Spec{knd: KindInt}))
		must.Nil(reg.Register(Spec{knd: KindInt64}))
		must.Value(reg.Associate(42, KindInt))

		// --- When ---
		have, err := reg.Associate(42, KindInt64)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, KindInt, have)
		assert.Len(t, 1, reg.specs)
		assert.Len(t, 2, reg.kinds)
	})

	t.Run("error - associate unknown kind", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		// --- When ---
		have, err := reg.Associate(42, KindInt64)

		// --- Then ---
		assert.ErrorEqual(t, "no spec for KindInt64(4)", err)
		assert.Equal(t, Kind(0), have)
		assert.Len(t, 0, reg.specs)
		assert.Len(t, 0, reg.kinds)
	})
}

func Test_Register_SpecForType(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		spec := TstIntSpec()
		reg := NewRegistry()
		must.Nil(reg.Register(spec))
		must.Value(reg.Associate(42, KindInt))

		// --- When ---
		have := reg.SpecForType(42)

		// --- Then ---
		assert.False(t, have.IsZero())
		assert.Equal(t, spec, have)
	})

	t.Run("not existing", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		// --- When ---
		have := reg.SpecForType(42i + 44)

		// --- Then ---
		assert.True(t, have.IsZero())
	})
}

func Test_Register_SpecForKind(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		spec := TstIntSpec()
		reg := NewRegistry()
		must.Nil(reg.Register(spec))

		// --- When ---
		have := reg.SpecForKind(KindInt)

		// --- Then ---
		assert.False(t, have.IsZero())
		assert.Equal(t, spec, have)
	})

	t.Run("not existing", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		// --- When ---
		have := reg.SpecForKind(0)

		// --- Then ---
		assert.True(t, have.IsZero())
	})
}

func Test_Registry_Create(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		must.Nil(reg.Register(TstIntSpec()))
		must.Value(reg.Associate(42, KindInt))

		// --- When ---
		have, err := reg.Create("name", 42)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, 42, have.TagValue())
		assert.Equal(t, KindInt, have.TagKind())
	})

	t.Run("error - not registered type", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		// --- When ---
		have, err := reg.Create("name", 42i+44)

		// --- Then ---
		assert.ErrorIs(t, ErrNoCreator, err)
		wMsg := "creator not found for name of type complex128"
		assert.ErrorEqual(t, wMsg, err)
		assert.Nil(t, have)
	})
}
