// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
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
