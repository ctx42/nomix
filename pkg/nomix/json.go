// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"encoding/json"
	"fmt"
)

// JSON is a tag for a [json.RawMessage] value.
type JSON = slice[byte]

// NewJSON returns a new instance of [JSON].
func NewJSON(name string, v json.RawMessage) *JSON {
	return &slice[byte]{
		name:     name,
		value:    v,
		kind:     KindJSON,
		stringer: jsonToString,
	}
}

// jsonToString converts [json.RawMessage] to its string representation.
func jsonToString(v []byte) string {
	return string(v)
}

// ParseJSON parses string representation of the raw [JSON] tag.
func ParseJSON(name, v string, _ ...Option) (*JSON, error) {
	if !json.Valid(json.RawMessage(v)) {
		return nil, fmt.Errorf("%s: %w", name, ErrInvFormat)
	}
	return NewJSON(name, json.RawMessage(v)), nil
}

// asJSON casts the value to [json.RawMessage] and verifies it is valid JSON.
// Returns the value and nil error if successful. Returns nil and [ErrInvType]
// if the value is not a valid JSON type.
func asJSON(val any, _ Options) (json.RawMessage, error) {
	var vv json.RawMessage
	switch v := val.(type) {
	case json.RawMessage:
		vv = v
	case []byte:
		vv = v
	default:
		return nil, ErrInvType
	}
	if !json.Valid(vv) {
		return nil, ErrInvFormat
	}
	return vv, nil
}
