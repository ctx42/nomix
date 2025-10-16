// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

// Definition represents a named tag definition. In other words, it wraps a
// [KindSpec] and a tag name.
type Definition struct {
	name string   // Tag name.
	spec KindSpec // Tag specification.
}

// Define defines named [Tag].
func Define(name string, spec KindSpec) *Definition {
	return &Definition{name: name, spec: spec}
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
