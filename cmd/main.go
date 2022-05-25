package main

import (
	"github.com/praj31/cli_counter/pkg/cache"
)

func main() {
	rdb := cache.GetRedisClient()

	rdb.ListCounters()
}
