// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
)

func ExampleTagSet() {
	set := NewTagSet()

	set.TagSet(NewInt("A", 42), NewBool("B", true), NewString("C", "foo"))

	fmt.Printf("There are %d tags in the set:\n", set.TagCount())
	fmt.Printf("- A: %v\n", set.TagGet("A").TagValue())
	fmt.Printf("- B: %v\n", set.TagGet("B").TagValue())
	fmt.Printf("- C: %v\n", set.TagGet("C").TagValue())

	fmt.Printf("\nGetting typed tags:\n")

	// Tag exists but is of a different type.
	tagA, err := set.TagGetBool("A")
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
}

func ExampleMetaSet() {
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
}
