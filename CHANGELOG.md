## v0.9.0 (Tue, 19 May 2026 20:33:33 UTC)
- refactor(nomix): rename IsVeraxError to IsNomixError.

## v0.8.2 (Tue, 19 May 2026 20:30:07 UTC)
- chore: bump verax to v0.9.0, convert to v0.9.1.

## v0.8.1 (Sat, 09 May 2026 19:52:16 UTC)
- feat(nomix)!: add Errorf constructors, adopt new xrr error API.
- refactor(nomix): adopt NewInternalErrorf at call sites.

## v0.8.0 (Fri, 08 May 2026 15:18:13 UTC)
- feat(nomix): add TagRule, test single-rule Define branch.

## v0.7.0 (Thu, 07 May 2026 12:51:03 UTC)
- refactor(nomix): improve coverage, use verax validation errors.

## v0.6.0 (Mon, 04 May 2026 20:43:23 UTC)
- chore: upgrade Go to 1.126 and all module dependencies.
- chore: Remove the email address from copyright lines.
- feat(nomix): add package-local error domain types.
- refactor(nomix): use package-local error type in Single and Slice.
- refactor(nomix)!: rename Spec to KindSpec, add spec serialization.

## v0.5.0 (Fri, 13 Feb 2026 13:08:14 UTC)
- Change Registry to make it more usable.
- Simplify Registry API by removing redundant methods.
- Change TagKind bit association.
- Add support for `driver.Valuer` interface in `single` and `slice` generic types.
- Expose structures allowing for custom tag implementations.
- style: Use an empty `Options` struct when options are not used.
- Move tag type definitions to separate `xtag` package.
- doc: Improve code documentation.

## v0.4.0 (Fri, 17 Oct 2025 19:56:52 UTC)
- Changes to tag validation logic.

## v0.3.0 (Fri, 17 Oct 2025 14:56:59 UTC)
- feat: Add the ability to define validation rules for a tag `Definition`.

## v0.2.0 (Fri, 17 Oct 2025 10:24:41 UTC)
- style: Rename method arguments in `nomix.MetaSet`.
- feat: Develop helper functions to cast `any` to other wanted types.
- test: Refactor time-related tests.
- test: Rearrange code and tests.
- test: Improve test coverage.
- feat!: Options passed by value instead of a pointer.
- feat: Develop helper functions to cast `any` to other wanted types.
- feat: Develop helper functions to cast `any` to other wanted types.
- feat: Add tag Create* constructor functions.
- feat: Implement `nomix.Creators` which provides a simple interface to create tags for registered types. It also allows you to register your own types on a package level.
- feat!: The `TagCreator` is not an interface.
- feat: Implement `fmt.Stringer` for `TagKind`.

## v0.1.1 (Sat, 11 Oct 2025 18:02:29 UTC)
- chore: Run `go mod tidy`.
- feat!: Move bit indicating a slice one position to the right. Allow int32 to be used for the bitmask.

## v0.1.0 (Fri, 10 Oct 2025 12:37:43 UTC)
- Initial commit.
- Remove unused tag kinds.
- doc: Write module documentation and usage example.

