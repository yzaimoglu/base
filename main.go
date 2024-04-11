package main

import (
	"github.com/rs/zerolog/log"
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/config"
)

func main() {
	config.Preconfiguration()
	config.ConfigureLogger()
	log.Info().Msg("Starting server...")

	a := base.NewBase()
	config.ConfigurePocketbase(a)

	if err := a.Pocketbase.Start(); err != nil {
		log.Fatal().Msgf("Could not start pocketbase: %s", err.Error())
	}
}
