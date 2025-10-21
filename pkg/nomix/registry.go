// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"reflect"
	"sync"
)

// specs is the global registry of [Spec]s.
var specs *Registry

func init() { specs = NewRegistry() }

// GlobalRegistry returns the global [Registry] instance.
func GlobalRegistry() *Registry { return specs }

// Registry represents a collection of [Spec]s.
type Registry struct {
	kinds map[Kind]Spec
	specs map[reflect.Type]Spec
	mx    sync.RWMutex
}

// NewRegistry returns a new [Registry] instance.
func NewRegistry() *Registry {
	return &Registry{
		kinds: make(map[Kind]Spec),
		specs: make(map[reflect.Type]Spec),
	}
}

// Register registers a [Spec] for the given [Kind]. Returns nil if
// successful, or an error if the kind is already registered. Each kind can
// have only one spec. Must be called before associating Go types with a
// [Spec].
func (reg *Registry) Register(spec Spec) error {
	reg.mx.Lock()
	defer reg.mx.Unlock()

	if _, ok := reg.kinds[spec.knd]; ok {
		format := "spec for %[1]s(%[1]d) already registered"
		return fmt.Errorf(format, spec.knd)
	}
	reg.kinds[spec.knd] = spec
	return nil
}

// Associate links a Go type to the given [Kind], overwriting any existing
// association. Returns the previous kind association or Kind(0) if none.
// Returns an error if no [Spec] is registered for the kind.
func (reg *Registry) Associate(typ any, knd Kind) (Kind, error) {
	reg.mx.Lock()
	defer reg.mx.Unlock()

	spec, ok := reg.kinds[knd]
	if !ok {
		return 0, fmt.Errorf("no spec for %[1]s(%[1]d)", knd)
	}
	rt := reflect.TypeOf(typ)
	was := reg.specs[rt]
	reg.specs[rt] = spec
	return was.knd, nil
}

// SpecForType retrieves the [Spec] for the given type. Requires prior type
// association with a [Spec]. Use [Spec.IsZero] to check if a spec is
// available for the type.
func (reg *Registry) SpecForType(typ any) Spec {
	reg.mx.RLock()
	defer reg.mx.RUnlock()
	return reg.specs[reflect.TypeOf(typ)]
}

// SpecForKind retrieves the [Spec] for the given [Kind]. Use
// [Spec.IsZero] to check if a spec exists for the kind.
func (reg *Registry) SpecForKind(knd Kind) Spec {
	reg.mx.RLock()
	defer reg.mx.RUnlock()
	return reg.kinds[knd]
}

// Create creates a new [Tag] for the given value. The value's type must be
// registered.
func (reg *Registry) Create(name string, val any, opts ...Option) (Tag, error) {
	reg.mx.RLock()
	defer reg.mx.RUnlock()
	valTyp := reflect.TypeOf(val)
	if spec := reg.specs[valTyp]; !spec.IsZero() {
		return spec.tcr(name, val, opts...)
	}
	return nil, fmt.Errorf("%w for %s of type %T", ErrNoCreator, name, val)
}
