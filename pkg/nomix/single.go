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
	_ Tag              = &single[int]{}
	_ TagValueComparer = &single[int]{}
	_ TagComparer      = &single[int]{}
)

// String represents string tag.
type single[T TagType] struct {
	name      string
	value     T
	kind      TagKind
	stringer  func(T) string
	sqlValuer func(T) (driver.Value, error)
}

func (tag *single[T]) TagName() string  { return tag.name }
func (tag *single[T]) TagKind() TagKind { return tag.kind }
func (tag *single[T]) TagValue() any    { return tag.value }
func (tag *single[T]) TagSet(v any) error {
	if v, ok := v.(T); ok {
		tag.value = v
		return nil
	}
	return ErrInvType
}

func (tag *single[T]) Get() T  { return tag.value }
func (tag *single[T]) Set(v T) { tag.value = v }

// Value implements [driver.Valuer] interface. When the type has no sqlValuer
// defined, then it returns the value directly and never returns an error.
func (tag *single[T]) Value() (driver.Value, error) {
	if tag.sqlValuer == nil {
		return tag.value, nil
	}
	return tag.sqlValuer(tag.value)
}

func (tag *single[T]) TagEqual(other Tag) bool {
	if other == nil {
		return false
	}
	if o, ok := other.(*single[T]); ok {
		return tag.value == o.value
	}
	return false
}

func (tag *single[T]) TagSame(other Tag) bool {
	if other == nil {
		return false
	}
	if o, ok := other.(*single[T]); ok {
		return tag.value == o.value && tag.name == o.name
	}
	return false
}

func (tag *single[T]) String() string { return tag.stringer(tag.value) }

func (tag *single[T]) ValidateWith(rule verax.Rule) error {
	if err := rule.Validate(tag.value); err != nil {
		return xrr.FieldError(tag.name, err)
	}
	return nil
}
