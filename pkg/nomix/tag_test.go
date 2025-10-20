package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_TagKind_String_tabular(t *testing.T) {
	tt := []struct {
		testN string

		knd TagKind
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
