// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// creatorRegistry is the global registry.
var creatorRegistry Creators

// RegisterCreator registers a tag creator for the given type. If the type
// already exists, it is overwritten.
func RegisterCreator(typ any, creator TagCreator) {
	creatorRegistry.Register(typ, creator)
}

// CreateTag creates a new [Tag] for the given value. The value's type must be
// registered.
func CreateTag(name string, val any) (Tag, error) {
	return creatorRegistry.Create(name, val)
}

func init() {
	creatorRegistry = NewCreators()

	creatorRegistry.Register(byte(1), AsTagCreator(CreateInt64))
	creatorRegistry.Register(int(1), AsTagCreator(CreateInt))
	creatorRegistry.Register(int8(1), AsTagCreator(CreateInt64))
	creatorRegistry.Register(int16(1), AsTagCreator(CreateInt64))
	creatorRegistry.Register(int32(1), AsTagCreator(CreateInt64))
	creatorRegistry.Register(int64(1), AsTagCreator(CreateInt64))

	creatorRegistry.Register(float32(1), AsTagCreator(CreateFloat64))
	creatorRegistry.Register(float64(1), AsTagCreator(CreateFloat64))

	creatorRegistry.Register(true, AsTagCreator(CreateBool))
	creatorRegistry.Register("string", AsTagCreator(CreateString))
	creatorRegistry.Register(json.RawMessage{}, AsTagCreator(CreateJSON))
	creatorRegistry.Register(time.Time{}, AsTagCreator(CreateTime))

	creatorRegistry.Register([]byte{}, AsTagCreator(CreateByteSlice))
	creatorRegistry.Register([]int{}, AsTagCreator(CreateIntSlice))
	creatorRegistry.Register([]int8{}, AsTagCreator(CreateInt64Slice))
	creatorRegistry.Register([]int16{}, AsTagCreator(CreateInt64Slice))
	creatorRegistry.Register([]int32{}, AsTagCreator(CreateInt64Slice))
	creatorRegistry.Register([]int64{}, AsTagCreator(CreateInt64Slice))

	creatorRegistry.Register([]float32{}, AsTagCreator(CreateFloat64Slice))
	creatorRegistry.Register([]float64{}, AsTagCreator(CreateFloat64Slice))

	creatorRegistry.Register([]bool{}, AsTagCreator(CreateBoolSlice))
	creatorRegistry.Register([]string{}, AsTagCreator(CreateStringSlice))
	creatorRegistry.Register([]time.Time{}, AsTagCreator(CreateTimeSlice))
}

// Creators represent a collection of tag creators for a given type.
type Creators struct {
	m map[reflect.Type]TagCreator
}

// NewCreators returns a new [Creators] instance.
func NewCreators() Creators {
	return Creators{m: make(map[reflect.Type]TagCreator)}
}

// Register registers a tag creator for the given type. If the type already
// exists, it is overwritten. Returns the previous creator function when the
// type was already registered, nil otherwise.
func (reg Creators) Register(typ any, creator TagCreator) TagCreator {
	rt := reflect.TypeOf(typ)
	tcr := reg.m[rt]
	reg.m[rt] = creator
	return tcr
}

// Create creates a new [Tag] for the given value. The value's type must be
// registered.
func (reg Creators) Create(name string, val any, opts ...Option) (Tag, error) {
	if creator, ok := reg.m[reflect.TypeOf(val)]; ok {
		return creator(name, val, opts...)
	}
	return nil, fmt.Errorf("%s: %w", name, ErrNoCreator)
}

// AsTagCreator adapts creation functions to the [TagCreator] interface.
func AsTagCreator[T Tag](cr GenTagCreator[T]) TagCreator {
	return func(name string, val any, opts ...Option) (Tag, error) {
		return cr(name, val, opts...)
	}
}
