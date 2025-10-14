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

// MetaDeleteAll deletes all metadata from the set.
func (set MetaSet) MetaDeleteAll() {
	for name := range set.m {
		delete(set.m, name)
	}
}

// MetaGetString gets a metadata value by name as a string. Returns the value
// and nil error if it exists and is the string type. Returns empty string and
// [ErrMissing] if the value is missing, or empty string and [ErrInvType] if
// the value is of a different type.
func (set MetaSet) MetaGetString(name string) (string, error) {
	if val, ok := set.m[name]; ok {
		v, err := asString(val)
		if err != nil {
			return "", fmt.Errorf("%s: %w", name, ErrInvType)
		}
		return v, nil
	}
	return "", fmt.Errorf("%s: %w", name, ErrMissing)
}

// MetaGetStringSlice gets a metadata value by name as a []string. Returns the
// slice and nil error if it exists and is the []string type. Returns nil and
// [ErrMissing] if the value is missing, or nil and [ErrInvType] if the value
// is of a different type.
func (set MetaSet) MetaGetStringSlice(name string) ([]string, error) {
	if valAny, ok := set.m[name]; ok {
		if val, ok := valAny.([]string); ok {
			return val, nil
		}
		return nil, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return nil, fmt.Errorf("%s: %w", name, ErrMissing)
}

// MetaGetInt64 gets a metadata value by name as an int64. Returns the value
// and nil error if it exists and is one of int, int8, int16, int32 or int64
// types. Returns 0 and [ErrMissing] if the value is missing, or 0 and
// [ErrInvType] if the value is of a different type.
func (set MetaSet) MetaGetInt64(name string) (int64, error) {
	if val, ok := set.m[name]; ok {
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
		return 0, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return 0, fmt.Errorf("%s: %w", name, ErrMissing)
}

// MetaGetInt64Slice gets a metadata value by name as a []int64. Returns the
// slice and nil error if it exists and is one of []int, []int8, []int16,
// []int32 or []int64 types. Returns nil and [ErrMissing] if the value is
// missing, or nil and [ErrInvType] if the value is of a different type.
func (set MetaSet) MetaGetInt64Slice(name string) ([]int64, error) {
	if val, ok := set.m[name]; ok {
		switch v := val.(type) {
		case []int:
			return asInt64Slice(v), nil
		case []int8:
			return asInt64Slice(v), nil
		case []int16:
			return asInt64Slice(v), nil
		case []int32:
			return asInt64Slice(v), nil
		case []int64:
			return v, nil
		}
		return nil, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return nil, fmt.Errorf("%s: %w", name, ErrMissing)
}

// MetaGetFloat64 gets a metadata value by name as a float64. Returns the value
// and nil error if it exists and is one of int, int8, int16, int32, int64,
// float32, float64 types. Returns 0.0 and [ErrMissing] if the value is missing,
// or 0.0 and [ErrInvType] if the value is of a different type.
//
// NOTE: For int64 values outside ±2^53 range, the result is undefined.
func (set MetaSet) MetaGetFloat64(name string) (float64, error) {
	if val, ok := set.m[name]; ok {
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
		return 0, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return 0, fmt.Errorf("%s: %w", name, ErrMissing)
}

// MetaGetFloat64Slice gets a metadata value by name as a []float. Returns the
// slice and nil error if it exists and is one of []int, []int8, []int16,
// []int32, []int64, []float32 or []float64 types. Returns nil and [ErrMissing]
// if the value is missing, or nil and [ErrInvType] if the value is of a
// different type.
//
// NOTE: For int64 values outside ±2^53 range, the result is undefined.
func (set MetaSet) MetaGetFloat64Slice(name string) ([]float64, error) {
	if val, ok := set.m[name]; ok {
		switch v := val.(type) {
		case []int:
			return asFloat64Slice(v), nil
		case []int8:
			return asFloat64Slice(v), nil
		case []int16:
			return asFloat64Slice(v), nil
		case []int32:
			return asFloat64Slice(v), nil
		case []int64:
			return asFloat64Slice(v), nil
		case []float32:
			return asFloat64Slice(v), nil
		case []float64:
			return v, nil
		}
		return nil, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return nil, fmt.Errorf("%s: %w", name, ErrMissing)
}

// MetaGetTime gets a metadata value by name as a [time.Time]. Returns the
// value and nil error if it exists and is one of [time.Time] or string time
// representation. You may customize string the parsing by using the
// [WithTimeFormat] and [WithTimeLoc] options. Returns zero value time and
// [ErrMissing] if the value is missing, or zero vale time and [ErrInvType] if
// the value is of a different type.
//
// To support string zero time values, use the [WithZeroTime] option.
func (set MetaSet) MetaGetTime(name string, opts ...Option) (time.Time, error) {
	def := DefaultOptions()
	def.timeFormat = "" // By default, tring type is not supported.
	for _, opt := range opts {
		opt(def)
	}
	if val, ok := set.m[name]; ok {
		switch v := val.(type) {
		case time.Time:
			return v, nil
		case string:
			if def.timeFormat != "" {
				return parseTime(name, v, def)
			}
		}
		return time.Time{}, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return time.Time{}, fmt.Errorf("%s: %w", name, ErrMissing)
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
func (set MetaSet) MetaGetTimeSlice(name string, opts ...Option) ([]time.Time, error) {
	def := DefaultOptions()
	def.timeFormat = "" // By default, tring type is not supported.
	for _, opt := range opts {
		opt(def)
	}
	if val, ok := set.m[name]; ok {
		switch v := val.(type) {
		case []time.Time:
			return v, nil
		case []string:
			times := make([]time.Time, len(v))
			for i, vv := range v {
				var err error
				times[i], err = parseTime(name, vv, def)
				if err != nil {
					return nil, err
				}
			}
			return times, nil
		}
		return nil, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return nil, fmt.Errorf("%s: %w", name, ErrMissing)
}

// MetaGetJOSN gets a metadata value by name as a [json.RawMessage]. Returns
// the value and nil error if it exists and is one of [json.RawMessage], string,
// []byte types. Returns nil and [ErrMissing] if the value is missing, or nil
// and [ErrInvType] if the value is of a different type.
func (set MetaSet) MetaGetJOSN(name string) (json.RawMessage, error) {
	var have []byte
	if val, ok := set.m[name]; ok {
		switch v := val.(type) {
		case []byte:
			have = v
		case json.RawMessage:
			have = v
		case string:
			have = []byte(v)
		default:
			return have, fmt.Errorf("%s: %w", name, ErrInvType)
		}
		if !json.Valid(have) {
			return nil, fmt.Errorf("%s: %w", name, ErrInvFormat)
		}
		return have, nil
	}
	return nil, fmt.Errorf("%s: %w", name, ErrMissing)
}

// MetaGetBool gets a metadata value by name as a bool. Returns the value and
// nil error if it exists and is the bool type. Returns false and [ErrMissing]
// if the value is missing, or empty string and [ErrInvType] if the value is of
// a different type.
func (set MetaSet) MetaGetBool(name string) (bool, error) {
	if val, ok := set.m[name]; ok {
		if v, ok := val.(bool); ok {
			return v, nil
		}
		return false, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return false, fmt.Errorf("%s: %w", name, ErrMissing)
}

// MetaGetBoolSlice gets a metadata value by name as a []bool. Returns the
// slice and nil error if it exists and is the []bool type. Returns nil and
// [ErrMissing] if the value is missing, or nil and [ErrInvType] if the value
// is of a different type.
func (set MetaSet) MetaGetBoolSlice(name string) ([]bool, error) {
	if valAny, ok := set.m[name]; ok {
		if val, ok := valAny.([]bool); ok {
			return val, nil
		}
		return nil, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return nil, fmt.Errorf("%s: %w", name, ErrMissing)
}

// MetaGetInt gets a metadata value by name as an int. Returns the value and
// nil error if it exists and is the int type. Returns false and [ErrMissing] if
// the value is missing, or empty string and [ErrInvType] if the value is of a
// different type.
func (set MetaSet) MetaGetInt(name string) (int, error) {
	if val, ok := set.m[name]; ok {
		if v, ok := val.(int); ok {
			return v, nil
		}
		return 0, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return 0, fmt.Errorf("%s: %w", name, ErrMissing)
}

// MetaGetIntSlice gets a metadata value by name as a []int. Returns the
// slice and nil error if it exists and is the []int type. Returns nil and
// [ErrMissing] if the value is missing, or nil and [ErrInvType] if the value
// is of a different type.
func (set MetaSet) MetaGetIntSlice(name string) ([]int, error) {
	if valAny, ok := set.m[name]; ok {
		if val, ok := valAny.([]int); ok {
			return val, nil
		}
		return nil, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return nil, fmt.Errorf("%s: %w", name, ErrMissing)
}
