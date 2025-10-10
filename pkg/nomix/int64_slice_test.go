// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewInt64Slice(t *testing.T) {
	// --- When ---
	tag := NewInt64Slice("name", 42, 44)

	// --- Then ---
	assert.SameType(t, &Int64Slice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []int64{42, 44}, tag.value)
	assert.Equal(t, KindInt64Slice, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, "[42, 44]", tag.stringer([]int64{42, 44}))
}
