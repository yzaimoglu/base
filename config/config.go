package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/rs/zerolog"
	utils_config "github.com/yzaimoglu/base/utils/config"
)

const (
	StandardTTL = 3 * time.Hour
)

func Preconfiguration() {
	godotenv.Load()
	if os.Args[1] != "migrate" {
		port := utils_config.GetUInteger(utils_config.EnvPort)
		if port == 0 {
			port = 3000
		}
		os.Args = append(os.Args, fmt.Sprintf("--http=0.0.0.0:%d", port))
	}
}

func MigrateConfiguration(app *pocketbase.PocketBase) {
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isGoRun,
	})
}

func ConfigureLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
