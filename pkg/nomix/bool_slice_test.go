// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewBoolSlice(t *testing.T) {
	// --- When ---
	tag := NewBoolSlice("name", true, false)

	// --- Then ---
	assert.SameType(t, &BoolSlice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []bool{true, false}, tag.value)
	assert.Equal(t, KindBoolSlice, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, "[true, false]", tag.stringer([]bool{true, false}))
}
