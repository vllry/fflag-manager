package main

import (
	"fmt"
	"os"
)

const (
	defaultListenPort   = "50052"
	defaultRedisAddress = "localhost"
	defaultRedisPort    = "6379"
)

// Define config globals.
var listenPort = ""
var redisAddress = ""
var redisPort = ""

func loadConfigGlobals() {
	listenPort = os.Getenv("FFLAG_MANAGER_PORT")
	if listenPort == "" {
		listenPort = defaultListenPort
	}

	redisAddress = os.Getenv("FFLAG_MASTER_REDIS_ADDRESS")
	if redisAddress == "" {
		redisAddress = defaultRedisAddress
	}

	redisPort = os.Getenv("FFLAG_MASTER_REDIS_PORT")
	if redisPort == "" {
		redisPort = defaultRedisPort
	}

	fmt.Println("Loaded config - Redis on ", redisAddress, ":", redisPort)
}
