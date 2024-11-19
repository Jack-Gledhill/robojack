package env

import "os"

// RequiredVar will error if the provided environment variable wasn't set
func RequiredVar(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(key + " was not provided")
	}

	return val
}

// DefaultVar attempts to get the value of an environment variable, or returns a custom default value if not set
func DefaultVar(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	return val
}
