package main

import (
	"github.com/rs/zerolog/log"
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/config"
	_ "github.com/yzaimoglu/base/migrations"
)

func main() {
	a := base.NewBase()
	config.MigrateConfiguration(a.Pocketbase)
	config.Preconfiguration()
	config.ConfigureLogger()

	log.Info().Msg("Starting the server...")
	config.ConfigurePocketbase(a)
	if err := a.Pocketbase.Start(); err != nil {
		log.Fatal().Msgf("Could not start pocketbase: %s", err.Error())
	}
}
