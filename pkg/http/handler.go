package rest

import (
	"github.com/cargaona/movies-api-DDD/pkg/domain/adding"
	"github.com/cargaona/movies-api-DDD/pkg/domain/deleting"
	"github.com/cargaona/movies-api-DDD/pkg/domain/listing"

	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Handler(a adding.Service, l listing.Service, d deleting.Service) http.Handler {
	router := httprouter.New()

	router.POST("/movies", addMovie(a))
	router.GET("/movies", getAllMovies(l))
	router.DELETE("/movies", deleteMovie(d))

	return router
}

func getAllMovies(l listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		list, err := l.GetAllMovies()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Context-Type", "application/json")
		json.NewEncoder(w).Encode(list)
	}
}
func addMovie(s adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var newMovie adding.Movie
		errorDecoding := decoder.Decode(&newMovie)
		if errorDecoding != nil {
			http.Error(w, errorDecoding.Error(), http.StatusBadRequest)

			return
		}

		errorAdding := s.AddMovie(newMovie)
		if errorAdding != nil{
			http.Error(w, errorAdding.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Internal Server Error")
			return
		}

		w.Header().Set("Context-Type", "application/json")
		json.NewEncoder(w).Encode("New Movie Added")
	}
}
func deleteMovie(d deleting.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)
		var deletedMovie deleting.Movie
		err := decoder.Decode(&deletedMovie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := d.DeleteMovie(deletedMovie); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			json.NewEncoder(w).Encode("Movie Not Found")
			return
		}

		w.Header().Set("Context-Type", "application/json")
		json.NewEncoder(w).Encode("Movie Deleted")
	}
}
