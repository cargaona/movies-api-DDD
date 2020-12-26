package main

import (
	"github.com/cargaona/movies-api-DDD/pkg/domain/adding"
	"github.com/cargaona/movies-api-DDD/pkg/domain/deleting"
	"github.com/cargaona/movies-api-DDD/pkg/domain/listing"

	rest "github.com/cargaona/movies-api-DDD/pkg/http"
	"github.com/cargaona/movies-api-DDD/pkg/storage/json"

	"log"
	"net/http"
)

func main() {
	// Initialize DB
	s, _ := json.NewStorage()

	// Initialize Services with the created DB.
	adder := adding.NewService(s)
	lister := listing.NewService(s)
	deleter := deleting.NewService(s)

	// Create and initialize HTTP server.
	router := rest.Handler(adder, lister, deleter)
	log.Fatal(http.ListenAndServe(":8080", router))
}
