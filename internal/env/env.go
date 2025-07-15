package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	keys, bol := os.LookupEnv(key)

	if !bol {
		return fallback
	}

	return keys
}

func GetInt(key string, fallback int) int {
	keys, bol := os.LookupEnv(key)

	if !bol {
		return fallback
	}

	valInt, err := strconv.Atoi(keys)

	if err != nil {
		return fallback
	}

	return valInt
}
