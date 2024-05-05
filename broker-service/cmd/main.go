package main

import (
	"fmt"
	"log"
	"net/http"

	security "github.com/Dimoonevs/SportsApp/broker-service/internal/routers"
)

const webPort = "8000"

func main() {
	log.Printf("Starting Broker Service on port %s", webPort)

	auth := security.AppBroker{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: auth.Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
