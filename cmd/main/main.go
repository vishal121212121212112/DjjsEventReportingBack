package main

import (
	"log"

	"event-reporting/app/api"
	"event-reporting/app/helpers/registry"
)

func main() {
	routes, cfg := registry.Build()
	r := api.NewRouter(*routes)

	addr := ":" + cfg.AppPort
	log.Printf("listening on %s (env=%s) ...", addr, cfg.AppEnv)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
