// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"encoding/json"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_JSONSpec(t *testing.T) {
	// --- When ---
	have := JSONSpec()

	// --- Then ---
	tag, err := have.TagCreate("name", []byte(`{}`))
	assert.NoError(t, err)
	assert.SameType(t, &JSON{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []byte(`{}`), tag.TagValue())
	assert.Equal(t, KindJSON, tag.TagKind())

	tag, err = have.TagParse("name", `{}`)
	assert.NoError(t, err)
	assert.SameType(t, &JSON{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []byte(`{}`), tag.TagValue())
	assert.Equal(t, KindJSON, tag.TagKind())
}

func Test_NewJSON(t *testing.T) {
	// --- When ---
	tag := NewJSON("name", []byte(`{}`))

	// --- Then ---
	assert.SameType(t, &JSON{}, tag)
	assert.Equal(t, "name", tag.name)
	assert.Equal(t, []byte(`{}`), tag.value)
	assert.Equal(t, KindJSON, tag.kind)
	assert.Equal(t, `{}`, tag.String())
}

func Test_CreateJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		tag, err := CreateJSON("name", []byte(`{}`))

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &JSON{}, tag)
		assert.Equal(t, "name", tag.name)
		assert.Equal(t, []byte(`{}`), tag.value)
		assert.Equal(t, KindJSON, tag.kind)
		assert.Equal(t, `{}`, tag.String())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		tag, err := CreateJSON("name", "abc")

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, tag)
	})
}

func Test_createJSON(t *testing.T) {
	t.Run("valid json.RawMessage", func(t *testing.T) {
		// --- When ---
		have, err := createJSON(json.RawMessage(`{"A": 1}`), Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, json.RawMessage(`{"A": 1}`), have)
	})

	t.Run("valid []byte", func(t *testing.T) {
		// --- When ---
		have, err := createJSON([]byte(`{"A": 1}`), Options{})

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, json.RawMessage(`{"A": 1}`), have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := createJSON("abc", Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})

	t.Run("error - invalid format", func(t *testing.T) {
		// --- When ---
		have, err := createJSON([]byte(`{!!!}`), Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := createJSON(nil, Options{})

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.Nil(t, have)
	})
}

func Test_ParseJSON_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		str  string
		have json.RawMessage
	}{
		{
			"object",
			`{"a": 2, "b": 3}`,
			json.RawMessage(`{"a": 2.0, "b": 3.0}`),
		},
		{
			"array",
			`[1, 2, 3]`,
			json.RawMessage(`[1, 2, 3]`),
		},
		{
			"object with an array field",
			`{"a": 2, "b": [3,4]}`,
			json.RawMessage(`{"a": 2.0, "b": [3.0, 4.0]}`),
		},
		{
			"object with an object field",
			`{"a": 2, "b": {"c": 3}}`,
			json.RawMessage(`{"a": 2.0, "b": {"c": 3.0}}`),
		},
	}

	for _, tc := range tt {
		t.Run(tc.str, func(t *testing.T) {
			// --- When ---
			tag, err := ParseJSON("name", tc.str)

			// --- Then ---
			assert.NoError(t, err)
			assert.JSON(t, string(tc.have), tag.String())
		})
	}
}

func Test_ParseJSON(t *testing.T) {
	t.Run("error - not supported string value", func(t *testing.T) {
		// --- When ---
		tag, err := ParseJSON("name", "bad")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, ErrInvFormat, err)
		assert.Nil(t, tag)
	})
}
