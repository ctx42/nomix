// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

// Spec is a specification for a tag kind.
type Spec struct {
	knd Kind       // Tag kind.
	tcr CreateFunc // Tag creation function.
	tpr ParseFunc  // Tag parsing function.
}

// NewSpec creates a new [Spec] instance.
func NewSpec(knd Kind, tcr CreateFunc, tpr ParseFunc) Spec {
	return Spec{knd: knd, tcr: tcr, tpr: tpr}
}

// TagCreate creates a new [Tag] matching the [Kind] in the spec.
func (spc Spec) TagCreate(
	name string,
	value any,
	opts ...Option,
) (Tag, error) {

	return spc.tcr(name, value, opts...)
}

// TagParse creates a [Tag] based on its string representation.
func (spc Spec) TagParse(name, val string, opts ...Option) (Tag, error) {
	return spc.tpr(name, val, opts...)
}

// TagKind returns the [Kind] the spec is for.
func (spc Spec) TagKind() Kind { return spc.knd }

func (spc Spec) IsZero() bool {
	return spc.knd == 0 && spc.tcr == nil && spc.tpr == nil
}
