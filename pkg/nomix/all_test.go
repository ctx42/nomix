// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"fmt"
	"strconv"
)

// TstIntSpec returns [Spec] used in testing. It is a very simple spec
// representing an int type.
func TstIntSpec() Spec {
	return Spec{knd: KindInt, tcr: TstIntCreate, tpr: TstIntParse}
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
func (t *TstRule) Validate(val any) error { return t.Err }
