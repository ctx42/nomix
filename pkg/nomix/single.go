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
	_ Tag              = &Single[int]{}
	_ TagValueComparer = &Single[int]{}
	_ TagComparer      = &Single[int]{}
)

// Single is a generic type for single value [Tag].
type Single[T comparable] struct {
	name      string                        // Tag name.
	value     T                             // Tag value.
	kind      TagKind                       // Tag kind.
	strValuer func(T) string                // T to string function.
	sqlValuer func(T) (driver.Value, error) // T to SQL value function.
}

// NewSingle returns a new instance of [Single].
func NewSingle[T comparable](
	name string,
	val T,
	kind TagKind,
	strValuer func(T) string,
	sqlValuer func(T) (driver.Value, error),
) *Single[T] {

	return &Single[T]{
		name:      name,
		value:     val,
		kind:      kind,
		strValuer: strValuer,
		sqlValuer: sqlValuer,
	}
}

func (tag *Single[T]) TagName() string  { return tag.name }
func (tag *Single[T]) TagKind() TagKind { return tag.kind }
func (tag *Single[T]) TagValue() any    { return tag.value }
func (tag *Single[T]) TagSet(v any) error {
	if v, ok := v.(T); ok {
		tag.value = v
		return nil
	}
	return ErrInvType
}

func (tag *Single[T]) Get() T  { return tag.value }
func (tag *Single[T]) Set(v T) { tag.value = v }

// Value implements [driver.Valuer] interface. When the type has no sqlValuer
// defined, then it returns the value directly and never returns an error.
func (tag *Single[T]) Value() (driver.Value, error) {
	if tag.sqlValuer == nil {
		return tag.value, nil
	}
	return tag.sqlValuer(tag.value)
}

func (tag *Single[T]) TagEqual(other Tag) bool {
	if other == nil {
		return false
	}
	if o, ok := other.(*Single[T]); ok {
		return tag.value == o.value
	}
	return false
}

func (tag *Single[T]) TagSame(other Tag) bool {
	if other == nil {
		return false
	}
	if o, ok := other.(*Single[T]); ok {
		return tag.value == o.value && tag.name == o.name
	}
	return false
}

func (tag *Single[T]) String() string { return tag.strValuer(tag.value) }

func (tag *Single[T]) ValidateWith(rule verax.Rule) error {
	if err := rule.Validate(tag.value); err != nil {
		return xrr.FieldError(tag.name, err)
	}
	return nil
}
