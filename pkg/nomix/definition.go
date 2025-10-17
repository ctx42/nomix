// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"github.com/ctx42/verax/pkg/verax"
	"github.com/ctx42/xrr/pkg/xrr"
)

// Definition represents a named tag definition. In other words, it wraps a
// [KindSpec] and a tag name.
type Definition struct {
	name string     // Tag name.
	spec KindSpec   // Tag specification.
	rule verax.Rule // Optional validation rule.
}

// Define defines named [Tag].
func Define(name string, spec KindSpec, rules ...verax.Rule) *Definition {
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
func (td *Definition) TagName() string { return td.name }

// TagKind returns the tag definition kind.
func (td *Definition) TagKind() TagKind { return td.spec.knd }

func (td *Definition) TagCreate(val any, opts ...Option) (Tag, error) {
	return td.spec.TagCreate(td.name, val, opts...)
}

func (td *Definition) TagParse(val string, opts ...Option) (Tag, error) {
	return td.spec.TagParse(td.name, val, opts...)
}

// Validate validates the given value against the definition.
//
// NOTE: The [TagCreator] is first used to create a [Tag] instance with the
// provided value this means that all supported by [TagCreator] types are
// supported.
func (td *Definition) Validate(val any) error {
	tag, err := td.TagCreate(val)
	if err != nil {
		return xrr.FieldError(td.name, err)
	}
	if err = td.rule.Validate(tag.TagValue()); err != nil {
		return xrr.FieldError(td.name, err)
	}
	return nil
}
