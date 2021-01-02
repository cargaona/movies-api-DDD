package main

import (
	"github.com/cargaona/movies-api-DDD/pkg/domain/adding"
	"github.com/cargaona/movies-api-DDD/pkg/domain/deleting"
	"github.com/cargaona/movies-api-DDD/pkg/domain/listing"
	"github.com/cargaona/movies-api-DDD/pkg/storage/sql"

	rest "github.com/cargaona/movies-api-DDD/pkg/http"
	"log"
	"net/http"
)

func main() {
	// Initialize DB
	s, _ := sql.NewStorage()

	// Initialize Services with the created DB.
	adder := adding.NewService(s)
	lister := listing.NewService(s)
	deleter := deleting.NewService(s)

	// Create and initialize HTTP server.
	router := rest.Handler(adder, lister, deleter)
	log.Fatal(http.ListenAndServe(":8081", router))
}
