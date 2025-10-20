// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"encoding/json"
	"fmt"
)

// JSON is a tag for a [json.RawMessage] value.
type JSON = Slice[byte]

// jsonSpec defines the [KindSpec] for [JSON] type.
var jsonSpec = KindSpec{
	knd: KindJSON,
	tcr: CreateFunc(CreateJSON),
	tpr: ParseFunc(ParseJSON),
}

// JSONSpec returns a [KindSpec] for [JSON] type.
func JSONSpec() KindSpec { return jsonSpec }

// NewJSON returns a new instance of [JSON].
func NewJSON(name string, v json.RawMessage) *JSON {
	return NewSlice(name, v, KindJSON, strValueJSON, nil)
}

// CreateJSON casts the value to [json.RawMessage]. Returns the [JSON] with the
// specified name and nil error if the value is a valid JSON represented as
// [json.RawMessage], []byte or string. Returns nil and error if the value's
// type is not a supported type or the value is not a valid JSON.
func CreateJSON(name string, val any, _ ...Option) (*JSON, error) {
	vv, err := createJSON(val, defaultOptions)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	return NewJSON(name, vv), nil
}

// createJSON casts the value to [json.RawMessage]. Returns the cast value and
// nil error if the value is a valid JSON represented as [json.RawMessage],
// []byte or string. Returns nil and error if the value's type is not a
// supported type or the value is not a valid JSON.
func createJSON(val any, _ Options) (json.RawMessage, error) {
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

// strValueJSON converts [json.RawMessage] to its string representation.
func strValueJSON(v []byte) string {
	return string(v)
}

// ParseJSON parses string representation of the raw [JSON] tag.
func ParseJSON(name, v string, _ ...Option) (*JSON, error) {
	if !json.Valid(json.RawMessage(v)) {
		return nil, fmt.Errorf("%s: %w", name, ErrInvFormat)
	}
	return NewJSON(name, json.RawMessage(v)), nil
}
