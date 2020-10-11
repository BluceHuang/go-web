package main

import (
	"goweb/routes"
	"goweb/util/config"
	"log"
)

func main() {
	httpPort := config.ServerConfig().HttpPort
	if err := routes.BuildRoute().Run(httpPort); err != nil {
		log.Fatalf("server run failed: %v", err)
	}
}
