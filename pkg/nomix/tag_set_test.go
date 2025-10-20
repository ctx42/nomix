// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/check"
)

func Test_NewTagSet(t *testing.T) {
	t.Run("no options", func(t *testing.T) {
		// --- When ---
		have := NewTagSet()

		// --- Then ---
		assert.NotNil(t, 0, have.m)
		assert.Len(t, 0, have.m)
	})

	t.Run("with initial map", func(t *testing.T) {
		// --- Given ---
		m := map[string]Tag{
			"A": NewInt("A", 1),
			"B": nil,
		}

		// --- When ---
		have := NewTagSet(WithTags(m))

		// --- Then ---
		assert.Equal(t, map[string]Tag{"A": NewInt("A", 1)}, have.m)
		assert.Same(t, m, have.m)
	})
}

func Test_TagSet_TagGet(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		have := set.TagGet("A")

		// --- Then ---
		assert.Equal(t, NewInt("A", 1), have)
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
		set := TagSet{m: map[string]Tag{}}

		// --- When ---
		set.TagSet(NewInt("A", 1), NewInt("B", 2))

		// --- Then ---
		assert.Len(t, 2, set.m)
		assert.HasKeyValue(t, "A", Tag(NewInt("A", 1)), set.m)
		assert.HasKeyValue(t, "B", Tag(NewInt("B", 2)), set.m)
	})

	t.Run("set existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}
		tagA := NewInt("A", 2)
		tagB := NewInt("B", 3)

		// --- When ---
		set.TagSet(tagA, tagB)

		// --- Then ---
		assert.Len(t, 2, set.m)
		assert.Same(t, tagA, set.m["A"])
		assert.Same(t, tagB, set.m["B"])
	})

	t.Run("nil values are ignored", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		set.TagSet(NewInt("A", 2), nil, NewInt("B", 3), nil)

		// --- Then ---
		assert.Len(t, 2, set.m)
		want := map[string]Tag{"A": NewInt("A", 2), "B": NewInt("B", 3)}
		assert.Equal(t, want, set.m)
	})
}

func Test_TagSet_TagDelete(t *testing.T) {
	t.Run("delete not existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		set.TagDelete("B")

		// --- Then ---
		assert.Len(t, 1, set.m)
		assert.HasKeyValue(t, "A", Tag(NewInt("A", 1)), set.m)
	})

	t.Run("delete existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
				"B": NewInt("B", 2),
			},
		}

		// --- When ---
		set.TagDelete("A")

		// --- Then ---
		assert.Len(t, 1, set.m)
		assert.HasKeyValue(t, "B", Tag(NewInt("B", 2)), set.m)
	})
}

func Test_TagSet_TagCount(t *testing.T) {
	t.Run("empty set", func(t *testing.T) {
		// --- Given ---
		set := TagSet{}

		// --- When ---
		have := set.TagCount()

		// --- Then ---
		assert.Equal(t, 0, have)
	})

	t.Run("set with keys", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
				"B": NewInt("B", 2),
			},
		}

		// --- When ---
		have := set.TagCount()

		// --- Then ---
		assert.Equal(t, 2, have)
	})
}

func Test_TagSet_TagGetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
				"B": NewInt("B", 2),
			},
		}

		// --- Then ---
		have := set.TagGetAll()

		// --- Then ---
		assert.Same(t, set.m, have)
	})

	t.Run("empty set", func(t *testing.T) {
		// --- Given ---
		set := NewTagSet()

		// --- Then ---
		assert.Nil(t, set.MetaGetAll())
	})
}

func Test_TagSet_TagDeleteAll(t *testing.T) {
	// --- Given ---
	set := TagSet{
		m: map[string]Tag{
			"A": NewInt("A", 1),
			"B": NewInt("B", 2),
		},
	}

	// --- When ---
	set.TagDeleteAll()

	// --- Then ---
	assert.Len(t, 0, set.m)
}

func Test_TagSet_MetaGetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
				"B": NewInt("B", 2),
			},
		}

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

