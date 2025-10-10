// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"encoding/json"
	"fmt"
	"time"
)

var _ Metadata = MetaSet{} // Compile time check.

// MetaSet represents a set of metadata key-values.
type MetaSet struct {
	m map[string]any
}

// NewMetaSet returns a new [MetaSet] instance. By default, the new map is
// initialized with the length equal to 10.
func NewMetaSet(opts ...func(*Options)) MetaSet {
	def := &Options{length: 10}
	for _, opt := range opts {
		opt(def)
	}
	set := MetaSet{}
	if m, ok := def.init.(map[string]any); ok {
		for k, v := range m {
			if v == nil {
				delete(m, k)
			}
		}
		set.m = m
	}
	if set.m == nil {
		set.m = make(map[string]any, def.length)
	}
	return set
}

func (set MetaSet) MetaGet(key string) any {
	return set.m[key]
}

func (set MetaSet) MetaSet(key string, value any) {
	if value == nil {
		return
	}
	set.m[key] = value
}

func (set MetaSet) MetaDelete(key string) { delete(set.m, key) }

// MetaCount returns the number of entries in the metadata set.
func (set MetaSet) MetaCount() int { return len(set.m) }

func (set MetaSet) MetaGetAll() map[string]any { return set.m }

// TODO(rz): Implement to MetaReset.

// MetaGetString gets a metadata value by name as a string. Returns the value
// and nil error if it exists and is the string type. Returns empty string and
// [ErrMissing] if the value is missing, or empty string and [ErrInvType] if
// the value is of a different type.
func (set MetaSet) MetaGetString(key string) (string, error) {
	if valAny, ok := set.m[key]; ok {
		if val, ok := valAny.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("%s: %w", key, ErrInvType)
	}
	return "", fmt.Errorf("%s: %w", key, ErrMissing)
}

// MetaGetStringSlice gets a metadata value by name as a []string. Returns the
// slice and nil error if it exists and is the []string type. Returns nil and
// [ErrMissing] if the value is missing, or nil and [ErrInvType] if the value
// is of a different type.
func (set MetaSet) MetaGetStringSlice(key string) ([]string, error) {
	if valAny, ok := set.m[key]; ok {
		if val, ok := valAny.([]string); ok {
			return val, nil
		}
		return nil, fmt.Errorf("%s: %w", key, ErrInvType)
	}
	return nil, fmt.Errorf("%s: %w", key, ErrMissing)
}

// MetaGetInt64 gets a metadata value by name as an int64. Returns the value
// and nil error if it exists and is one of int, int8, int16, int32 or int64
// types. Returns 0 and [ErrMissing] if the value is missing, or 0 and
// [ErrInvType] if the value is of a different type.
func (set MetaSet) MetaGetInt64(key string) (int64, error) {
	if val, ok := set.m[key]; ok {
		switch v := val.(type) {
		case int:
			return int64(v), nil
		case int8:
			return int64(v), nil
		case int16:
			return int64(v), nil
		case int32:
			return int64(v), nil
		case int64:
			return v, nil
		}
		return 0, fmt.Errorf("%s: %w", key, ErrInvType)
	}
	return 0, fmt.Errorf("%s: %w", key, ErrMissing)
}

// upgradableToInt64 lists types that can be upgraded to int64 without loss of
// precision.
type upgradableToInt64 interface {
	int | int8 | int16 | int32 | int64
}

// upgradeToInt64Slice upgrades a slice of values that can be upgraded to []int64
// without loss of precision.
func upgradeToInt64Slice[T upgradableToInt64](v []T) []int64 {
	upgraded := make([]int64, len(v))
	for i, val := range v {
		upgraded[i] = int64(val)
	}
	return upgraded
}

// MetaGetInt64Slice gets a metadata value by name as a []int64. Returns the
// slice and nil error if it exists and is one of []int, []int8, []int16,
// []int32 or []int64 types. Returns nil and [ErrMissing] if the value is
// missing, or nil and [ErrInvType] if the value is of a different type.
func (set MetaSet) MetaGetInt64Slice(key string) ([]int64, error) {
	if val, ok := set.m[key]; ok {
		switch v := val.(type) {
		case []int:
			return upgradeToInt64Slice(v), nil
		case []int8:
			return upgradeToInt64Slice(v), nil
		case []int16:
			return upgradeToInt64Slice(v), nil
		case []int32:
			return upgradeToInt64Slice(v), nil
		case []int64:
			return v, nil
		}
		return nil, fmt.Errorf("%s: %w", key, ErrInvType)
	}
	return nil, fmt.Errorf("%s: %w", key, ErrMissing)
}

