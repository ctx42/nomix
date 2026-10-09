// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac.
// SPDX-License-Identifier: MIT

package xtag_test

import (
	"fmt"
	"time"

	"github.com/ctx42/nomix/pkg/nomix"
	"github.com/ctx42/nomix/pkg/xtag"
)

func ExampleRegisterAll() {
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
}
