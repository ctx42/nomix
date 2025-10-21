// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"github.com/ctx42/verax/pkg/verax"
	"github.com/ctx42/xrr/pkg/xrr"
)

// Definition represents a named tag definition. In other words, it wraps a
// [Spec] and a tag name.
type Definition struct {
	name string     // Tag name.
	spec Spec       // Tag specification.
	rule verax.Rule // Optional validation rule.
}

// Define defines named [Tag].
func Define(name string, spec Spec, rules ...verax.Rule) *Definition {
	def := &Definition{
		name: name,
		spec: spec,
	}
	if len(rules) > 0 {
		def.rule = verax.Set(rules)
	}
	return def
}

// TagName returns the tag definition name.
func (def *Definition) TagName() string { return def.name }

// TagKind returns the tag definition kind.
func (def *Definition) TagKind() Kind { return def.spec.knd }

// TagCreate creates a new [Tag] matching the definition. It does not validate
// the value.
func (def *Definition) TagCreate(val any, opts ...Option) (Tag, error) {
	tag, err := def.spec.TagCreate(def.name, val, opts...)
	if err != nil {
		return nil, err
	}
	if err = def.validate(tag.TagValue()); err != nil {
		return nil, err
	}
	return tag, nil
}

func (def *Definition) TagParse(val string, opts ...Option) (Tag, error) {
	tag, err := def.spec.TagParse(def.name, val, opts...)
	if err != nil {
		return nil, err
	}
	if err = def.validate(tag.TagValue()); err != nil {
		return nil, err
	}
	return tag, nil
}

// Validate validates the given value against the definition.
//
// NOTE: The [Creator] is first used to create a [Tag] instance with the
// provided value; hence all types supported by [Creator] are supported.
func (def *Definition) Validate(val any) error {
	_, err := def.TagCreate(val) // Also does validation.
	return err
}

func (def *Definition) validate(val any) error {
	if def.rule == nil {
		return nil
	}
	if err := def.rule.Validate(val); err != nil {
		return xrr.FieldError(def.name, err)
	}
	return nil
}
