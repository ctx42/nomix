// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

// String is a tag for a single byte value.
type String = single[string]

// NewString returns a new instance of [String].
func NewString(name, v string) *String {
	return &single[string]{
		name:     name,
		value:    v,
		kind:     KindString,
		stringer: stringToString,
	}
}

// ParseString creates string type tag. Never returns an error.
func ParseString(name, v string, _ ...Option) (*String, error) {
	return NewString(name, v), nil
}

// stringToString converts string to string.
func stringToString(v string) string { return v }
