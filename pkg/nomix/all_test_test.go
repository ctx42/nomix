// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"github.com/ctx42/testing/pkg/tester"
)

// TstTag returns a new [Tag] instance with the given name and value.
func TstTag[T comparable](t tester.T, name string, knd Kind, value T) *TagMock {
	tag := NewTagMock(t)
	tag.OnTagName().Return(name).Optional()
	tag.OnTagKind().Return(knd).Optional()
	tag.OnTagValue().Return(value).Optional()
	return tag
}