func Test_TagSet_TagGetString(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewString("A", "1"),
			},
		}

		// --- When ---
		have, err := set.TagGetString("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, NewString("A", "1"), have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewString("A", "1"),
			},
		}

		// --- When ---
		have, err := set.TagGetString("B")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "B: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - wrong type", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		have, err := set.TagGetString("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_TagSet_TagGetStringSlice(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewStringSlice("A", "abc", "xyz"),
			},
		}

		// --- When ---
		have, err := set.TagGetStringSlice("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, NewStringSlice("A", "abc", "xyz"), have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewStringSlice("A", "abc", "xyz"),
			},
		}

		// --- When ---
		have, err := set.TagGetStringSlice("B")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "B: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - wrong type", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		have, err := set.TagGetStringSlice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_TagSet_TagGetInt64(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt64("A", 1),
			},
		}

		// --- When ---
		have, err := set.TagGetInt64("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, NewInt64("A", 1), have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt64("A", 1),
			},
		}

		// --- When ---
		have, err := set.TagGetInt64("B")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "B: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - wrong type", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		have, err := set.TagGetInt64("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_TagSet_TagGetInt64Slice(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt64Slice("A", 42, 44),
			},
		}

		// --- When ---
		have, err := set.TagGetInt64Slice("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, NewInt64Slice("A", 42, 44), have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt64Slice("A", 42, 44),
			},
		}

		// --- When ---
		have, err := set.TagGetInt64Slice("B")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "B: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - wrong type", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		have, err := set.TagGetInt64Slice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_TagSet_TagGetFloat64(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewFloat64("A", 1.1),
			},
		}

		// --- When ---
		have, err := set.TagGetFloat64("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, NewFloat64("A", 1.1), have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewFloat64("A", 1.1),
			},
		}

		// --- When ---
		have, err := set.TagGetFloat64("B")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "B: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - wrong type", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		have, err := set.TagGetFloat64("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_TagSet_TagGetFloat64Slice(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewFloat64Slice("A", 4.2, 4.4),
			},
		}

		// --- When ---
		have, err := set.TagGetFloat64Slice("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, NewFloat64Slice("A", 4.2, 4.4), have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewFloat64Slice("A", 4.2, 4.4),
			},
		}

		// --- When ---
		have, err := set.TagGetFloat64Slice("B")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "B: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - wrong type", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		have, err := set.TagGetFloat64Slice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_TagSet_TagGetTime(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		tim := time.Now()
		set := TagSet{
			m: map[string]Tag{
				"A": NewTime("A", tim),
			},
		}

		// --- When ---
		have, err := set.TagGetTime("A")

		// --- Then ---
		assert.NoError(t, err)
		exactTime := check.WithTypeChecker(time.Time{}, check.Exact)
		assert.Equal(t, NewTime("A", tim), have, exactTime)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewTime("A", time.Now()),
			},
		}

		// --- When ---
		have, err := set.TagGetTime("B")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "B: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - wrong type", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		have, err := set.TagGetTime("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_TagSet_TagGetTimeSlice(t *testing.T) {
	tim0 := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)
	tim1 := time.Date(2001, 1, 2, 3, 4, 5, 0, time.UTC)

	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewTimeSlice("A", tim0, tim1),
			},
		}

		// --- When ---
		have, err := set.TagGetTimeSlice("A")

		// --- Then ---
		assert.NoError(t, err)
		want := NewTimeSlice("A", tim0, tim1)
		exactTime := check.WithTypeChecker(time.Time{}, check.Exact)
		assert.Equal(t, want, have, exactTime)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewTimeSlice("A", tim0, tim1),
			},
		}

		// --- When ---
		have, err := set.TagGetTimeSlice("B")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "B: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - wrong type", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		have, err := set.TagGetTimeSlice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_TagSet_TagGetJSON(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		data := json.RawMessage(`{"field": "value"}`)
		set := TagSet{
			m: map[string]Tag{
				"A": NewJSON("A", data),
			},
		}

		// --- When ---
		have, err := set.TagGetJSON("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, NewJSON("A", data), have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		data := json.RawMessage(`{"field": "value"}`)
		set := TagSet{
			m: map[string]Tag{
				"A": NewJSON("A", data),
			},
		}

		// --- When ---
		have, err := set.TagGetJSON("B")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "B: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - wrong type", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		have, err := set.TagGetJSON("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_TagSet_TagGetBool(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewBool("A", true),
			},
		}

		// --- When ---
		have, err := set.TagGetBool("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, NewBool("A", true), have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewBool("A", true),
			},
		}

		// --- When ---
		have, err := set.TagGetBool("B")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "B: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - wrong type", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		have, err := set.TagGetBool("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_TagSet_TagGetBoolSlice(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewBoolSlice("A", true, false),
			},
		}

		// --- When ---
		have, err := set.TagGetBoolSlice("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, NewBoolSlice("A", true, false), have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewBoolSlice("A", true, false),
			},
		}

		// --- When ---
		have, err := set.TagGetBoolSlice("B")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "B: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - wrong type", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		have, err := set.TagGetBoolSlice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_TagSet_TagGetInt(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 42),
			},
		}

		// --- When ---
		have, err := set.TagGetInt("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, NewInt("A", 42), have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 42),
			},
		}

		// --- When ---
		have, err := set.TagGetInt("B")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "B: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - wrong type", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewString("A", "abc"),
			},
		}

		// --- When ---
		have, err := set.TagGetInt("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}

func Test_TagSet_TagGetIntSlice(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewIntSlice("A", 42, 44),
			},
		}

		// --- When ---
		have, err := set.TagGetIntSlice("A")

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, NewIntSlice("A", 42, 44), have)
	})

	t.Run("error - not existing", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewIntSlice("A", 42, 44),
			},
		}

		// --- When ---
		have, err := set.TagGetIntSlice("B")

		// --- Then ---
		assert.ErrorIs(t, ErrMissing, err)
		assert.ErrorContain(t, "B: ", err)
		assert.Nil(t, have)
	})

	t.Run("error - wrong type", func(t *testing.T) {
		// --- Given ---
		set := TagSet{
			m: map[string]Tag{
				"A": NewInt("A", 1),
			},
		}

		// --- When ---
		have, err := set.TagGetIntSlice("A")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "A: ", err)
		assert.Nil(t, have)
	})
}
