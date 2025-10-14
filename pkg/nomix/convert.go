package nomix

import (
	"time"
)

// asString casts the value to a string. Returns the string and nil error if
// the value is a string. Returns "" and [ErrInvType] if the value is not a
// string.
func asString(val any, _ *Options) (string, error) {
	switch v := val.(type) {
	case string:
		return v, nil
	default:
		return "", ErrInvType
	}
}

// asStringSlice casts the value to []string. Returns the slice and nil error
// if the value is a []string. Returns nil and [ErrInvType] if not a []string.
func asStringSlice(val any, _ *Options) ([]string, error) {
	switch v := val.(type) {
	case []string:
		return v, nil
	default:
		return nil, ErrInvType
	}
}

// asInt64 casts the value to int64. Returns the int64 and nil error if the
// value is a byte, int, int8, int16, int32, or int64. Returns 0 and
// [ErrInvType] if the value is not a supported integer type.
func asInt64(val any, _ *Options) (int64, error) {
	switch v := val.(type) {
	case int:
		return int64(v), nil
	case byte:
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
	return 0, ErrInvType
}

// convertableToInt64 lists types that can be upgraded to int64 without loss of
// precision.
type convertableToInt64 interface {
	int | int8 | int16 | int32 | int64
}

// asInt64Slice upgrades a slice of values that can be upgraded to []int64
// without loss of precision.
func toInt64Slice[T convertableToInt64](v []T, _ *Options) []int64 {
	upgraded := make([]int64, len(v))
	for i, val := range v {
		upgraded[i] = int64(val)
	}
	return upgraded
}

// asInt64Slice casts the value to []int64. Returns the slice and nil error if
// the value is a []int, []int8, []int16, []int32, or []int64. Returns 0 and
// [ErrInvType] if the value is not a supported integer type.
func asInt64Slice(val any, opts *Options) ([]int64, error) {
	switch v := val.(type) {
	case []int:
		return toInt64Slice(v, opts), nil
	case []int8:
		return toInt64Slice(v, opts), nil
	case []int16:
		return toInt64Slice(v, opts), nil
	case []int32:
		return toInt64Slice(v, opts), nil
	case []int64:
		return v, nil
	}
	return nil, ErrInvType
}

// asFloat64 casts the value to float64. Returns the float64 and nil error if
// the value is a byte, int, int8, int16, int32, int64, float32, or float64.
// Returns 0.0 and [ErrInvType] if the value is not a supported numeric type.
//
// NOTE: For int64 values outside ±2^53 range, the result is undefined.
func asFloat64(val any, _ *Options) (float64, error) {
	switch v := val.(type) {
	case int:
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
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	}
	return 0, ErrInvType
}

// convertableToFloat64 lists types that can be upgraded to float64 without
// loss of precision.
type convertableToFloat64 interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// toFloat64Slice upgrades a slice of values that can be upgraded to
// []float64 without loss of precision.
//
// NOTE: For int64 values outside ±2^53 range, the result is undefined.
func toFloat64Slice[T convertableToFloat64](v []T, _ *Options) []float64 {
	upgraded := make([]float64, len(v))
	for i, val := range v {
		upgraded[i] = float64(val)
	}
	return upgraded
}

// asFloat64Slice casts the value to []float64. Returns the slice and nil error
// if the value is a []int, []int8, []int16, []int32, []int64, []float32, or
// []float64. Returns 0.0 and [ErrInvType] if the value is not a supported
// numeric type.
func asFloat64Slice(val any, opts *Options) ([]float64, error) {
	switch v := val.(type) {
	case []int:
		return toFloat64Slice(v, opts), nil
	case []int8:
		return toFloat64Slice(v, opts), nil
	case []int16:
		return toFloat64Slice(v, opts), nil
	case []int32:
		return toFloat64Slice(v, opts), nil
	case []int64:
		return toFloat64Slice(v, opts), nil
	case []float32:
		return toFloat64Slice(v, opts), nil
	case []float64:
		return v, nil
	}
	return nil, ErrInvType
}

// asTime casts the value to [time.Time] or when the value is a string parses
// it time but only when [Options.timeFormat] is set. Returns the time and nil
// error on success. Returns zero value time and [ErrInvType] if the value is
// not a supported type.
func asTime(val any, opts *Options) (time.Time, error) {
	switch v := val.(type) {
	case time.Time:
		return v, nil
	case string:
		if opts != nil && opts.timeFormat != "" {
			return parseTime(v, opts)
		}
	}
	return time.Time{}, ErrInvType
}

// asTimeSlice casts the value to []time.Time or when the value is a []string
// parses its values as time but only when [Options.timeFormat] is set. Returns
// the slice and nil error on success. Returns nil and [ErrInvType] if the
// value is not a supported type.
func asTimeSlice(val any, opts *Options) ([]time.Time, error) {
	switch v := val.(type) {
	case []time.Time:
		return v, nil
	case []string:
		if opts != nil && opts.timeFormat != "" {
			times := make([]time.Time, len(v))
			for i, str := range v {
				var err error
				if times[i], err = parseTime(str, opts); err != nil {
					return nil, err
				}
			}
			return times, nil
		}
	}
	return nil, ErrInvType
}
