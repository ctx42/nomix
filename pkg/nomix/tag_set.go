// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
)

// TagSet represents a set of tags.
type TagSet struct {
	m map[string]Tag
}

// NewTagSet returns a new instance of [TagSet].
func NewTagSet(opts ...func(*Options)) TagSet {
	def := &Options{length: 10}
	for _, opt := range opts {
		opt(def)
	}
	set := TagSet{}
	if m, ok := def.init.(map[string]Tag); ok {
		for k, v := range m {
			if v == nil {
				delete(m, k)
			}
		}
		set.m = m
	}
	if set.m == nil {
		set.m = make(map[string]Tag, def.length)
	}
	return set
}

func (set TagSet) TagGet(name string) Tag {
	return set.m[name]
}

func (set TagSet) TagSet(tags ...Tag) {
	for _, tag := range tags {
		if tag == nil {
			continue
		}
		set.m[tag.TagName()] = tag
	}
}

func (set TagSet) TagDelete(name string) {
	delete(set.m, name)
}

// TagCount returns the number of entries in the tag set.
func (set TagSet) TagCount() int {
	return len(set.m)
}

func (set TagSet) TagGetAll() map[string]Tag { return set.m }

// TagDeleteAll deletes all tags from the set.
func (set TagSet) TagDeleteAll() {
	for name := range set.m {
		delete(set.m, name)
	}
}

func (set TagSet) MetaGetAll() map[string]any {
	if len(set.m) == 0 {
		return nil
	}
	m := make(map[string]any, len(set.m))
	for k, v := range set.m {
		m[k] = v.TagValue()
	}
	return m
}

// TagGetString gets a tag by name as a [String]. Returns the tag and nil error
// if it exists and is the [KindString] kind. Returns nil and [ErrMissing] if
// the tag is missing, or nil and [ErrInvType] if the tag is of a different
// kind.
func (set TagSet) TagGetString(name string) (*String, error) {
	return getTag[*String](set, name)
}

// TagGetStringSlice gets a tag by name as a [StringSlice]. Returns the tag and
// nil error if it exists and is the [KindStringSlice] kind. Returns nil and
// [ErrMissing] if the tag is missing, or nil and [ErrInvType] if the tag is of
// a different kind.
func (set TagSet) TagGetStringSlice(name string) (*StringSlice, error) {
	return getTag[*StringSlice](set, name)
}

// TagGetInt64 gets a tag by name as a [Int64]. Returns the tag and nil error
// if it exists and is the [KindInt64] kind. Returns nil and [ErrMissing] if
// the tag is missing, or nil and [ErrInvType] if the tag is of a different
// kind.
func (set TagSet) TagGetInt64(name string) (*Int64, error) {
	return getTag[*Int64](set, name)
}

// TagGetInt64Slice gets a tag by name as a [Int64Slice]. Returns the tag and
// nil error if it exists and is the [KindInt64Slice] kind. Returns nil and
// [ErrMissing] if the tag is missing, or nil and [ErrInvType] if the tag is of
// a different kind.
func (set TagSet) TagGetInt64Slice(name string) (*Int64Slice, error) {
	return getTag[*Int64Slice](set, name)
}

// TagGetFloat64 gets a tag by name as a [Float64]. Returns the tag and nil
// error if it exists and is the [KindFloat64] kind. Returns nil and
// [ErrMissing] if the tag is missing, or nil and [ErrInvType] if the tag is of
// a different kind.
func (set TagSet) TagGetFloat64(name string) (*Float64, error) {
	return getTag[*Float64](set, name)
}

// TagGetFloat64Slice gets a tag by name as a [Float64Slice]. Returns the tag
// and nil error if it exists and is the [KindFloat64Slice] kind. Returns nil
// and [ErrMissing] if the tag is missing, or nil and [ErrInvType] if the tag
// is of a different kind.
func (set TagSet) TagGetFloat64Slice(name string) (*Float64Slice, error) {
	return getTag[*Float64Slice](set, name)
}

// TagGetTime gets a tag by name as a [Time]. Returns the tag and nil error if
// it exists and is the [KindTime] kind. Returns nil and [ErrMissing] if the
// tag is missing, or nil and [ErrInvType] if the tag is of a different kind.
func (set TagSet) TagGetTime(name string) (*Time, error) {
	return getTag[*Time](set, name)
}

// TagGetTimeSlice gets a tag by name as a [TimeSlice]. Returns the tag
// and nil error if it exists and is the [KindTimeSlice] kind. Returns nil
// and [ErrMissing] if the tag is missing, or nil and [ErrInvType] if the tag
// is of a different kind.
func (set TagSet) TagGetTimeSlice(name string) (*TimeSlice, error) {
	return getTag[*TimeSlice](set, name)
}

// TagGetJSON gets a tag by name as a [JSON]. Returns the tag and nil error if
// it exists and is the [KindUUID] kind. Returns nil and [ErrMissing] if the
// tag is missing, or nil and [ErrInvType] if the tag is of a different kind.
func (set TagSet) TagGetJSON(name string) (*JSON, error) {
	return getTag[*JSON](set, name)
}

// TagGetBool gets a tag by name as a [Bool]. Returns the tag and nil error if
// it exists and is the [KindBool] kind. Returns nil and [ErrMissing] if the
// tag is missing, or nil and [ErrInvType] if the tag is of a different kind.
func (set TagSet) TagGetBool(name string) (*Bool, error) {
	return getTag[*Bool](set, name)
}

// TagGetBoolSlice gets a tag by name as a [BoolSlice]. Returns the tag
// and nil error if it exists and is the [KindBoolSlice] kind. Returns nil
// and [ErrMissing] if the tag is missing, or nil and [ErrInvType] if the tag
// is of a different kind.
func (set TagSet) TagGetBoolSlice(name string) (*BoolSlice, error) {
	return getTag[*BoolSlice](set, name)
}

// TagGetInt gets a tag by name as a [Int]. Returns the tag and nil error if it
// exists and is the [KindInt] kind. Returns nil and [ErrMissing] if the tag is
// missing, or nil and [ErrInvType] if the tag is of a different kind.
func (set TagSet) TagGetInt(name string) (*Int, error) {
	return getTag[*Int](set, name)
}

// TagGetIntSlice gets a tag by name as a [IntSlice]. Returns the tag and nil
// error if it exists and is the [KindIntSlice] kind. Returns nil and
// [ErrMissing] if the tag is missing, or nil and [ErrInvType] if the tag is of
// a different kind.
func (set TagSet) TagGetIntSlice(name string) (*IntSlice, error) {
	return getTag[*IntSlice](set, name)
}

// getTag returns the tag from the set or zero value if it doesn't exist.
func getTag[T any](set TagSet, name string) (T, error) {
	var zero T
	if tag := set.TagGet(name); tag != nil {
		if v, ok := tag.(T); ok {
			return v, nil
		}
		return zero, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return zero, fmt.Errorf("%s: %w", name, ErrMissing)
}
