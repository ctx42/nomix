// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewTagSet(t *testing.T) {
	t.Run("no options", func(t *testing.T) {
		// --- When ---
		have := NewTagSet()

		// --- Then ---
		assert.Equal(t, 0, have.TagCount())
	})

	t.Run("with initial map", func(t *testing.T) {
		// --- Given ---
		m := map[string]Tag{"A": NewTagMock(t), "B": nil}

		// --- When ---
		have := NewTagSet(WithTags(m))

		// --- Then ---
		assert.Same(t, m["A"], have.TagGet("A"))
		assert.Same(t, m, have.TagGetAll())
	})
}

func Test_TagSet_TagGet(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		m := map[string]Tag{"A": NewTagMock(t)}
		set := NewTagSet(WithTags(m))

		// --- When ---
		have := set.TagGet("A")

		// --- Then ---
		assert.Same(t, m["A"], have)
	})

	t.Run("not existing", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()

		// --- When ---
		have := set.TagGet("B")

		// --- Then ---
		assert.Nil(t, have)
	})
}

func Test_TagSet_TagSet(t *testing.T) {
	t.Run("set", func(t *testing.T) {
		// --- Given ---
		tagA := TstTag(t, "A", KindInt, 1)
		tagB := TstTag(t, "B", KindInt, 2)
		set := NewTagSet()

		// --- When ---
		set.TagSet(tagA, tagB)

		// --- Then ---
		assert.Equal(t, 2, set.TagCount())
		assert.Same(t, tagA, set.TagGet("A"))
		assert.Same(t, tagB, set.TagGet("B"))
	})

	t.Run("set existing", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()
		set.TagSet(TstTag(t, "A", KindInt, 0))
		tagA := TstTag(t, "A", KindInt, 1)
		tagB := TstTag(t, "B", KindInt, 2)

		// --- When ---
		set.TagSet(tagA, tagB)

		// --- Then ---
		assert.Equal(t, 2, set.TagCount())
		assert.Equal(t, 1, set.TagGet("A").TagValue())
	})

	t.Run("nil values are ignored", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()
		set.TagSet(TstTag(t, "A", KindInt, 0))

		tagA := TstTag(t, "A", KindInt, 1)
		tagB := TstTag(t, "B", KindInt, 2)

		// --- When ---
		set.TagSet(tagA, nil, tagB, nil)

		// --- Then ---
		assert.Equal(t, 2, set.TagCount())
		assert.Equal(t, 1, set.TagGet("A").TagValue())
		assert.Equal(t, 2, set.TagGet("B").TagValue())
	})
}

func Test_TagSet_TagDelete(t *testing.T) {
	t.Run("delete not existing", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()
		set.TagSet(TstTag(t, "A", KindInt, 1))

		// --- When ---
		set.TagDelete("B")

		// --- Then ---
		assert.Equal(t, 1, set.TagCount())
		assert.Equal(t, 1, set.TagGet("A").TagValue())
	})

	t.Run("delete existing", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()
		set.TagSet(TstTag(t, "A", KindInt, 1), TstTag(t, "B", KindInt, 2))

		// --- When ---
		set.TagDelete("A")

		// --- Then ---
		assert.Equal(t, 1, set.TagCount())
		assert.Equal(t, 2, set.TagGet("B").TagValue())
	})
}

func Test_TagSet_TagCount(t *testing.T) {
	t.Run("empty set", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()

		// --- When ---
		have := set.TagCount()

		// --- Then ---
		assert.Equal(t, 0, have)
	})

	t.Run("set with keys", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()
		set.TagSet(TstTag(t, "A", KindInt, 1), TstTag(t, "B", KindInt, 2))

		// --- When ---
		have := set.TagCount()

		// --- Then ---
		assert.Equal(t, 2, have)
	})
}

func Test_TagSet_TagGetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		tagA := TstTag(t, "A", KindInt, 1)
		tagB := TstTag(t, "B", KindInt, 2)
		set := NewTagSet()
		set.TagSet(tagA, tagB)

		// --- Then ---
		have := set.TagGetAll()

		// --- Then ---
		want := map[string]Tag{"A": tagA, "B": tagB}
		assert.Equal(t, want, have)
	})

	t.Run("empty set", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()

		// --- Then ---
		assert.Nil(t, set.MetaGetAll())
	})
}

func Test_TagSet_TagDeleteAll(t *testing.T) {
	set := NewTagSet()
	set.TagSet(TstTag(t, "A", KindInt, 1), TstTag(t, "B", KindInt, 2))

	// --- When ---
	set.TagDeleteAll()

	// --- Then ---
	assert.Equal(t, 0, set.TagCount())
}

func Test_TagSet_MetaGetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()
		set.TagSet(TstTag(t, "A", KindInt, 1), TstTag(t, "B", KindInt, 2))

		// --- Then ---
		have := set.MetaGetAll()

		// --- Then ---
		assert.Equal(t, map[string]any{"A": 1, "B": 2}, have)
	})

	t.Run("empty set", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()

		// --- Then ---
		assert.Nil(t, set.MetaGetAll())
	})
}
