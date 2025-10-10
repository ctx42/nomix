// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/must"
)

func Test_DefaultParsingOpts(t *testing.T) {
	// --- When ---
	have := DefaultOptions()

	// --- Then ---
	assert.Equal(t, 0, have.length)
	assert.Nil(t, have.init)
	assert.Equal(t, time.RFC3339Nano, have.timeFormat)
	assert.Nil(t, have.location)
	assert.False(t, have.locationAsString)
	assert.Empty(t, have.zeroTime)
	assert.Equal(t, 10, have.intBase)
	assert.Fields(t, 7, have)
}

func Test_WithLen(t *testing.T) {
	// --- Given ---
	opts := &Options{}

	// --- When ---
	WithLen(10)(opts)

	// --- Then ---
	assert.Equal(t, 10, opts.length)
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
	m := map[string]Tag{"A": NewInt("A", 1)}
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
	assert.Equal(t, time.RFC822, opts.timeFormat)
}

func Test_WithTimeLoc(t *testing.T) {
	// --- Given ---
	WAW := must.Value(time.LoadLocation("Europe/Warsaw"))
	opts := &Options{}

	// --- When ---
	WithTimeLoc(WAW)(opts)

	// --- Then ---
	assert.Same(t, WAW, opts.location)
}

func Test_WithLocString(t *testing.T) {
	// --- Given ---
	opts := &Options{}

	// --- When ---
	WithLocString(opts)

	// --- Then ---
	assert.True(t, opts.locationAsString)
}

func Test_WithBaseHEX(t *testing.T) {
	// --- Given ---
	opts := &Options{}

	// --- When ---
	WithBaseHEX(opts)

	// --- Then ---
	assert.Equal(t, 16, opts.intBase)
}
