// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewIntSlice(t *testing.T) {
	// --- When ---
	tag := NewIntSlice("name", 42, 44)

	// --- Then ---
	assert.SameType(t, &IntSlice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []int{42, 44}, tag.value)
	assert.Equal(t, KindIntSlice, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, "[42, 44]", tag.stringer([]int{42, 44}))
}
