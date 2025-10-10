[![Go Report Card](https://goreportcard.com/badge/github.com/ctx42/nomix)](https://goreportcard.com/report/github.com/ctx42/nomix)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/ctx42/nomix)
![Tests](https://github.com/ctx42/nomix/actions/workflows/go.yml/badge.svg?branch=master)

# nomix: Tagging and Metadata for Go

<!-- TOC -->
* [nomix: Tagging and Metadata for Go](#nomix-tagging-and-metadata-for-go)
  * [Installation](#installation)
  * [Introduction](#introduction)
    * [Kinds](#kinds)
    * [Slices](#slices)
  * [Instances](#instances)
  * [Tag Set](#tag-set)
  * [Meta Set](#meta-set)
<!-- TOC -->

In the world of software engineering, we often deal with a vast array of
assets like files, user profiles, products in an online store. These assets 
form the backbone of application data structures, but as systems grow more 
complex, simply storing the core data isn't enough. That's where metadata comes 
into play. Metadata provides a way to add extra layers of information and 
organization to these assets, making systems more adaptable. 

`nomix` – from nomos (law/order) or “nomen” (name).

## Installation

To use `nomix` in your Go project, install it with:

```bash
go get github.com/ctx42/nomix
```

## Introduction

The `nomix` module distinguishes two main types of information: 

- **metadata** is a set of named values of any type (`map[string]any`).   
- **tags** is a set of named values each with a specific type (`map[string]Tag`).

The `metadata` values are more flexible where `tags` are more restricted values
that implement the `Tag` interface.

```go
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
```

### Kinds

The `TagKind` identifies the underlying Go type and can be used to decide what 
database type to map it to. The `nomix` module defines and implements the 
following set of _base kinds_:

- `nomix.KindString` - in Go represented by `string` type.
- `nomix.KindInt64`  - in Go represented by `int64` type.
- `nomix.KindFloat64`  - in Go represented by `float64` type.
- `nomix.KindTime`  - in Go represented by `time.Time` type.
- `nomix.KindJSON`  - in Go represented by `json.RawMessage` type.

The _base kinds_ represent known Go types and at the same time reflect the most 
common database types. The idea is to provide the small set of base types and 
`derive` the other types from them.

The derived kinds implemented in `nomix` are: 

- `nomix.KindBool` - in Go represented by `bool` type.
- `nomix.KindInt` - in Go represented by `int` type.

### Slices

All the kinds (except `nomix.KindJSON`) have the `SliceKind` equivalents.

- `nomix.KindByteSlice`  - in Go represented by `[]byte` type.
- `nomix.KindStringSlice` - in Go represented by `[]string` type.
- `nomix.KindInt64Slice` - in Go represented by `[]int64` type.
- `nomix.KindFloat64Slice` - in Go represented by `[]float64` type.
- `nomix.KindTimeSlice` - in Go represented by `[]time.Time` type.
- `nomix.KindBoolSlice` - in Go represented by `[]bool` type.
- `nomix.KindIntSlice` - in Go represented by `[]int` type.

## Instances

Tge `nomix` module provides a `Tag` interface implementations for all the kinds:

- `nomix.String` and `nomix.StringSlice` 
- `nomix.Int64` and `nomix.Int64Slice` 
- `nomix.Float64` and `nomix.Float64Slice` 
- `nomix.Time` and `nomix.TimeSlice` 
- `nomix.Bool` and `nomix.BoolSlice` 
- `nomix.Int` and `nomix.IntSlice` 
- `nomix.ByteSlice` 

## Tag Set

The `Tag` interface provides the basis for operating on sets of differently 
typed and named values in a generic way.

Anything can become a set of tags by implementing the `Tagger` interface.

```go
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
``` 

But `nomix` module provides a `TagSet` type that implements the `Tagger` and 
adds some additional functionality around getting typed values from the set. 

```go
set := NewTagSet()

set.TagSet(NewInt("A", 42), NewBool("B", true), NewString("C", "foo"))

fmt.Printf("There are %d tags in the set:\n", set.TagCount())
fmt.Printf("- A: %v\n", set.TagGet("A").TagValue())
fmt.Printf("- B: %v\n", set.TagGet("B").TagValue())
fmt.Printf("- C: %v\n", set.TagGet("C").TagValue())

fmt.Printf("\nGetting typed tags:\n")

// Tag exists but is of a different type.
tagA, err := set.TagGetInt64("A")
fmt.Printf("  A: %v; err: %v\n", tagA, err)

tagC, err := set.TagGetString("C")
fmt.Printf("  C: %v;   err: %v\n", tagC, err)

// Output:
// There are 3 tags in the set:
// - A: 42
// - B: true
// - C: foo
//
// Getting typed tags:
//   A: <nil>; err: A: invalid element type
//   C: foo;   err: <nil>
```

## Meta Set

Similarly to tags the `nomix` module provides equivalent interfaces and 
structures for metadata.

```go
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
```

Similar example as for tags with `MetaSet`.

```go
set := NewMetaSet()

set.MetaSet("A", 42)
set.MetaSet("B", true)
set.MetaSet("C", "foo")

fmt.Printf("There are %d entries in the set:\n", set.MetaCount())
fmt.Printf("- A: %v\n", set.MetaGet("A"))
fmt.Printf("- B: %v\n", set.MetaGet("B"))
fmt.Printf("- C: %v\n", set.MetaGet("C"))

fmt.Printf("\nGetting metadata values:\n")

// Tag exists but is of a different type.
metaA, err := set.MetaGetBool("A")
fmt.Printf("  A: %v; err: %v\n", metaA, err)

metaC, err := set.MetaGetString("C")
fmt.Printf("  C: %v;   err: %v\n", metaC, err)

// Output:
// There are 3 entries in the set:
// - A: 42
// - B: true
// - C: foo
//
// Getting metadata values:
//   A: false; err: A: invalid element type
//   C: foo;   err: <nil>
```

