package common

import (
	"os"
	"time"
)

// Server timeouts
const (
	ServerReadTimeout  = 5 * time.Second
	ServerWriteTimeout = 10 * time.Second
	ServerIdleTimeout  = 120 * time.Second
)

// GetEnvWithDefault gets an environment variable or returns a default value
func GetEnvWithDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
