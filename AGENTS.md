# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run all tests
go test -v -race ./...

# Run tests for a single package
go test -v -race ./pkg/nomix/...
go test -v -race ./pkg/xtag/...

# Run a single test
go test -v -race -run TestFunctionName ./pkg/nomix/...

# Regenerate mocks (in pkg/nomix/)
go generate ./pkg/nomix/...
```

## Architecture

`nomix` is a two-package Go module for handling typed, named metadata (called **tags**) in a generic way, with database-friendly type constraints.

### Package `pkg/nomix` — core abstractions

**Tag kind system** (`tag.go`): `Kind` is an `int16` bitfield. The lower byte holds base kinds (`KindString`, `KindInt64`, `KindFloat64`, `KindTime`, `KindJSON`, `KindUUID`). The upper byte holds derived kinds (e.g. `KindBool = KindInt64 | 0x0100`). `KindSlice` (`0x80`) is a modifier ORed into any kind to produce its slice variant.

**Core interfaces** (`tag.go`, `meta.go`):
- `Tag` — named, typed value (`TagName`, `TagKind`, `TagValue`).
- `Tagger` — set of `Tag` values, keyed by name.
- `Metadata` — untyped set (`map[string]any`), keyed by name.

**Generic implementations** (`single.go`, `slice.go`):
- `Single[T comparable]` — single-value `Tag`. Holds a `strValuer` (to string) and `sqlValuer` (to `driver.Value`) function.
- `Slice[T comparable]` — multi-value `Tag` for slice types.
- Both implement `driver.Valuer`, `fmt.Stringer`, `Comparer`, and `ValueComparer`.

**Spec / Definition / Registry** (`spec.go`, `definition.go`, `registry.go`):
- `Spec` pairs a `Kind` with a `CreateFunc` and a `ParseFunc`; it is the factory for a tag kind.
- `Definition` wraps a `Spec` with a fixed name and optional `verax.Rule` validators.
- `Registry` maps `Kind → Spec` and `reflect.Type → Spec`. A package-level global registry is available via `GlobalRegistry()`.

**Sets** (`tag_set.go`, `meta_set.go`):
- `TagSet` — `map[string]Tag`-backed, rich typed access; use `GetTag[T]` / `GetTagValue[T]` generics for typed retrieval.
- `MetaSet` — `map[string]any`-backed, with typed getter methods (`MetaGetInt`, `MetaGetString`, `MetaGetTime`, etc.) and `Options`-controlled parsing.

**Options** (`options.go`): `Option` functions configure `MetaSet` and tag parsing. Key options: `WithTimeFormat`, `WithTimeLoc`, `WithZeroTime`, `WithRadixHEX`, `WithLen`, `WithMeta`, `WithTags`.

**Test helpers** (`all_test.go`): `TstIntSpec`, `TstIntCreate`, `TstIntParse`, and `TstRule` are exported test helpers used across internal tests.

**Code generation** (`00_generate.go`, `00_generate_main.go`): runs `mocker` from `github.com/ctx42/testing` to generate `tag_mock_test.go` (mock for the `Tag` interface).

### Package `pkg/xtag` — concrete tag types

Implements ready-to-use `Single` and `Slice` tags for all supported base/derived types and exposes a `Spec` for each. Each type file (e.g. `int.go`, `bool.go`, `time.go`) provides:

- A type alias (e.g. `type Int = nomix.Single[int]`)
- A constructor (e.g. `NewInt(name, val)`)
- `CreateXxx` / `ParseXxx` matching `nomix.CreateFunc` / `nomix.ParseFunc`
- `XxxSpec()` returning the `nomix.Spec` for that kind

`RegisterAll(reg)` registers all specs and type associations at once. Supported types: `String`, `Int`, `Int64`, `Float64`, `Bool`, `Time`, `JSON`, and slice variants (`StringSlice`, `IntSlice`, `Int64Slice`, `Float64Slice`, `BoolSlice`, `TimeSlice`, `ByteSlice`).

### Dependency highlights

| Dependency              | Role                                         |
|-------------------------|----------------------------------------------|
| `github.com/ctx42/verax` | Validation rules used in `Definition`.       |
| `github.com/ctx42/xrr`  | Field-scoped error wrapping (`xrr.NewField`). |
| `github.com/ctx42/testing` | `mocker` for generating interface mocks.   |

### Key invariants

- `Kind` bits: lower byte = base type, bit 7 = slice modifier, upper byte = derived-type discriminator.
- `Single` and `Slice` are never constructed without a `strValuer`; `sqlValuer` may be nil (falls back to returning the value directly).
- `TagSet` and `MetaSet` silently ignore nil values on set operations.
- `Registry.Register` is idempotent-safe per kind (returns error on duplicate); `Associate` overwrites silently.