// MetaGetFloat64 gets a metadata value by name as a float64. Returns the value
// and nil error if it exists and is one of int, int8, int16, int32, int64,
// float32, float64 types. Returns 0.0 and [ErrMissing] if the value is missing,
// or 0.0 and [ErrInvType] if the value is of a different type.
//
// NOTE: For int64 values outside ±2^53 range, the result is undefined.
func (set MetaSet) MetaGetFloat64(key string) (float64, error) {
	if val, ok := set.m[key]; ok {
		switch v := val.(type) {
		case int:
			return float64(v), nil
		case int8:
			return float64(v), nil
		case int16:
			return float64(v), nil
		case int32:
			return float64(v), nil
		case int64:
			return float64(v), nil
		case float32:
			return float64(v), nil
		case float64:
			return v, nil
		}
		return 0, fmt.Errorf("%s: %w", key, ErrInvType)
	}
	return 0, fmt.Errorf("%s: %w", key, ErrMissing)
}

// upgradableToFloat64 lists types that can be upgraded to float64 without loss
// of precision.
type upgradableToFloat64 interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// upgradeToFloat64Slice upgrades a slice of values that can be upgraded to
// []float64 without loss of precision.
//
// NOTE: For int64 values outside ±2^53 range, the result is undefined.
func upgradeToFloat64Slice[T upgradableToFloat64](v []T) []float64 {
	upgraded := make([]float64, len(v))
	for i, val := range v {
		upgraded[i] = float64(val)
	}
	return upgraded
}

// MetaGetFloat64Slice gets a metadata value by name as a []float. Returns the
// slice and nil error if it exists and is one of []int, []int8, []int16,
// []int32, []int64, []float32 or []float64 types. Returns nil and [ErrMissing]
// if the value is missing, or nil and [ErrInvType] if the value is of a
// different type.
//
// NOTE: For int64 values outside ±2^53 range, the result is undefined.
func (set MetaSet) MetaGetFloat64Slice(key string) ([]float64, error) {
	if val, ok := set.m[key]; ok {
		switch v := val.(type) {
		case []int:
			return upgradeToFloat64Slice(v), nil
		case []int8:
			return upgradeToFloat64Slice(v), nil
		case []int16:
			return upgradeToFloat64Slice(v), nil
		case []int32:
			return upgradeToFloat64Slice(v), nil
		case []int64:
			return upgradeToFloat64Slice(v), nil
		case []float32:
			return upgradeToFloat64Slice(v), nil
		case []float64:
			return v, nil
		}
		return nil, fmt.Errorf("%s: %w", key, ErrInvType)
	}
	return nil, fmt.Errorf("%s: %w", key, ErrMissing)
}

// MetaGetTime gets a metadata value by name as a [time.Time]. Returns the
// value and nil error if it exists and is one of [time.Time] or string time
// representation. You may customize string the parsing by using the
// [WithTimeFormat] and [WithTimeLoc] options. Returns zero value time and
// [ErrMissing] if the value is missing, or zero vale time and [ErrInvType] if
// the value is of a different type.
//
// To support string zero time values, use the [WithZeroTime] option.
func (set MetaSet) MetaGetTime(key string, opts ...Option) (time.Time, error) {
	def := DefaultOptions()
	def.timeFormat = "" // By default, tring type is not supported.
	for _, opt := range opts {
		opt(def)
	}
	if val, ok := set.m[key]; ok {
		switch v := val.(type) {
		case time.Time:
			return v, nil
		case string:
			if def.timeFormat != "" {
				return parseTime(key, v, def)
			}
		}
		return time.Time{}, fmt.Errorf("%s: %w", key, ErrInvType)
	}
	return time.Time{}, fmt.Errorf("%s: %w", key, ErrMissing)
}

