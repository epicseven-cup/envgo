package getorelseenv

import (
	"errors"
	"os"
	"strconv"
)

// GetValueOrDefault retrieves the value of the environment variable specified by key.
// If the variable is not set or is an empty string, it returns the defaultValue.
//
// The generic type T is used to handle general type conversion (e.g., string to int).
// It returns an error if type conversion fails for a non-empty environment variable.
//
// # By default when error, it returns the defaultValue and the error
//
// Example:
//
//	intValue, err := GetValueOrDefault[int]("MY_INT_VAR", 0)
func GetValueOrDefault[T any](key string, defaultValue T) (T, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue, nil
	}

	// Checking what to convert to base on the defaultValue type
	switch any(defaultValue).(type) {
	case string:
		return any(v).(T), nil
	case int:
		n, err := strconv.Atoi(v)
		if err != nil {
			return defaultValue, err
		}
		return any(n).(T), nil
	case float32:
		f32, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return defaultValue, err
		}
		return any(f32).(T), nil
	case float64:
		f64, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return defaultValue, err
		}
		return any(f64).(T), nil
	case bool:
		b, err := strconv.ParseBool(v)
		if err != nil {
			return defaultValue, err
		}
		return any(b).(T), nil
	default:
		return defaultValue, errors.New("No Matching convseration type found")
	}
}
