package getorelseenv

import "os"

// GetValueOrDefault retrieves a value from the environment variable with the given key.
//
// Args:
//
//	key: The name of the environment variable to retrieve.
//	value: The default value to return if the environment variable is not set.
//
// Returns:
//
//	The value of the environment variable if it is set, otherwise the
func GetValueOrDefault(key string, value string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return value
	}
	return v
}
