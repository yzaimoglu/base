package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	EnvPort = "PORT"
)

func Preconfiguration() {
	godotenv.Load()
	port := GetInteger(EnvPort)
	if port == 0 {
		port = 3000
	}
	os.Args = append(os.Args, fmt.Sprintf("--http=127.0.0.1:%d", port))
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
