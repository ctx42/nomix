// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac.
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"strconv"

	"github.com/ctx42/testing/pkg/tester"
)

// TstIntSpec returns [KindSpec] used in testing. It is a very simple spec
// representing an int type.
func TstIntSpec() KindSpec {
	return KindSpec{knd: KindInt, tcr: TstIntCreate, tpr: TstIntParse}
}

// TstIntCreate is a test create function matching [CreateFunc] signature.
func TstIntCreate(name string, val any, opts ...Option) (Tag, error) {
	switch v := val.(type) {
	case int:
		return NewSingle(name, v, KindInt, strconv.Itoa, nil), nil
	case string:
		def := NewOptions(opts...)
		if def.Radix == 16 && v == "AA" {
			return NewSingle(name, 170, KindInt, strconv.Itoa, nil), nil
		}
		return nil, ErrInvFormat
	default:
		return nil, ErrInvType
	}
}

// TstIntParse is a test parse function matching [ParseFunc] signature.
func TstIntParse(name, val string, opts ...Option) (Tag, error) {
	def := NewOptions(opts...)
	v, err := strconv.ParseInt(val, def.Radix, 0)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, ErrInvFormat)
	}
	return NewSingle(name, int(v), KindInt, strconv.Itoa, nil), nil
}

// TstRule implements [verax.Rule] interface for use in testing.
type TstRule struct{ Err error }

// Validate returns [TstRule.Err].
func (rul *TstRule) Validate(val any) error { return rul.Err }

// TstTag returns a new [Tag] instance with the given name and value.
func TstTag[T comparable](t tester.T, name string, knd Kind, value T) *TagMock {
	t.Helper()
	tag := NewTagMock(t)
	tag.OnTagName().Return(name).Optional()
	tag.OnTagKind().Return(knd).Optional()
	tag.OnTagValue().Return(value).Optional()
	return tag
}

// tstTagger is a [Tagger] implementation other than [TagSet].
type tstTagger map[string]Tag

func (tgr tstTagger) TagGet(name string) Tag { return tgr[name] }

func (tgr tstTagger) TagSet(tags ...Tag) {
	for _, tag := range tags {
		tgr[tag.TagName()] = tag
	}
}

func (tgr tstTagger) TagDelete(name string) { delete(tgr, name) }

// tstMetadata is a [Metadata] implementation other than [MetaSet].
type tstMetadata map[string]any

func (mtd tstMetadata) MetaGet(name string) any { return mtd[name] }

func (mtd tstMetadata) MetaSet(name string, value any) { mtd[name] = value }

func (mtd tstMetadata) MetaDelete(name string) { delete(mtd, name) }
