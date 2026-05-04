// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"

	"github.com/ctx42/verax/pkg/spec"
)

// KindSpecName identifies [KindSpec] in a [spec.Spec].
const KindSpecName = "kind-spec"

// KindSpec is a specification for a tag kind.
type KindSpec struct {
	knd Kind       // Tag kind.
	tcr CreateFunc // Tag creation function.
	tpr ParseFunc  // Tag parsing function.
}

// NewKindSpec creates a new [KindSpec] instance.
func NewKindSpec(knd Kind, tcr CreateFunc, tpr ParseFunc) KindSpec {
	return KindSpec{knd: knd, tcr: tcr, tpr: tpr}
}

// TagCreate creates a new [Tag] matching the [Kind] in the spec.
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

// TagKind returns the [Kind] the spec is for.
func (ks KindSpec) TagKind() Kind { return ks.knd }

func (ks KindSpec) IsZero() bool {
	return ks.knd == 0 && ks.tcr == nil && ks.tpr == nil
}

func (ks KindSpec) Spec() (*spec.Spec, error) {
	spc := spec.NewSpec(KindSpecName).SetArg(spec.ArgValue, int16(ks.knd))
	return spc, nil
}

// KindSpecFromSpec creates a [KindSpec] from a [spec.Spec]. The [Kind] encoded
// in the spec must be registered in the global [Registry] at the time this
// function is called.
func KindSpecFromSpec(reg *Registry, spc *spec.Spec) (KindSpec, error) {
	if spc.Name != KindSpecName {
		msg := fmt.Sprintf("%s: invalid spec name: %q", KindSpecName, spc.Name)
		return KindSpec{}, NewInternalError(msg, spec.ECInvSpec)
	}
	iKnd, err := getSpecArg[int16](spc.Args, spec.ArgValue, KindSpecName)
	if err != nil {
		return KindSpec{}, err
	}

	ks := reg.SpecForKind(Kind(iKnd))
	if ks.IsZero() {
		msg := fmt.Sprintf("%s: invalid spec kind: %d", KindSpecName, iKnd)
		return KindSpec{}, NewInternalError(msg, spec.ECInvSpec)
	}
	return ks, nil
}
