// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package nomix

import (
	"errors"
	"fmt"
	"slices"
	"time"
)

// GetTag retrieves the [Tag] of type T from the set. Returns the [Tag] if
// found, or the zero value for T and error.
func GetTag[T any](set TagSet, name string) (T, error) {
	var zero T
	if tag := set.TagGet(name); tag != nil {
		if v, ok := tag.(T); ok {
			return v, nil
		}
		return zero, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return zero, fmt.Errorf("%s: %w", name, ErrMissing)
}

// GetTagValue retrieves the value of the [Tag] of type T from the set. Returns
// the value if found, or the zero value for T ane error.
func GetTagValue[T any](set TagSet, name string) (T, error) {
	var zero T
	if tag := set.TagGet(name); tag != nil {
		if v, ok := tag.TagValue().(T); ok {
			return v, nil
		}
		return zero, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return zero, fmt.Errorf("%s: %w", name, ErrMissing)
}

// GetMetaValue retrieves the value of type T from the set. Returns the value
// if found, or the zero value for T and error.
func GetMetaValue[T any](set MetaSet, name string) (T, error) {
	var zero T
	if tag := set.MetaGet(name); tag != nil {
		if v, ok := tag.(T); ok {
			return v, nil
		}
		return zero, fmt.Errorf("%s: %w", name, ErrInvType)
	}
	return zero, fmt.Errorf("%s: %w", name, ErrMissing)
}

// CreateInt64 casts the value to int64. Returns the int64 and nil error if the
// value is a byte, int, int8, int16, int32, or int64. Returns 0 and
// [ErrInvType] if the value is not a supported integer type.
func CreateInt64(val any) (int64, error) {
	switch v := val.(type) {
	case int64:
		return v, nil
	case byte:
		return int64(v), nil
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	}
	return 0, ErrInvType
}

// convertableToInt64 lists types that can be upgraded to int64 without loss of
// precision.
type convertableToInt64 interface {
	byte | int | int8 | int16 | int32 | int64
}

// createInt64Slice upgrades a slice of values that can be upgraded to []int64
// without loss of precision.
func toInt64Slice[T convertableToInt64](v []T) []int64 {
	upgraded := make([]int64, len(v))
	for i, val := range v {
		upgraded[i] = int64(val)
	}
	return upgraded
}

// CreateInt64Slice casts the value to []int64. Returns the []int64 and nil
// error if the value is a []int, []int8, []int16, []int32, or []int64. Returns
// nil and [ErrInvType] if the value's type is not a supported numeric slice
// type.
func CreateInt64Slice(val any) ([]int64, error) {
	switch v := val.(type) {
	case []int64:
		return v, nil
	case []int:
		return toInt64Slice(v), nil
	case []byte:
		return toInt64Slice(v), nil
	case []int8:
		return toInt64Slice(v), nil
	case []int16:
		return toInt64Slice(v), nil
	case []int32:
		return toInt64Slice(v), nil
	}
	return nil, ErrInvType
}

// CreateFloat64 casts the value to float64. Returns the float64 and nil error
// if the value is a byte, int, int8, int16, int32, int64, float32, or float64.
// Returns 0.0 and [ErrInvType] if the value's type is not a supported numeric
// type.
//
// NOTE: For values outside ±2^53 range, the function will return an error.
func CreateFloat64(val any) (float64, error) {
	const maxSafeFloat64 = 1 << 53

	switch v := val.(type) {
	case float64:
		return v, nil
	case int:
		if v > maxSafeFloat64-1 || v < -(maxSafeFloat64-1) {
			msg := "int value out of range for precise float64 conversion"
			return 0, errors.New(msg)
		}
		return float64(v), nil
	case byte:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		if v > maxSafeFloat64-1 || v < -(maxSafeFloat64-1) {
			msg := "int64 value out of range for precise float64 conversion"
			return 0, errors.New(msg)
		}
		return float64(v), nil
	case float32:
		return float64(v), nil
	}
	return 0, ErrInvType
}

// convertableToFloat64 lists types that can be upgraded to float64 without
// loss of precision.
type convertableToFloat64 interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// toFloat64Slice upgrades a slice of values that can be upgraded to []float64
// without loss of precision.
//
// NOTE: For values outside ±2^53 range, the function will return an error.
func toFloat64Slice[T convertableToFloat64](val []T) ([]float64, error) {
	upgraded := make([]float64, len(val))
	for i, v := range val {
		var err error
		if upgraded[i], err = CreateFloat64(v); err != nil {
			return nil, err
		}
	}
	return upgraded, nil
}

// CreateFloat64Slice casts the value to []float64. Returns the slice and nil
// error if the value is a []int, []int8, []int16, []int32, []int64, []float32,
// or []float64. Returns nil and [ErrInvType] if the value's type is not a
// supported numeric slice type.
//
// NOTE: For values outside ±2^53 range, the function will return an error.
func CreateFloat64Slice(val any) ([]float64, error) {
	switch v := val.(type) {
	case []float64:
		return v, nil
	case []int:
		return toFloat64Slice(v)
	case []int8:
		return toFloat64Slice(v)
	case []int16:
		return toFloat64Slice(v)
	case []int32:
		return toFloat64Slice(v)
	case []int64:
		return toFloat64Slice(v)
	case []float32:
		return toFloat64Slice(v)
	}
	return nil, ErrInvType
}

// CreateTime casts the value to [time.Time] or when the value is a string
// parses it in the given location but only when the format is provided.
// Returns the time and nil error on success. Returns zero value time and error
// if the value's type is not a supported type or the value is not a valid time
// representation.
func CreateTime(val any, opts Options) (time.Time, error) {
	switch v := val.(type) {
	case time.Time:
		return v, nil
	case string:
		if slices.Contains(opts.zeroTime, v) {
			return time.Time{}, nil
		}
		return ParseTime(v, opts)
	}
	return time.Time{}, ErrInvType
}

// CreateTimeSlice casts the value to []time.Time, or when the value is a
// []string, it parses its elements in the given location but only when the
// format is provided. With the "zvt" argument you may provide a list of
// strings representing zero value time. Returns the []time.Time and nil error
// on success. Returns nil and error if the value's type is not a supported
// type or the value is not a valid time representation.
func CreateTimeSlice(val any, opts Options) ([]time.Time, error) {
	switch v := val.(type) {
	case []time.Time:
		return v, nil
	case []string:
		if opts.TimeFormat == "" {
			return nil, ErrInvType
		}
		times := make([]time.Time, len(v))
		for i, str := range v {
			if slices.Contains(opts.zeroTime, str) {
				times[i] = time.Time{}
				continue
			}
			var err error
			if times[i], err = ParseTime(str, opts); err != nil {
				return nil, err
			}
		}
		return times, nil
	}
	return nil, ErrInvType
}

// ParseTime parses a string representation of [time.Time]. Returns the time
// and nil error if the value is a valid time representation. Returns zero
// value time and [ErrInvFormat] if the value is not a valid time
// representation.
func ParseTime(val string, opts Options) (time.Time, error) {
	if opts.TimeFormat == "" {
		return time.Time{}, ErrInvType
	}
	if slices.Contains(opts.zeroTime, val) {
		return time.Time{}, nil
	}
	var tim time.Time
	var err error
	if opts.Location != nil {
		tim, err = time.ParseInLocation(opts.TimeFormat, val, opts.Location)
	} else {
		tim, err = time.Parse(opts.TimeFormat, val)
	}
	if err != nil {
		return time.Time{}, ErrInvFormat
	}
	if tim.Location().String() == "" {
		tim = tim.UTC()
	}
	return tim, nil
}

// TagParserNotImpl returns an error indicating that the tag parser is not
// implemented.
func TagParserNotImpl(name, _ string, _ ...Option) (Tag, error) {
	return nil, fmt.Errorf("%s: tag parser %w", name, ErrNotImpl)
}
