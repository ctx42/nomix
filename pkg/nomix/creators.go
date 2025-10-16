// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// creators is the global collection of creators for registered types.
var creators Creators

// RegisterCreator registers a tag creator for the given type. If the type
// already exists, it is overwritten.
func RegisterCreator(typ any, creator TagCreator) {
	creators.Register(typ, creator)
}

// CreateTag creates a new [Tag] for the given value. The value's type must be
// registered.
func CreateTag(name string, val any) (Tag, error) {
	return creators.Create(name, val)
}

func init() {
	creators = NewCreators()

	creators.Register(byte(1), CreateFunc(CreateInt64))
	creators.Register(int(1), CreateFunc(CreateInt))
	creators.Register(int8(1), CreateFunc(CreateInt64))
	creators.Register(int16(1), CreateFunc(CreateInt64))
	creators.Register(int32(1), CreateFunc(CreateInt64))
	creators.Register(int64(1), CreateFunc(CreateInt64))

	creators.Register(float32(1), CreateFunc(CreateFloat64))
	creators.Register(float64(1), CreateFunc(CreateFloat64))

	creators.Register(true, CreateFunc(CreateBool))
	creators.Register("string", CreateFunc(CreateString))
	creators.Register(json.RawMessage{}, CreateFunc(CreateJSON))
	creators.Register(time.Time{}, CreateFunc(CreateTime))

	creators.Register([]byte{}, CreateFunc(CreateByteSlice))
	creators.Register([]int{}, CreateFunc(CreateIntSlice))
	creators.Register([]int8{}, CreateFunc(CreateInt64Slice))
	creators.Register([]int16{}, CreateFunc(CreateInt64Slice))
	creators.Register([]int32{}, CreateFunc(CreateInt64Slice))
	creators.Register([]int64{}, CreateFunc(CreateInt64Slice))

	creators.Register([]float32{}, CreateFunc(CreateFloat64Slice))
	creators.Register([]float64{}, CreateFunc(CreateFloat64Slice))

	creators.Register([]bool{}, CreateFunc(CreateBoolSlice))
	creators.Register([]string{}, CreateFunc(CreateStringSlice))
	creators.Register([]time.Time{}, CreateFunc(CreateTimeSlice))
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
		return creator.TagCreate(name, val, opts...)
	}
	return nil, fmt.Errorf("%s: %w", name, ErrNoCreator)
}

// Creator wraps a [GenTagCreator] to implement the [TagCreator] interface.
type Creator[T Tag] struct{ GenTagCreator[T] }

// CreateFunc creates a new [Creator] based on [GenTagCreator].
func CreateFunc[T Tag](cr GenTagCreator[T]) *Creator[T] {
	return &Creator[T]{cr}
}

func (c *Creator[T]) TagCreate(
	name string,
	value any,
	opts ...Option,
) (Tag, error) {

	return c.GenTagCreator(name, value, opts...)
}
