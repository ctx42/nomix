// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"github.com/ctx42/verax/pkg/verax"
	"github.com/ctx42/xrr/pkg/xrr"
)

// Compile time checks.
var (
	_ Tag              = &slice[int]{}
	_ TagValueComparer = &slice[int]{}
	_ TagComparer      = &slice[int]{}
)

// String represents string tag.
type slice[T TagType] struct {
	name     string
	value    []T
	kind     TagKind
	stringer func([]T) string
}

func (tag *slice[T]) TagName() string  { return tag.name }
func (tag *slice[T]) TagKind() TagKind { return tag.kind }
func (tag *slice[T]) TagValue() any    { return tag.value }

func (tag *slice[T]) Value() []T { return tag.value }
func (tag *slice[T]) Set(v []T)  { tag.value = v }

func (tag *slice[T]) TagSet(v any) error {
	if v, ok := v.([]T); ok {
		tag.value = v
		return nil
	}
	return ErrInvType
}

func (tag *slice[T]) TagEqual(other Tag) bool {
	if other == nil {
		return false
	}
	o, ok := other.(*slice[T])
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

func (tag *slice[T]) TagSame(other Tag) bool {
	if other == nil {
		return false
	}
	o, ok := other.(*slice[T])
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

func (tag *slice[T]) String() string { return tag.stringer(tag.value) }

func (tag *slice[T]) ValidateWith(rule verax.Rule) error {
	if err := rule.Validate(tag.value); err != nil {
		return xrr.FieldError(tag.name, err)
	}
	return nil
}
