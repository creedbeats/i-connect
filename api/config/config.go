package config

import (
	"os"
)

// Config func to get env value
func Get(key string) string {
    return os.Getenv(key)
}