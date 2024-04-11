package base

import (
	"github.com/pocketbase/pocketbase"
)

type Base struct {
	Pocketbase *pocketbase.PocketBase
	Caches     *Cache
}

func NewBase() *Base {
	return &Base{
		Pocketbase: pocketbase.NewWithConfig(pocketbase.Config{
			HideStartBanner: false,
		}),
		Caches: NewCache(10_000),
	}
}
