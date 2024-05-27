package main

import (
	"github.com/rs/zerolog/log"
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/config"
	_ "github.com/yzaimoglu/base/migrations"
)

func main() {
	config.ConfigureLogger()
	a := base.NewBase()
	config.MigrateConfiguration(a.Pocketbase)

	log.Info().Msg("Starting the server...")
	config.Preconfiguration()
	config.ConfigurePocketbase(a)
	if err := a.Pocketbase.Start(); err != nil {
		log.Fatal().Msgf("Could not start pocketbase: %s", err.Error())
	}
}
