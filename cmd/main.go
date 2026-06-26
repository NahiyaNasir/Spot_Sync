package main

import (
	"spot_sync/internal/config"
	"spot_sync/internal/server.go"
)

func main() {
	// load environment variables
	cfg := config.LoadEnv()

	// connect to the database
	db := config.ConnectDatabase(cfg)
	// start the server
	server.Start(db, cfg)

}