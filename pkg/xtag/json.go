// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package xtag

import (
	"encoding/json"
	"fmt"

	"github.com/ctx42/nomix/pkg/nomix"
)

// JSON is a tag for a [json.RawMessage] value.
type JSON = nomix.Slice[byte]

// jsonSpec defines the [nomix.Spec] for [JSON] type.
var jsonSpec = nomix.NewSpec(
	nomix.KindJSON,
	nomix.TagCreateFunc(CreateJSON),
	nomix.TagParseFunc(ParseJSON),
)

// JSONSpec returns a [nomix.Spec] for [JSON] type.
func JSONSpec() nomix.Spec { return jsonSpec }

// NewJSON returns a new instance of [JSON].
func NewJSON(name string, v json.RawMessage) *JSON {
	return nomix.NewSlice(name, v, nomix.KindJSON, strValueJSON, nil)
}

// CreateJSON casts the value to [json.RawMessage]. Returns the [JSON] with the
// specified name and nil error if the value is a valid JSON represented as
// [json.RawMessage], []byte or string. Returns nil and error if the value's
// type is not a supported type or the value is not a valid JSON.
func CreateJSON(name string, val any, _ ...nomix.Option) (*JSON, error) {
	vv, err := createJSON(val, nomix.Options{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewJSON(name, vv), nil
}

// createJSON casts the value to [json.RawMessage]. Returns the cast value and
// nil error if the value is a valid JSON represented as [json.RawMessage],
// []byte or string. Returns nil and error if the value's type is not a
// supported type or the value is not a valid JSON.
func createJSON(val any, _ nomix.Options) (json.RawMessage, error) {
	var vv json.RawMessage
	switch v := val.(type) {
	case json.RawMessage:
		vv = v
	case []byte:
		vv = v
	default:
		return nil, nomix.ErrInvType
	}
	if !json.Valid(vv) {
		return nil, nomix.ErrInvFormat
	}
	return vv, nil
}

// strValueJSON converts [json.RawMessage] to its string representation.
func strValueJSON(v []byte) string {
	return string(v)
}

// ParseJSON parses string representation of the raw [JSON] tag.
func ParseJSON(name, v string, _ ...nomix.Option) (*JSON, error) {
	if !json.Valid(json.RawMessage(v)) {
		return nil, fmt.Errorf("%s: %w", name, nomix.ErrInvFormat)
	}
	return NewJSON(name, json.RawMessage(v)), nil
}
