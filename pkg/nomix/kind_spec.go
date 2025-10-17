// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

// KindSpec is a specification for a tag kind.
type KindSpec struct {
	knd TagKind       // Tag kind.
	tcr TagCreateFunc // Tag creation function.
	tpr TagParseFunc  // Tag parsing function.
}

// NewKindSpec creates a new [KindSpec] instance.
func NewKindSpec(knd TagKind, tcr TagCreateFunc, tpr TagParseFunc) KindSpec {
	return KindSpec{knd: knd, tcr: tcr, tpr: tpr}
}

// TagCreate creates a new [Tag] matching the [tagKind] in the spec.
func (ks KindSpec) TagCreate(
	name string,
	value any,
	opts ...Option,
) (Tag, error) {

	return ks.tcr(name, value, opts...)
}

// TagParse creates a [Tag] based on its string representation.
func (ks KindSpec) TagParse(name, val string, opts ...Option) (Tag, error) {
	return ks.tpr(name, val, opts...)
}

func (ks KindSpec) IsZero() bool {
	return ks.knd == 0 && ks.tcr == nil && ks.tpr == nil
}
