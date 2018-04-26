package main

import (
	"github.com/gomodule/redigo/redis"
)

var globalDbPool *redis.Pool

func main() {
	loadConfigGlobals()
	globalDbPool = newDbPool()
}
