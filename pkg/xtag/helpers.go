// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac.
// SPDX-License-Identifier: MIT

package xtag

import (
	"strings"
)

// formatSlice formats the slice as "[e0, e1, ...]", converting each element
// to a string with the format function.
func formatSlice[T any](v []T, format func(T) string) string {
	var buf strings.Builder
	buf.WriteByte('[')
	for i, val := range v {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(format(val))
	}
	buf.WriteByte(']')
	return buf.String()
}
