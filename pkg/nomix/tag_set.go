// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

// TagSet represents a set of tags.
type TagSet struct {
	m map[string]Tag
}

// NewTagSet returns a new instance of [TagSet].
func NewTagSet(opts ...Option) TagSet {
	def := NewOptions(opts...)
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
		set.m = make(map[string]Tag, def.Length)
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
