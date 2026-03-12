package config

import (
	"os"
	"strconv"
)

func Get[T any](key string, fallback T) T {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	switch any(fallback).(type) {
	case string:
		return any(val).(T)
	case int:
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return fallback
		}
		return any(intVal).(T)
	case bool:
		boolVal, err := strconv.ParseBool(val)
		if err != nil {
			return fallback
		}
		return any(boolVal).(T)
	default:
		return fallback
	}
}