// MetaGetTimeSlice gets a metadata value by name as a slice of [time.Time]
// values. Returns the value and nil error if it exists and is a slice of
// [time.Time] or slice of string time representations. You may customize
// string the parsing by using the [WithTimeFormat] and [WithTimeLoc] options.
// Returns nil and [ErrMissing] if the value is missing, zero vale time and
// [ErrInvType] if the value is of a different type or, zero value time and
// [ErrInvFormat] if the value is not a valid time string.
//
// To support string zero time values, use the [WithZeroTime] option.
func (set MetaSet) MetaGetTimeSlice(key string, opts ...Option) ([]time.Time, error) {
	def := DefaultOptions()
	def.timeFormat = "" // By default, tring type is not supported.
	for _, opt := range opts {
		opt(def)
	}
	if val, ok := set.m[key]; ok {
		switch v := val.(type) {
		case []time.Time:
			return v, nil
		case []string:
			times := make([]time.Time, len(v))
			for i, vv := range v {
				var err error
				times[i], err = parseTime(key, vv, def)
				if err != nil {
					return nil, err
				}
			}
			return times, nil
		}
		return nil, fmt.Errorf("%s: %w", key, ErrInvType)
	}
	return nil, fmt.Errorf("%s: %w", key, ErrMissing)
}

// MetaGetJOSN gets a metadata value by name as a [json.RawMessage]. Returns
// the value and nil error if it exists and is one of [json.RawMessage], string,
// []byte types. Returns nil and [ErrMissing] if the value is missing, or nil
// and [ErrInvType] if the value is of a different type.
func (set MetaSet) MetaGetJOSN(key string) (json.RawMessage, error) {
	var have []byte
	if val, ok := set.m[key]; ok {
		switch v := val.(type) {
		case []byte:
			have = v
		case json.RawMessage:
			have = v
		case string:
			have = []byte(v)
		default:
			return have, fmt.Errorf("%s: %w", key, ErrInvType)
		}
		if !json.Valid(have) {
			return nil, fmt.Errorf("%s: %w", key, ErrInvFormat)
		}
		return have, nil
	}
	return nil, fmt.Errorf("%s: %w", key, ErrMissing)
}

// MetaGetBool gets a metadata value by name as a bool. Returns the value and
// nil error if it exists and is the bool type. Returns false and [ErrMissing]
// if the value is missing, or empty string and [ErrInvType] if the value is of
// a different type.
func (set MetaSet) MetaGetBool(key string) (bool, error) {
	if val, ok := set.m[key]; ok {
		if v, ok := val.(bool); ok {
			return v, nil
		}
		return false, fmt.Errorf("%s: %w", key, ErrInvType)
	}
	return false, fmt.Errorf("%s: %w", key, ErrMissing)
}

// MetaGetBoolSlice gets a metadata value by name as a []bool. Returns the
// slice and nil error if it exists and is the []bool type. Returns nil and
// [ErrMissing] if the value is missing, or nil and [ErrInvType] if the value
// is of a different type.
func (set MetaSet) MetaGetBoolSlice(key string) ([]bool, error) {
	if valAny, ok := set.m[key]; ok {
		if val, ok := valAny.([]bool); ok {
			return val, nil
		}
		return nil, fmt.Errorf("%s: %w", key, ErrInvType)
	}
	return nil, fmt.Errorf("%s: %w", key, ErrMissing)
}

// MetaGetInt gets a metadata value by name as an int. Returns the value and
// nil error if it exists and is the int type. Returns false and [ErrMissing] if
// the value is missing, or empty string and [ErrInvType] if the value is of a
// different type.
func (set MetaSet) MetaGetInt(key string) (int, error) {
	if val, ok := set.m[key]; ok {
		if v, ok := val.(int); ok {
			return v, nil
		}
		return 0, fmt.Errorf("%s: %w", key, ErrInvType)
	}
	return 0, fmt.Errorf("%s: %w", key, ErrMissing)
}

// MetaGetIntSlice gets a metadata value by name as a []int. Returns the
// slice and nil error if it exists and is the []int type. Returns nil and
// [ErrMissing] if the value is missing, or nil and [ErrInvType] if the value
// is of a different type.
func (set MetaSet) MetaGetIntSlice(key string) ([]int, error) {
	if valAny, ok := set.m[key]; ok {
		if val, ok := valAny.([]int); ok {
			return val, nil
		}
		return nil, fmt.Errorf("%s: %w", key, ErrInvType)
	}
	return nil, fmt.Errorf("%s: %w", key, ErrMissing)
}
