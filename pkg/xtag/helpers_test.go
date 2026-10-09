// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac.
// SPDX-License-Identifier: MIT

package xtag

import (
	"strconv"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_formatSlice_tabular(t *testing.T) {
	tt := []struct {
		testN string

		val  []int
		want string
	}{
		{"nil", nil, "[]"},
		{"empty", []int{}, "[]"},
		{"one", []int{42}, "[42]"},
		{"two", []int{42, 44}, "[42, 44]"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have := formatSlice(tc.val, strconv.Itoa)

			// --- Then ---
			assert.Equal(t, tc.want, have)
		})
	}
}
