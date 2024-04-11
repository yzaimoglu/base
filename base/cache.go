package base

import (
	"github.com/maypok86/otter"
	"github.com/rs/zerolog/log"
)

type Cache struct {
	String *otter.CacheWithVariableTTL[string, string]
}

func NewCache(totalSize int) *Cache {
	cacheSize := 1
	singleSize := totalSize / cacheSize

	log.Info().Msg("Creating string cache...")
	stringCache, err := otter.MustBuilder[string, string](singleSize).CollectStats().WithVariableTTL().Build()
	if err != nil {
		log.Error().Msgf("Cache could not be initialized: %s", err.Error())
	}
	log.Info().Msg("Created string cache.")

	return &Cache{
		String: &stringCache,
	}
}
