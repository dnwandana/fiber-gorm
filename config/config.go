package config

import "os"

// Env function return env value from .env file or OS env value
func Env(key string) string {
	return os.Getenv(key)
}
