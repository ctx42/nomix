// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package xtag provides tags for primitive types.
package xtag

import (
	"encoding/json"
	"time"

	"github.com/ctx42/nomix/pkg/nomix"
)

// RegisterAll registers all [nomix.Spec] the package provides in the
// given [nomix.Registry] and associates types with them.
func RegisterAll(reg *nomix.Registry) {
	mustRegisterKind(reg, int64Spec)
	mustRegisterKind(reg, intSpec)
	mustRegisterKind(reg, float64Spec)
	mustRegisterKind(reg, boolSpec)
	mustRegisterKind(reg, stringSpec)
	mustRegisterKind(reg, timeSpec)
	mustRegisterKind(reg, jsonSpec)
	mustRegisterKind(reg, byteSliceSpec)
	mustRegisterKind(reg, intSliceSpec)
	mustRegisterKind(reg, int64SliceSpec)
	mustRegisterKind(reg, float64SliceSpec)
	mustRegisterKind(reg, boolSliceSpec)
	mustRegisterKind(reg, stringSliceSpec)
	mustRegisterKind(reg, timeSliceSpec)

	mustAssociateType(reg, byte(1), nomix.KindInt64)
	mustAssociateType(reg, int(1), nomix.KindInt)
	mustAssociateType(reg, int8(1), nomix.KindInt64)
	mustAssociateType(reg, int16(1), nomix.KindInt64)
	mustAssociateType(reg, int32(1), nomix.KindInt64)
	mustAssociateType(reg, int64(1), nomix.KindInt64)
	mustAssociateType(reg, float32(1), nomix.KindFloat64)
	mustAssociateType(reg, float64(1), nomix.KindFloat64)
	mustAssociateType(reg, true, nomix.KindBool)
	mustAssociateType(reg, "string", nomix.KindString)
	mustAssociateType(reg, time.Time{}, nomix.KindTime)
	mustAssociateType(reg, json.RawMessage{}, nomix.KindJSON)

	mustAssociateType(reg, []byte{}, nomix.KindByteSlice)
	mustAssociateType(reg, []int{}, nomix.KindIntSlice)
	mustAssociateType(reg, []int8{}, nomix.KindInt64Slice)
	mustAssociateType(reg, []int16{}, nomix.KindInt64Slice)
	mustAssociateType(reg, []int32{}, nomix.KindInt64Slice)
	mustAssociateType(reg, []int64{}, nomix.KindInt64Slice)
	mustAssociateType(reg, []float32{}, nomix.KindFloat64Slice)
	mustAssociateType(reg, []float64{}, nomix.KindFloat64Slice)
	mustAssociateType(reg, []bool{}, nomix.KindBoolSlice)
	mustAssociateType(reg, []string{}, nomix.KindStringSlice)
	mustAssociateType(reg, []time.Time{}, nomix.KindTimeSlice)
}

// mustAssociateType calls [nomix.Registry.Register], and panics on error.
func mustRegisterKind(reg *nomix.Registry, spec nomix.Spec) {
	if err := reg.Register(spec); err != nil {
		panic(err)
	}
}

// mustAssociateType calls [nomix.Registry.Associate], and panics on error.
func mustAssociateType(
	reg *nomix.Registry,
	typ any,
	knd nomix.Kind,
) nomix.Kind {

	was, err := reg.Associate(typ, knd)
	if err != nil {
		panic(err)
	}
	return was
}
