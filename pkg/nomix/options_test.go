// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/must"
)

func Test_NewOptions(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		// --- When ---
		have := NewOptions()

		// --- Then ---
		assert.Equal(t, 0, have.Length)
		assert.Nil(t, have.init)
		assert.Equal(t, time.RFC3339Nano, have.TimeFormat)
		assert.Nil(t, have.Location)
		assert.False(t, have.LocationAsString)
		assert.Empty(t, have.zeroTime)
		assert.Equal(t, 10, have.Radix)
		assert.Fields(t, 7, have)
	})

	t.Run("with changes", func(t *testing.T) {
		// --- When ---
		have := NewOptions(WithTimeFormat("2006"))

		// --- Then ---
		assert.Equal(t, 0, have.Length)
		assert.Nil(t, have.init)
		assert.Equal(t, "2006", have.TimeFormat)
		assert.Nil(t, have.Location)
		assert.False(t, have.LocationAsString)
		assert.Empty(t, have.zeroTime)
		assert.Equal(t, 10, have.Radix)
		assert.Fields(t, 7, have)
	})
}

func Test_WithLen(t *testing.T) {
	// --- Given ---
	opts := &Options{}

	// --- When ---
	WithLen(10)(opts)

	// --- Then ---
	assert.Equal(t, 10, opts.Length)
}

func Test_WithMeta(t *testing.T) {
	// --- Given ---
	m := map[string]any{"A": 1}
	opts := &Options{}

	// --- When ---
	WithMeta(m)(opts)

	// --- Then ---
	assert.Same(t, m, opts.init)
}

func Test_WithTags(t *testing.T) {
	// --- Given ---
	m := map[string]Tag{"A": nil}
	opts := &Options{}

	// --- When ---
	WithTags(m)(opts)

	// --- Then ---
	assert.Same(t, m, opts.init)
}

func Test_WithTimeFormat(t *testing.T) {
	// --- Given ---
	opts := &Options{}

	// --- When ---
	WithTimeFormat(time.RFC822)(opts)

	// --- Then ---
	assert.Equal(t, time.RFC822, opts.TimeFormat)
}

func Test_WithTimeLoc(t *testing.T) {
	// --- Given ---
	WAW := must.Value(time.LoadLocation("Europe/Warsaw"))
	opts := &Options{}

	// --- When ---
	WithTimeLoc(WAW)(opts)

	// --- Then ---
	assert.Same(t, WAW, opts.Location)
}

func Test_WithLocString(t *testing.T) {
	// --- Given ---
	opts := &Options{}

	// --- When ---
	WithLocString(opts)

	// --- Then ---
	assert.True(t, opts.LocationAsString)
}

func Test_WithZeroTime(t *testing.T) {
	// --- Given ---
	opts := &Options{}

	// --- When ---
	WithZeroTime("a", "b", "c")(opts)

	// --- Then ---
	assert.Equal(t, []string{"a", "b", "c"}, opts.zeroTime)
}

func Test_WithRadixHEX(t *testing.T) {
	// --- Given ---
	opts := &Options{}

	// --- When ---
	WithRadixHEX(opts)

	// --- Then ---
	assert.Equal(t, 16, opts.Radix)
}
