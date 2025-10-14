package nomix

// asString converts value v to its string representation for supported types.
func asString(v any) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	default:
		return "", ErrInvType
	}
}

// convertableToInt64 lists types that can be upgraded to int64 without loss of
// precision.
type convertableToInt64 interface {
	int | int8 | int16 | int32 | int64
}

// asInt64Slice upgrades a slice of values that can be upgraded to []int64
// without loss of precision.
func asInt64Slice[T convertableToInt64](v []T) []int64 {
	upgraded := make([]int64, len(v))
	for i, val := range v {
		upgraded[i] = int64(val)
	}
	return upgraded
}

// convertableToFloat64 lists types that can be upgraded to float64 without
// loss of precision.
type convertableToFloat64 interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// asFloat64Slice upgrades a slice of values that can be upgraded to
// []float64 without loss of precision.
//
// NOTE: For int64 values outside ±2^53 range, the result is undefined.
func asFloat64Slice[T convertableToFloat64](v []T) []float64 {
	upgraded := make([]float64, len(v))
	for i, val := range v {
		upgraded[i] = float64(val)
	}
	return upgraded
}
