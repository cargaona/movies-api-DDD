package main

import (
	"github.com/cargaona/movies-api-DDD/pkg/adding"
	"github.com/cargaona/movies-api-DDD/pkg/deleting"
	rest "github.com/cargaona/movies-api-DDD/pkg/http"
	"github.com/cargaona/movies-api-DDD/pkg/listing"
	"github.com/cargaona/movies-api-DDD/pkg/storage/json"
	"log"
	"net/http"
)

func main(){
	s, _ := json.NewStorage()
	adder := adding.NewService(s)
	lister := listing.NewService(s)
	deleter := deleting.NewService(s)
	router := rest.Handler(adder, lister, deleter)
	log.Fatal(http.ListenAndServe(":8080", router))
}