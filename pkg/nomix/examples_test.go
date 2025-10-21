// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix_test

import (
	"database/sql/driver"
	"fmt"
	"strconv"

	"github.com/ctx42/verax/pkg/verax"

	"github.com/ctx42/nomix/pkg/nomix"
	"github.com/ctx42/nomix/pkg/xtag"
)

func ExampleCreateFunc() {
	type Int = nomix.Single[int]

	// CreateInt is a function creating a new [Int] instance by casting
	// argument "val" to int.
	var CreateInt = func(name string, val any, opts ...nomix.Option) (*Int, error) {
		if v, ok := val.(int); ok {
			sqlValue := func(val int) (driver.Value, error) {
				return int64(val), nil
			}
			return nomix.NewSingle(name, v, nomix.KindInt, strconv.Itoa, sqlValue), nil
		}
		return nil, nomix.ErrInvType
	}

	var tcf nomix.CreateFunc
	tcf = nomix.TagCreateFunc(CreateInt)

	tag, err := tcf("A", 42)
	format := "name: %s; kind: %s; value: %v; err: %v\n"
	fmt.Printf(format, tag.TagName(), tag.TagKind(), tag.TagValue(), err)
	// Output:
	// name: A; kind: KindInt; value: 42; err: <nil>
}

func ExampleNewSpec() {
	spec := nomix.NewSpec(
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
}

func ExampleDefine() {
	def := nomix.Define("name", xtag.IntSpec(), verax.Max(42))

	tag, err := def.TagCreate(42)
	fmt.Printf("- success: %s err: %v\n", tag, err)

	tag, err = def.TagCreate(44)
	fmt.Printf("- failure: %v err: %v\n", tag, err)

	// Output:
	// - success: 42 err: <nil>
	// - failure: <nil> err: name: must be no greater than 42
}

func ExampleRegistry() {
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
}

func ExampleTagSet() {
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
}

func ExampleMetaSet() {
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
}
