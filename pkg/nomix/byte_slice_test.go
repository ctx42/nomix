// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewByteSlice(t *testing.T) {
	// --- When ---
	tag := NewByteSlice("name", 42, 44)

	// --- Then ---
	assert.SameType(t, &ByteSlice{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []byte{42, 44}, tag.value)
	assert.Equal(t, KindByteSlice, tag.kind)
	assert.NotNil(t, tag.stringer)
	assert.Equal(t, "[42, 44]", tag.stringer([]byte{42, 44}))
}
