// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"
)

// specs is the global registry of [KindSpec]s.
var specs Registry

// mxSpecs protects the [specs] variable.
var mxSpecs sync.RWMutex

// RegisterKind registers a [KindSpec] for the given [TagKind]. Returns nil if
// successful, or an error if the kind is already registered. Each kind can
// have only one spec. Must be called before associating Go types with a
// [KindSpec].
func RegisterKind(spec KindSpec) error {
	mxSpecs.Lock()
	defer mxSpecs.Unlock()
	return specs.Register(spec)
}

// AssociateType links a Go type to the given [TagKind], overwriting any
// existing association. Returns the previous kind association or TagKind(0) if
// none. Returns an error if no [KindSpec] is registered for the kind.
func AssociateType(typ any, knd TagKind) (TagKind, error) {
	mxSpecs.Lock()
	defer mxSpecs.Unlock()
	return specs.Associate(typ, knd)
}

// SpecForType retrieves the [KindSpec] for the given type. Requires prior type
// association with a [KindSpec]. Use [KindSpec.IsZero] to check if a spec is
// available for the type.
func SpecForType(typ any) KindSpec {
	mxSpecs.RLock()
	defer mxSpecs.RUnlock()
	return specs.SpecForType(typ)
}

// SpecForKind retrieves the [KindSpec] for the given [TagKind]. Use
// [KindSpec.IsZero] to check if a spec exists for the kind.
func SpecForKind(knd TagKind) KindSpec {
	mxSpecs.RLock()
	defer mxSpecs.RUnlock()
	return specs.SpecForKind(knd)
}

// CreateTag creates a new [Tag] for the given name and value. The value's type
// must be first registered with [AssociateType].
func CreateTag(name string, val any) (Tag, error) {
	mxSpecs.RLock()
	defer mxSpecs.RUnlock()
	return specs.Create(name, val)
}

func init() {
	specs = NewRegistry()

	mustRegisterKind(int64Spec)
	mustRegisterKind(intSpec)
	mustRegisterKind(float64Spec)
	mustRegisterKind(boolSpec)
	mustRegisterKind(stringSpec)
	mustRegisterKind(timeSpec)
	mustRegisterKind(jsonSpec)
	mustRegisterKind(byteSliceSpec)
	mustRegisterKind(intSliceSpec)
	mustRegisterKind(int64SliceSpec)
	mustRegisterKind(float64SliceSpec)
	mustRegisterKind(boolSliceSpec)
	mustRegisterKind(stringSliceSpec)
	mustRegisterKind(timeSliceSpec)

	mustAssociateType(int(1), KindInt)
	mustAssociateType(int8(1), KindInt64)
	mustAssociateType(int16(1), KindInt64)
	mustAssociateType(int32(1), KindInt64)
	mustAssociateType(int64(1), KindInt64)
	mustAssociateType(float32(1), KindFloat64)
	mustAssociateType(float64(1), KindFloat64)

	mustAssociateType(true, KindBool)
	mustAssociateType("string", KindString)
	mustAssociateType(time.Time{}, KindTime)
	mustAssociateType(json.RawMessage{}, KindJSON)

	mustAssociateType([]byte{}, KindByteSlice)
	mustAssociateType([]int{}, KindIntSlice)
	mustAssociateType([]int8{}, KindInt64Slice)
	mustAssociateType([]int16{}, KindInt64Slice)
	mustAssociateType([]int32{}, KindInt64Slice)
	mustAssociateType([]int64{}, KindInt64Slice)
	mustAssociateType([]float32{}, KindFloat64Slice)
	mustAssociateType([]float64{}, KindFloat64Slice)

	mustAssociateType([]bool{}, KindBoolSlice)
	mustAssociateType([]string{}, KindStringSlice)
	mustAssociateType([]time.Time{}, KindTimeSlice)
}

// Registry represents a collection of [KindSpec]s.
type Registry struct {
	kinds map[TagKind]KindSpec
	specs map[reflect.Type]KindSpec
}

// NewRegistry returns a new [Registry] instance.
func NewRegistry() Registry {
	return Registry{
		kinds: make(map[TagKind]KindSpec),
		specs: make(map[reflect.Type]KindSpec),
	}
}

// Register registers a [KindSpec] for the given [TagKind]. Returns nil if
// successful, or an error if the kind is already registered. Each kind can
// have only one spec. Must be called before associating Go types with a
// [KindSpec].
func (reg Registry) Register(spec KindSpec) error {
	if _, ok := reg.kinds[spec.knd]; ok {
		format := "KindSpec for %[1]s(%[1]d) already registered"
		return fmt.Errorf(format, spec.knd)
	}
	reg.kinds[spec.knd] = spec
	return nil
}

// Associate links a Go type to the given [TagKind], overwriting any existing
// association. Returns the previous kind association or TagKind(0) if none.
// Returns an error if no [KindSpec] is registered for the kind.
func (reg Registry) Associate(typ any, knd TagKind) (TagKind, error) {
	spec, ok := reg.kinds[knd]
	if !ok {
		return 0, fmt.Errorf("no spec for %[1]s(%[1]d)", knd)
	}
	rt := reflect.TypeOf(typ)
	was := reg.specs[rt]
	reg.specs[rt] = spec
	return was.knd, nil
}

// SpecForType retrieves the [KindSpec] for the given type. Requires prior type
// association with a [KindSpec]. Use [KindSpec.IsZero] to check if a spec is
// available for the type.
func (reg Registry) SpecForType(typ any) KindSpec {
	return reg.specs[reflect.TypeOf(typ)]
}

// SpecForKind retrieves the [KindSpec] for the given [TagKind]. Use
// [KindSpec.IsZero] to check if a spec exists for the kind.
func (reg Registry) SpecForKind(knd TagKind) KindSpec {
	return reg.kinds[knd]
}

// Create creates a new [Tag] for the given value. The value's type must be
// registered.
func (reg Registry) Create(name string, val any, opts ...Option) (Tag, error) {
	valTyp := reflect.TypeOf(val)
	if spec := reg.specs[valTyp]; !spec.IsZero() {
		return spec.tcr(name, val, opts...)
	}
	return nil, fmt.Errorf("%w for %s of type %T", ErrNoCreator, name, val)
}

// mustRegisterKind is like [RegisterKind], but panics if there is an error.
func mustRegisterKind(spec KindSpec) {
	if err := RegisterKind(spec); err != nil {
		panic(err)
	}
}

// mustAssociateType is like [AssociateType], but panics if there is an error.
func mustAssociateType(typ any, knd TagKind) TagKind {
	was, err := AssociateType(typ, knd)
	if err != nil {
		panic(err)
	}
	return was
}
