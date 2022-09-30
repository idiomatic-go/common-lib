package util

import (
	"os"
	"strings"
)

const ReferencePrefix = "$"

func Getenv(key string) string {
	v := os.Getenv(key)
	if !strings.HasPrefix(v, ReferencePrefix) {
		return v
	}
	return os.Getenv(strings.TrimPrefix(v, ReferencePrefix))
}
