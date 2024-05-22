package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	EnvPort  = "PORT"
	EnvDebug = "DEBUG"

	StandardTTL = 3 * time.Hour
)

func Preconfiguration() {
	godotenv.Load()
	port := GetUInteger(EnvPort)
	if port == 0 {
		port = 3000
	}
	os.Args = append(os.Args, fmt.Sprintf("--http=0.0.0.0:%d", port))
}

func IsDebug() bool {
	return GetBoolean(EnvDebug)
}

func ConfigureLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
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
