// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"database/sql/driver"

	"github.com/ctx42/verax/pkg/verax"
	"github.com/ctx42/xrr/pkg/xrr"
)

// Compile time checks.
var (
	_ Tag           = &Slice[int]{}
	_ ValueComparer = &Slice[int]{}
	_ Comparer      = &Slice[int]{}
)

// Slice is a generic type for multi value [Tag].
type Slice[T comparable] struct {
	name      string                          // Tag name.
	value     []T                             // Tag value.
	kind      Kind                            // Tag kind.
	strValuer func([]T) string                // T to string function.
	sqlValuer func([]T) (driver.Value, error) // T to SQL value function.
}

// NewSlice returns a new instance of [Slice].
func NewSlice[T comparable](
	name string,
	val []T,
	kind Kind,
	strValuer func([]T) string,
	sqlValuer func([]T) (driver.Value, error),
) *Slice[T] {

	return &Slice[T]{
		name:      name,
		value:     val,
		kind:      kind,
		strValuer: strValuer,
		sqlValuer: sqlValuer,
	}
}

func (tag *Slice[T]) TagName() string { return tag.name }
func (tag *Slice[T]) TagKind() Kind   { return tag.kind }
func (tag *Slice[T]) TagValue() any   { return tag.value }

func (tag *Slice[T]) TagSet(v any) error {
	if v, ok := v.([]T); ok {
		tag.value = v
		return nil
	}
	return ErrInvType
}

func (tag *Slice[T]) Get() []T  { return tag.value }
func (tag *Slice[T]) Set(v []T) { tag.value = v }

func (tag *Slice[T]) Value() (driver.Value, error) {
	if tag.sqlValuer == nil {
		return tag.value, nil
	}
	return tag.sqlValuer(tag.value)
}

func (tag *Slice[T]) TagEqual(other Tag) bool {
	if other == nil {
		return false
	}
	o, ok := other.(*Slice[T])
	if !ok {
		return false
	}
	if len(tag.value) != len(o.value) {
		return false
	}
	for i, v := range tag.value {
		if v != o.value[i] {
			return false
		}
	}
	return true
}

func (tag *Slice[T]) TagSame(other Tag) bool {
	if other == nil {
		return false
	}
	o, ok := other.(*Slice[T])
	if !ok {
		return false
	}
	if tag.name != o.name {
		return false
	}
	if len(tag.value) != len(o.value) {
		return false
	}
	for i, v := range tag.value {
		if v != o.value[i] {
			return false
		}
	}
	return true
}

func (tag *Slice[T]) String() string { return tag.strValuer(tag.value) }

func (tag *Slice[T]) ValidateWith(rule verax.Rule) error {
	if err := rule.Validate(tag.value); err != nil {
		return xrr.FieldError(tag.name, err)
	}
	return nil
}
