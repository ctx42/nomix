// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

// Kind describes the type of [Tag] value.
type Kind int16

// String implements [fmt.Stringer].
//
// nolint: cyclop
func (tk Kind) String() string {
	switch tk {
	case KindString:
		return "KindString"
	case KindInt64:
		return "KindInt64"
	case KindFloat64:
		return "KindFloat64"
	case KindTime:
		return "KindTime"
	case KindUUID:
		return "KindUUID"
	case KindJSON:
		return "KindJSON"
	case KindBool:
		return "KindBool"
	case KindInt:
		return "KindInt"
	case KindByteSlice:
		return "KindByteSlice"
	case KindStringSlice:
		return "KindStringSlice"
	case KindInt64Slice:
		return "KindInt64Slice"
	case KindFloat64Slice:
		return "KindFloat64Slice"
	case KindTimeSlice:
		return "KindTimeSlice"
	case KindUUIDSlice:
		return "KindUUIDSlice"
	case KindBoolSlice:
		return "KindBoolSlice"
	case KindIntSlice:
		return "KindIntSlice"
	default:
		return "KindUnknown"
	}
}

// Base [Tag] kinds.
const (
	KindString  Kind = 0b00000000_00000010
	KindInt64   Kind = 0b00000000_00000100
	KindFloat64 Kind = 0b00000000_00001000
	KindTime    Kind = 0b00000000_00010000
	KindJSON    Kind = 0b00000000_00100000
	KindUUID    Kind = 0b00000000_01000000
)

// Derived [Tag] kinds.
// Derived kinds are types that are derived from base kinds.
const (
	KindBool = 0b00000001_00000000 | KindInt64
	KindInt  = 0b00000010_00000000 | KindInt64
)

// Multi value (slice) [Tag] kinds.
const (
	KindByteSlice    = 0b00000000_00000001 | KindSlice
	KindStringSlice  = KindString | KindSlice
	KindInt64Slice   = KindInt64 | KindSlice
	KindFloat64Slice = KindFloat64 | KindSlice
	KindTimeSlice    = KindTime | KindSlice
	KindUUIDSlice    = KindUUID | KindSlice
	KindBoolSlice    = KindBool | KindSlice
	KindIntSlice     = KindInt | KindSlice
)

// KindSlice is a [Kind] type modifier indicating it is a slice.
const KindSlice Kind = 0b00000000_10000000

// Tag is an interface representing a tag.
//
// Tags are named and typed values that can be used to annotate objects.
type Tag interface {
	// TagName returns tag name.
	TagName() string

	// TagKind returns the [Kind] holding the information about the type of
	// the tag value. Use it to interpret the value returned by the
	// [Tag.TagValue] method.
	TagKind() Kind

	// TagValue returns tag value.
	// You may use the value returned by the [Tag.Kind] method
	// to cast it to the proper type.
	TagValue() any
}

// Tagger is an interface for managing a collection of distinctly named [Tag]
// instances (set). The implementations must not allow for nil values to be
// stored in the set.
type Tagger interface {
	// TagGet retrieves from the set a [Tag] by its name. If the name doesn't
	// exist in the set, it returns nil.
	TagGet(name string) Tag

	// TagSet adds instances of [Tag] to the set. If the tag name already
	// exists in the set, it will be overwritten. The nil instances are ignored.
	TagSet(tag ...Tag)

	// TagDelete removes from the set the [Tag] by name. If the name does not
	// exist, the method has no effect.
	TagDelete(name string)
}

// AllGetter is an interface for retrieving all tags in the set.
type AllGetter interface {
	// TagGetAll returns all tags in the set as a map. May return nil. The
	// returned map should be treated as read-only.
	TagGetAll() map[string]Tag
}

// Comparer is an interface for comparing tags.
type Comparer interface {
	// TagSame returns true if both tags have the same name, kind and value.
	TagSame(other Tag) bool
}

// ValueComparer is an interface for comparing tag values.
type ValueComparer interface {
	// TagEqual returns true if both tags are having the same kind and value.
	TagEqual(other Tag) bool
}

// Creator is an interface for creating [Tag] instances.
type Creator interface {
	// TagCreate creates the appropriate [Tag] instance based on the value's
	// type. It returns the [ErrNoCreator] error if the value's type is not
	// supported. The name must be set by the implementer.
	TagCreate(value any, opts ...Option) (Tag, error)
}

// NamedCreator is an interface for creating named [Tag] instances.
type NamedCreator interface {
	// TagCreate creates the appropriate [Tag] instance based on the value's
	// type. It returns the [ErrNoCreator] error if the value's type is not
	// supported.
	TagCreate(name string, value any, opts ...Option) (Tag, error)
}

// Parser is an interface for creating [Tag] instances from their string
// representation.
type Parser interface {
	TagParse(name string, val string, opts ...Option) (Tag, error)
}

// NamedParser is an interface for creating named [Tag] instances from their
// string representation. The name must be set by the implementer.
type NamedParser interface {
	TagParse(val string, opts ...Option) (Tag, error)
}

// CreateFunc function signature for creating [Tag] instances.
type CreateFunc func(name string, val any, opts ...Option) (Tag, error)

// ParseFunc function signature for creating [Tag] instances from their string
// representation.
type ParseFunc func(name, val string, opts ...Option) (Tag, error)

// TagCreateFunc creates a [CreateFunc] from a function that creates concrete
// tag instances.
//
// Examples:
//
//	TagCreateFunc(func(name string, val any, opts ...Option) (MyTag, error))
func TagCreateFunc[T Tag](fn func(name string, val any, opts ...Option) (T, error)) CreateFunc {
	return func(name string, val any, opts ...Option) (Tag, error) {
		return fn(name, val, opts...)
	}
}

// TagParseFunc creates a [ParseFunc] from a function that creates concrete
// tag instances from their string representation.
//
// Examples:
//
//	TagParseFunc(func(name string, val string, opts ...Option) (T, error))
func TagParseFunc[T Tag](fn func(name string, val string, opts ...Option) (T, error)) ParseFunc {
	return func(name string, val string, opts ...Option) (Tag, error) {
		return fn(name, val, opts...)
	}
}
