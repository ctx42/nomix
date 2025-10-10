// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewFloat64Slice(t *testing.T) {
	// --- When ---
	tag := NewFloat64Slice("name", 42.1, 44.2)

	// --- Then ---
	assert.SameType(t, &Float64Slice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []float64{42.1, 44.2}, tag.value)
	assert.Equal(t, KindFloat64Slice, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, "[42.1, 44.2]", tag.stringer([]float64{42.1, 44.2}))
}
