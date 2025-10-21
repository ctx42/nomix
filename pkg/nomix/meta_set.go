// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

var _ Metadata = MetaSet{} // Compile time check.

// MetaSet represents a set of metadata key-values.
type MetaSet struct{ m map[string]any }

// NewMetaSet returns a new [MetaSet] instance. By default, the new map is
// initialized with the length equal to 10.
func NewMetaSet(opts ...Option) MetaSet {
	def := NewOptions(opts...)
	set := MetaSet{}
	if m, ok := def.init.(map[string]any); ok {
		for k, v := range m {
			if v == nil {
				delete(m, k)
			}
		}
		set.m = m
	}
	if set.m == nil {
		set.m = make(map[string]any, def.Length)
	}
	return set
}

func (set MetaSet) MetaGet(key string) any { return set.m[key] }

func (set MetaSet) MetaSet(key string, value any) {
	if value == nil {
		return
	}
	set.m[key] = value
}

func (set MetaSet) MetaDelete(key string) { delete(set.m, key) }

// MetaCount returns the number of entries in the metadata set.
func (set MetaSet) MetaCount() int { return len(set.m) }

func (set MetaSet) MetaGetAll() map[string]any { return set.m }

// MetaDeleteAll deletes all metadata from the set.
func (set MetaSet) MetaDeleteAll() {
	for name := range set.m {
		delete(set.m, name)
	}
}
