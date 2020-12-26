package json

import (
	"encoding/json"
	"github.com/cargaona/movies-api-DDD/pkg/adding"
	"github.com/cargaona/movies-api-DDD/pkg/deleting"
	"github.com/cargaona/movies-api-DDD/pkg/listing"
	"github.com/cargaona/movies-api-DDD/pkg/storage"
	"github.com/nanobox-io/golang-scribble"
	"github.com/pkg/errors"
	"log"
	"path"
	"runtime"
	"time"
)

type Storage struct {
	db *scribble.Driver
}

func NewStorage() (*Storage, error) {
	var err error
	s := new(Storage)
	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Storage) AddMovie(m adding.Movie) error {
	id, err := storage.GetID("movie")
	if err != nil {
		log.Fatal(err)
	}

	newB := Movie{
		ID:       id,
		Name:     m.Name,
		Director: director{m.Director.Name, m.Director.LastName},
		Year:     m.Year,
		Created:  time.Now(),
	}

	if err := s.db.Write("movies", newB.ID, newB); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetAllMovies() ([]listing.Movie, error) {
	list := []listing.Movie{}
	records, err := s.db.ReadAll("movies")
	if err != nil {
		return nil, err
	}
	for _, r := range records {
		var m Movie
		var movie listing.Movie

		if err := json.Unmarshal([]byte(r), &m); err != nil {
			return nil, err
		}
		movie.Director.LastName = m.Director.LastName
		movie.Director.Name = m.Director.Name
		movie.Name = m.Name
		movie.Year = m.Year

		list = append(list, movie)
	}
	return list, nil
}

func (s *Storage) DeleteMovie(m deleting.Movie) error {
	moviesList, _ := s.db.ReadAll("movies")
	for _, movies := range moviesList {
		var movie Movie
		if err := json.Unmarshal([]byte(movies), &movie); err != nil {
			return err
		}
		if m.Name == movie.Name {
			if err := s.db.Delete("movies", movie.ID); err != nil {
				return err
			}
			return nil	
		}
	}
	return errors.New("")
}
