[![Go Report Card](https://goreportcard.com/badge/github.com/ctx42/nomix)](https://goreportcard.com/report/github.com/ctx42/nomix)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/ctx42/nomix)
[![Tests](https://github.com/ctx42/nomix/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/ctx42/nomix/actions/workflows/go.yml)

# nomix

Typed, named metadata (tags) for Go, restricted to database-friendly types.

<!-- TOC -->
* [nomix](#nomix)
  * [Introduction](#introduction)
  * [Installation](#installation)
  * [Base and Derived Types](#base-and-derived-types)
  * [Tag Kind](#tag-kind)
  * [Tag Interface](#tag-interface)
  * [Typed Tags](#typed-tags)
  * [Tag Spec](#tag-spec)
  * [Tag Definition](#tag-definition)
  * [Registry](#registry)
  * [Tag Sets](#tag-sets)
<!-- TOC -->

In the world of software engineering, we often deal with a vast array of assets
like files, profiles, products, and so on. These assets form the backbone of
application data structures, but as systems grow more complex, simply storing
the core data isn't enough. That's where metadata comes into play. Metadata
provides a way to add extra layers of information and organization to these
assets, making systems more adaptable.

The `nomix` – from *nomos* (law/order) or *nomen* (name) library provides a
simple interface to deal with multi type labeled metadata value sets in a
generic way. In the context of the module we call them **tags**.

Before explaining the library's API, let's introduce its approach and
philosophy.

## Introduction

When we talk about tags, we usually think about a set of name-value pairs
representing a set of data that describes and gives information about other
data.

![tagged_gopher.png](doc/tagged_gopher.png)

Using `map[string]any` in Go is a simple way to model this kind of data, but it
has several drawbacks:

- Managing `any` values is challenging.
- Comparing `any` values is challenging.
- Restricting `any` to specific supported types is not straightforward.
- Mapping `any` to database types is complex.

We need a more sophisticated solution that is user-friendly for programmers,
supports easy definition types and validation rules, and simplifies database
storage. To enable efficient searching and indexing of tags, we must limit
supported types to those widely available across databases.

## Installation

Requires Go 1.26 or later. To use `nomix` in your Go project, install it with:

```bash
go get github.com/ctx42/nomix
```

The core abstractions live in `pkg/nomix` and the ready-to-use tag types in
`pkg/xtag`:

```go
import (
	"github.com/ctx42/nomix/pkg/nomix"
	"github.com/ctx42/nomix/pkg/xtag"
)
```

## Base and Derived Types

The requirement to limit types to a small, widely supported set across most
database engines guided our design. We named these **Base Types** and selected
the following seven during the module's design:

- `byte`
- `string`
- `int64`
- `float64`
- `time`
- `JSON`
- `UUID`

In the database layer, the **Base Types** can map to multiple types, termed
**Derived Types**, in the application layer. Examples of derived types include:

- `bool`
- `int`
- ... and any type you define, such as a `uint16`

## Tag Kind

Each type is assigned a `Kind`, an `int16` number.

![tag-kind.png](doc/tag-kind.png)

The lower octet is reserved by the library to hold *base types*, the seven
higher bits are for use by a library user to create *derived types*, except the
values `0x01` and `0x02` the library's own `KindBool` and `KindInt` use. Notice
that there is also a bit indicating *slice type* which allows us to create
typed slices (arrays) of the base types.

For example, if we wanted to define `[]int` *derived type* we could do it as
follows:

![tag-int.png](doc/tag-int.png)

Based on the above idea, we can seamlessly map types from Go to a database.

![type-mapping.png](doc/type-mapping.png)

## Tag Interface

With the basis of the module explained, we can define the main `nomix`
interface:

```go
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
    // You may use the value returned by the [Tag.TagKind] method
    // to cast it to the proper type.
    TagValue() any
}
```

The `Tag` interface is implemented by all the **Named** and **Typed Tags**. It
still has the `any` returned from the `TagValue` method, but we need it if we
want to be able to handle tag sets with all the supported *base types* and
`derived types` by the same interface.

## Typed Tags

To create typed tags the `nomix` package provides generic `Single` and `Slice`
types which both implement the `Tag` interface and at the same time provide
typed methods to access their typed values.

```go
// Single is a generic type for single value [Tag].
type Single[T comparable] struct { ... }

// Slice is a generic type for multi value [Tag].
type Slice[T comparable] struct { ... }
```

Both of the types have constructor functions (`NewSingle` and `NewSlice`) to
create the instances.

Let's say we would like to define an `int` *derived type* based on `int64` *base
type* and its constructor function.

```go
// Int is a tag representing a single integer value.
type Int = nomix.Single[int]

// NewInt returns a new instance of [Int].
func NewInt(name string, val int) *Int {
    return nomix.NewSingle(name, val, nomix.KindInt, strconv.Itoa, sqlValueInt)
}

// sqlValueInt converts int to its int64 representation. Never returns an error.
func sqlValueInt(val int) (driver.Value, error) { return int64(val), nil }
```

That's it. The `nomix.NewSingle` is the `nomix.Single[T comparable]` constructor
function which you call defining the tag name, value, kind, a function that
returns the string representation of the *derived type* and a function returning
`driver.Value` for it.

You can see the full implementation of `Int` and many other base and derived
types in `xtag` package.

## Tag Spec

The `KindSpec` is a structure providing tag kind specification, it puts
together the `Kind` and the `CreateFunc` and `ParseFunc` functions which create
tags of given *kind* and parse tags of that kind from strings respectively. The
functions have signatures:

```go
// CreateFunc function signature for creating [Tag] instances.
type CreateFunc func(name string, val any, opts ...Option) (Tag, error)

// ParseFunc function signature for creating [Tag] instances from their string
// representation.
type ParseFunc func(name, val string, opts ...Option) (Tag, error)
```

The `TagCreateFunc` and `TagParseFunc` adapters turn functions returning
concrete tag types into a `CreateFunc` and a `ParseFunc`.

The *spec* for the previously seen `Int` tag looks like this:

<!-- gmmce:pkg/nomix/ExampleNewKindSpec -->
```go
spec := nomix.NewKindSpec(
	nomix.KindInt,
	nomix.TagCreateFunc(xtag.CreateInt),
	nomix.TagParseFunc(xtag.ParseInt),
)

tagA, errA := spec.TagCreate("A", 42)
tagB, errB := spec.TagParse("B", "42")

fmt.Printf("- A: %v err: %v\n", tagA.TagValue(), errA)
fmt.Printf("- B: %v err: %v\n", tagB.TagValue(), errB)
// Output:
// - A: 42 err: <nil>
// - B: 42 err: <nil>
```

The `KindSpec` main job is to simplify creation of tags for given *base* or
*derived* types.

## Tag Definition

The `Definition` defines a named tag using a `KindSpec` and an optional set of
[validation rules](https://github.com/ctx42/verax).

<!-- gmmce:pkg/nomix/ExampleDefine -->
```go
def := nomix.Define("name", xtag.IntSpec(), verax.Max(42))

tag, err := def.TagCreate(42)
fmt.Printf("- success: %s err: %v\n", tag, err)

tag, err = def.TagCreate(44)
fmt.Printf("- failure: %v err: %v\n", tag, err)

// Output:
// - success: 42 err: <nil>
// - failure: <nil> err: name: must be less or equal to 42
```

With `Definition` you can create definitions for tags you are using in your
system along with their validation and then use it to create instances of the
tags.

## Registry

The `Registry` allows you to register *specs* (for *kinds*) and then associate
Go types with them.

<!-- gmmce:pkg/nomix/ExampleRegistry -->
```go
reg := nomix.NewRegistry()

_ = reg.Register(xtag.IntSpec())       // Register spec.
_, _ = reg.Associate(0, nomix.KindInt) // Associate the int type with spec.

spec := reg.SpecForKind(nomix.KindInt) // Get spec for KindInt.
tag, err := spec.TagCreate("A", 42)
format := "name: %s; kind: %s; value: %v; err: %v\n"
fmt.Printf(format, tag.TagName(), tag.TagKind(), tag.TagValue(), err)

spec = reg.SpecForType(0) // Get spec for int type.
tag, err = spec.TagCreate("B", 44)
fmt.Printf(format, tag.TagName(), tag.TagKind(), tag.TagValue(), err)

// Convenience function to create tags for registered types.
tag, err = reg.Create("C", 11)
fmt.Printf(format, tag.TagName(), tag.TagKind(), tag.TagValue(), err)

// Output:
// name: A; kind: KindInt; value: 42; err: <nil>
// name: B; kind: KindInt; value: 44; err: <nil>
// name: C; kind: KindInt; value: 11; err: <nil>
```

The `nomix` package also provides a global registry if you wish to register all
your specs and types, for example, in an `init` function.

```go
reg := nomix.GlobalRegistry()
```

The `xtag` package registers all its specs and type associations in one call.
Kinds already registered keep their spec and types already associated keep
their kind, so `RegisterAll` is safe to call more than once on the same
registry:

<!-- gmmce:pkg/xtag/ExampleRegisterAll -->
```go
reg := nomix.NewRegistry()
xtag.RegisterAll(reg) // Safe to call more than once.

tagA, errA := reg.Create("A", 42)
tagB, errB := reg.Create("B", time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC))

format := "%s: kind: %s; value: %v; err: %v\n"
fmt.Printf(format, tagA.TagName(), tagA.TagKind(), tagA, errA)
fmt.Printf(format, tagB.TagName(), tagB.TagKind(), tagB, errB)
// Output:
// A: kind: KindInt; value: 42; err: <nil>
// B: kind: KindTime; value: 2000-01-02T03:04:05Z; err: <nil>
```

## Tag Sets

The `nomix.TagSet` is a structure helping to operate on sets of typed tags.

<!-- gmmce:pkg/nomix/ExampleTagSet -->
```go
set := nomix.NewTagSet()

set.TagSet(
	xtag.NewInt("A", 42),
	xtag.NewBool("B", true),
	xtag.NewString("C", "foo"),
)

fmt.Printf("There are %d tags in the set:\n", set.TagCount())
fmt.Printf("- A: %v\n", set.TagGet("A").TagValue())
fmt.Printf("- B: %v\n", set.TagGet("B").TagValue())
fmt.Printf("- C: %v\n", set.TagGet("C").TagValue())
fmt.Printf("- D: %v\n", set.TagGet("D"))

// Output:
// There are 3 tags in the set:
// - A: 42
// - B: true
// - C: foo
// - D: <nil>
```

The `nomix.MetaSet` does the same for untyped metadata.

<!-- gmmce:pkg/nomix/ExampleMetaSet -->
```go
set := nomix.NewMetaSet()

set.MetaSet("A", 42)
set.MetaSet("B", true)
set.MetaSet("C", "foo")

fmt.Printf("There are %d entries in the set:\n", set.MetaCount())
fmt.Printf("- A: %v\n", set.MetaGet("A"))
fmt.Printf("- B: %v\n", set.MetaGet("B"))
fmt.Printf("- C: %v\n", set.MetaGet("C"))
fmt.Printf("- D: %v\n", set.MetaGet("D"))

// Output:
// There are 3 entries in the set:
// - A: 42
// - B: true
// - C: foo
// - D: <nil>
```

For convenience the `xtag` package provides *specs*, and *typed tags* for most
used types:

- `String` and `StringSlice`
- `Int64` and `Int64Slice`
- `Float64` and `Float64Slice`
- `Time` and `TimeSlice`
- `Bool` and `BoolSlice`
- `Int` and `IntSlice`
- `ByteSlice`
- `JSON`
