// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// specs is the global registry of [KindSpec]s.
var specs Registry

// RegisterSpec registers a tag creator for the given type. If the type
// already exists, it is overwritten.
func RegisterSpec(typ any, spec KindSpec) { specs.Register(typ, spec) }

// CreateTag creates a new [Tag] for the given name and value. The value's type
// must be first registered with [RegisterSpec].
func CreateTag(name string, val any) (Tag, error) {
	return specs.Create(name, val)
}

// CreatorForKind returns the [TagCreateFunc] for the given [TagKind].
func CreatorForKind(knd TagKind) (TagCreateFunc, error) {
	return specs.CreatorForKind(knd)
}

// CreatorForType returns the [TagCreateFunc] for the given type.
func CreatorForType(typ any) (TagCreateFunc, error) {
	return specs.CreatorForType(typ)
}

func init() {
	specs = NewRegistry()

	RegisterSpec(byte(1), int64Spec)
	RegisterSpec(int(1), intSpec)
	RegisterSpec(int8(1), int64Spec)
	RegisterSpec(int16(1), int64Spec)
	RegisterSpec(int32(1), int64Spec)
	RegisterSpec(int64(1), int64Spec)
	RegisterSpec(float32(1), float64Spec)
	RegisterSpec(float64(1), float64Spec)

	RegisterSpec(true, boolSpec)
	RegisterSpec("string", stringSpec)
	RegisterSpec(time.Time{}, timeSpec)
	RegisterSpec(json.RawMessage{}, jsonSpec)

	RegisterSpec([]byte{}, byteSliceSpec)
	RegisterSpec([]int{}, intSliceSpec)
	RegisterSpec([]int8{}, int64SliceSpec)
	RegisterSpec([]int16{}, int64SliceSpec)
	RegisterSpec([]int32{}, int64SliceSpec)
	RegisterSpec([]int64{}, int64SliceSpec)
	RegisterSpec([]float32{}, float64SliceSpec)
	RegisterSpec([]float64{}, float64SliceSpec)

	RegisterSpec([]bool{}, boolSliceSpec)
	RegisterSpec([]string{}, stringSliceSpec)
	RegisterSpec([]time.Time{}, timeSliceSpec)
}

// Registry represents a collection of [KindSpec]s.
type Registry struct {
	specs map[reflect.Type]KindSpec
}

// NewRegistry returns a new [Registry] instance.
func NewRegistry() Registry {
	return Registry{
		specs: make(map[reflect.Type]KindSpec),
	}
}

// Register registers a tag creator for the given type. If the type already
// exists, it is overwritten. Returns the previous creator function when the
// type was already registered, nil otherwise.
func (reg Registry) Register(typ any, spec KindSpec) KindSpec {
	rt := reflect.TypeOf(typ)
	have := reg.specs[rt]
	reg.specs[rt] = spec
	return have
}

// Create creates a new [Tag] for the given value. The value's type must be
// registered.
func (reg Registry) Create(name string, val any, opts ...Option) (Tag, error) {
	valTyp := reflect.TypeOf(val)
	for typ, spec := range reg.specs {
		if typ == valTyp {
			return spec.tcr(name, val, opts...)
		}
	}
	return nil, fmt.Errorf("%w for %s of type %T", ErrNoCreator, name, val)
}

// CreatorForKind returns the [TagCreateFunc] for the given [TagKind].
func (reg Registry) CreatorForKind(knd TagKind) (TagCreateFunc, error) {
	for _, spec := range reg.specs {
		if spec.knd == knd {
			return spec.tcr, nil
		}
	}
	return nil, fmt.Errorf("%w: %s (%d)", ErrNoCreator, knd.String(), knd)
}

// CreatorForType returns the [TagCreateFunc] for the given type.
func (reg Registry) CreatorForType(typ any) (TagCreateFunc, error) {
	if tcr, ok := reg.specs[reflect.TypeOf(typ)]; ok {
		return tcr.tcr, nil
	}
	return nil, fmt.Errorf("%w: for type %T", ErrNoCreator, typ)
}
