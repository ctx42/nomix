// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewStringSlice(t *testing.T) {
	// --- When ---
	tag := NewStringSlice("name", "abc", "xyz")

	// --- Then ---
	assert.SameType(t, &StringSlice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []string{"abc", "xyz"}, tag.value)
	assert.Equal(t, KindStringSlice, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, `["abc", "xyz"]`, tag.stringer([]string{"abc", "xyz"}))
}
