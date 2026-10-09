package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_Kind_String_tabular(t *testing.T) {
	tt := []struct {
		testN string

		knd Kind
	}{
		{"KindString", KindString},
		{"KindInt64", KindInt64},
		{"KindFloat64", KindFloat64},
		{"KindTime", KindTime},
		{"KindUUID", KindUUID},
		{"KindJSON", KindJSON},
		{"KindBool", KindBool},
		{"KindInt", KindInt},
		{"KindByteSlice", KindByteSlice},
		{"KindStringSlice", KindStringSlice},
		{"KindInt64Slice", KindInt64Slice},
		{"KindFloat64Slice", KindFloat64Slice},
		{"KindTimeSlice", KindTimeSlice},
		{"KindUUIDSlice", KindUUIDSlice},
		{"KindBoolSlice", KindBoolSlice},
		{"KindIntSlice", KindIntSlice},
		{"KindUnknown", 0},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := tc.knd.String()

			// --- Then ---
			assert.Equal(t, tc.testN, have)
		})
	}
}

func Test_TagCreateFunc(t *testing.T) {
	// --- Given ---
	fn := func(name string, val any, _ ...Option) (*Single[int], error) {
		return NewSingle(name, val.(int), KindInt, nil, nil), nil
	}

	// --- When ---
	have := TagCreateFunc(fn)

	// --- Then ---
	tag, err := have("name", 42)
	assert.NoError(t, err)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, 42, tag.TagValue())
	assert.Equal(t, KindInt, tag.TagKind())
}

func Test_TagParseFunc(t *testing.T) {
	// --- Given ---
	fn := func(name, val string, _ ...Option) (*Single[string], error) {
		return NewSingle(name, val, KindString, nil, nil), nil
	}

	// --- When ---
	have := TagParseFunc(fn)

	// --- Then ---
	tag, err := have("name", "abc")
	assert.NoError(t, err)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, "abc", tag.TagValue())
	assert.Equal(t, KindString, tag.TagKind())
}
