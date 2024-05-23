package config_utils

import (
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

const (
	EnvPort  = "PORT"
	EnvName  = "NAME"
	EnvDebug = "DEBUG"
)

func IsDebug() bool {
	return GetBoolean(EnvDebug)
}

func Port() uint16 {
	return uint16(GetUInteger(EnvPort))
}

func GetString(key string) string {
	return os.Getenv(key)
}

func GetBoolean(key string) bool {
	s := os.Getenv(key)
	b, err := strconv.ParseBool(s)
	if err != nil {
		log.Info().Msgf("Could not parse bool, returning false: %s", err.Error())
		return false
	}
	return b
}

func GetInteger(key string) int64 {
	s := os.Getenv(key)
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Info().Msgf("Could not parse int, returning 0: %s", err.Error())
		return 0
	}
	return i
}

func GetUInteger(key string) uint64 {
	s := os.Getenv(key)
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Info().Msgf("Could not parse uint, returning 0: %s", err.Error())
		return 0
	}
	return i
}
