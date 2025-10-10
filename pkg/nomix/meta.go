package nomix

// MetaAllGetter is an interface wrapping MetaGetAll method.
type MetaAllGetter interface {
	// MetaGetAll returns all tags in the set as a map. May return nil. The
	// returned map should be treated as read-only.
	MetaGetAll() map[string]any
}

// MetaAllSetter is an interface wrapping MetaSetAll method.
type MetaAllSetter interface {
	// MetaSetAll sets metadata on the implementor. If the value with the given
	// name already exists in the set, it will be overwritten. The nil values
	// must be ignored.
	MetaSetAll(map[string]any)
}

// MetaFromGetter is an interface wrapping MetaFromGetter method.
type MetaFromGetter interface {
	// MetaSetFrom sets metadata on the implementor from the [MetaAllGetter].
	MetaSetFrom(src MetaAllGetter)
}

// MetaAppender is an interface wrapping MetaAppend method.
type MetaAppender interface {
	// MetaAppend appends all implementor metadata to the passed map.
	MetaAppend(map[string]any)
}

// Metadata is an interface for managing a collection of distinctly named
// metadata values. The implementations must not allow for nil values to be
// stored in the set.
type Metadata interface {
	// MetaGet retrieves from the set a metadata value by its name. If the name
	// does not exist in the set, it returns nil.
	MetaGet(name string) any

	// MetaSet adds a named value to the set. If the value with the given name
	// already exists in the set, it will be overwritten. Setting a nil value
	// must be implemented as a no-op.
	MetaSet(name string, value any)

	// MetaDelete removes from the set the metadata value by name. If the name
	// does not exist, the method has no effect.
	MetaDelete(name string)
}
