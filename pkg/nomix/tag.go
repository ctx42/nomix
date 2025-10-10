// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"time"
)

// TagType represents supported tag types.
type TagType interface {
	byte | string | int64 | float64 | time.Time |
		int | bool | time.Duration
}

// TagKind describes the type of [Tag] value.
type TagKind uint16

// Base [Tag] kinds.
const (
	KindString  TagKind = 0b00000000_00000010
	KindInt64   TagKind = 0b00000000_00000100
	KindFloat64 TagKind = 0b00000000_00001000
	KindTime    TagKind = 0b00000000_00010000
	KindJSON    TagKind = 0b00000000_00100000
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
	KindBoolSlice    = KindBool | KindSlice
	KindIntSlice     = KindInt | KindSlice
)

// KindSlice is a [TagKind] type modifier indicating it is a slice.
const KindSlice TagKind = 0b10000000_00000000

// Tag is an interface representing a tag.
//
// Tags are named and typed values that can be used to annotate objects.
type Tag interface {
	// TagName returns tag name.
	TagName() string

	// TagKind returns the [TagKind] holding the information about the type of
	// the tag value. Use it interprets the value returned by the [Tag.TagValue]
	// method.
	TagKind() TagKind

	// TagValue returns tag value.
	// You may use the value returned by the [Tag.TagKind] method
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

// TagAllGetter is an interface for retrieving all tags in the set.
type TagAllGetter interface {
	// TagGetAll returns all tags in the set as a map. May return nil. The
	// returned map should be treated as read-only.
	TagGetAll() map[string]Tag
}

// TagComparer is an interface for comparing tags.
type TagComparer interface {
	// TagSame returns true if both tags have the same name, kind and value.
	TagSame(other Tag) bool
}

// TagValueComparer is an interface for comparing tag values.
type TagValueComparer interface {
	// TagEqual returns true if both tags are having the same kind and value.
	TagEqual(other Tag) bool
}
