// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

// StringSlice is a tag for a slice of strings.
type StringSlice = slice[string]

// NewStringSlice returns a new instance of [StringSlice].
func NewStringSlice(name string, v ...string) *StringSlice {
	return &slice[string]{
		name:     name,
		value:    v,
		kind:     KindStringSlice,
		stringer: stringSliceToString,
	}
}

// stringSliceToString converts a string slice to its string representation.
func stringSliceToString(v []string) string {
	ret := "["
	for i, val := range v {
		if i > 0 {
			ret += ", "
		}
		ret += `"` + val + `"`
	}
	return ret + "]"
}

// asStringSlice casts the value to []string. Returns the slice and nil error
// if the value is a []string. Returns nil and [ErrInvType] if not a []string.
func asStringSlice(val any, _ *Options) ([]string, error) {
	switch v := val.(type) {
	case []string:
		return v, nil
	default:
		return nil, ErrInvType
	}
}
