// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac.
// SPDX-License-Identifier: MIT

package xtag

import (
	"encoding/json"
	"testing"

	"github.com/ctx42/testing/pkg/assert"

	"github.com/ctx42/nomix/pkg/nomix"
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
	assert.Equal(t, nomix.KindJSON, tag.TagKind())

	tag, err = have.TagParse("name", `{}`)
	assert.NoError(t, err)
	assert.SameType(t, &JSON{}, tag)
	assert.Equal(t, "name", tag.TagName())
	assert.Equal(t, []byte(`{}`), tag.TagValue())
	assert.Equal(t, nomix.KindJSON, tag.TagKind())
}

func Test_NewJSON(t *testing.T) {
	// --- When ---
	have := NewJSON("name", []byte(`{}`))

	// --- Then ---
	assert.SameType(t, &JSON{}, have)
	assert.Equal(t, "name", have.TagName())
	assert.Equal(t, []byte(`{}`), have.TagValue())
	assert.Equal(t, nomix.KindJSON, have.TagKind())
	assert.Equal(t, `{}`, have.String())
}

func Test_CreateJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := CreateJSON("name", []byte(`{}`))

		// --- Then ---
		assert.NoError(t, err)
		assert.SameType(t, &JSON{}, have)
		assert.Equal(t, "name", have.TagName())
		assert.Equal(t, []byte(`{}`), have.TagValue())
		assert.Equal(t, nomix.KindJSON, have.TagKind())
		assert.Equal(t, `{}`, have.String())
	})

	t.Run("string", func(t *testing.T) {
		// --- When ---
		have, err := CreateJSON("name", `{"A": 1}`)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"A": 1}`), have.TagValue())
		assert.Equal(t, nomix.KindJSON, have.TagKind())
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := CreateJSON("name", 42)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.ErrorContain(t, "name: ", err)
		assert.Nil(t, have)
	})
}

func Test_createJSON(t *testing.T) {
	t.Run("valid json.RawMessage", func(t *testing.T) {
		// --- When ---
		have, err := createJSON(json.RawMessage(`{"A": 1}`))

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, json.RawMessage(`{"A": 1}`), have)
	})

	t.Run("valid []byte", func(t *testing.T) {
		// --- When ---
		have, err := createJSON([]byte(`{"A": 1}`))

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, json.RawMessage(`{"A": 1}`), have)
	})

	t.Run("valid string", func(t *testing.T) {
		// --- When ---
		have, err := createJSON(`{"A": 1}`)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, json.RawMessage(`{"A": 1}`), have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- When ---
		have, err := createJSON(42)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.Nil(t, have)
	})

	t.Run("error - invalid format", func(t *testing.T) {
		// --- When ---
		have, err := createJSON([]byte(`{!!!}`))

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvFormat, err)
		assert.Nil(t, have)
	})

	t.Run("error - invalid string format", func(t *testing.T) {
		// --- When ---
		have, err := createJSON("abc")

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvFormat, err)
		assert.Nil(t, have)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- When ---
		have, err := createJSON(nil)

		// --- Then ---
		assert.ErrorIs(t, nomix.ErrInvType, err)
		assert.Nil(t, have)
	})
}

func Test_ParseJSON(t *testing.T) {
	t.Run("error - not supported string value", func(t *testing.T) {
		// --- When ---
		have, err := ParseJSON("name", "bad")

		// --- Then ---
		assert.ErrorEqual(t, "name: invalid element format", err)
		assert.ErrorIs(t, nomix.ErrInvFormat, err)
		assert.Nil(t, have)
	})
}

func Test_ParseJSON_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		str  string
		want json.RawMessage
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
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := ParseJSON("name", tc.str)

			// --- Then ---
			assert.NoError(t, err)
			assert.JSON(t, string(tc.want), have.String())
		})
	}
}
