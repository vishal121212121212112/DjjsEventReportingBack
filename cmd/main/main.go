package main

import (
	"log"
	"os"

	"event-reporting/app/api"
	"event-reporting/app/helpers/registry"
)

func main() {
	routes := registry.Build()
	r := api.NewRouter(*routes)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Printf("listening on %s ...", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
