package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	utils_config "github.com/yzaimoglu/base/utils/config"
)

const (
	StandardTTL = 3 * time.Hour
)

func Preconfiguration() {
	godotenv.Load()
	port := utils_config.GetUInteger(utils_config.EnvPort)
	if port == 0 {
		port = 3000
	}
	os.Args = append(os.Args, fmt.Sprintf("--http=0.0.0.0:%d", port))
}

func ConfigureLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
