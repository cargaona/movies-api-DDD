package sql

import (
	"database/sql"
	"github.com/cargaona/movies-api-DDD/pkg/domain/adding"
	"github.com/cargaona/movies-api-DDD/pkg/domain/deleting"
	"github.com/cargaona/movies-api-DDD/pkg/domain/listing"
	"github.com/cargaona/movies-api-DDD/pkg/storage"
	"log"
	"time"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func (s *Storage) DeleteMovie(movie deleting.Movie) error {
	panic("implement me")
}

func (s *Storage) GetAllMovies() ([]listing.Movie, error) {
	panic("implement me")
}

func NewStorage() (*Storage, error) {
	uri := "postgres://movies:example@127.0.0.1:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	s := new(Storage)
	s.db = db
	return s, nil
}

func (s *Storage) AddMovie(m adding.Movie) error {
	id, err := storage.GetID("movie")
	if err != nil {
		log.Fatal(err)
	}
	newMovie := Movie{
		ID:       id,
		Name:     m.Name,
		Director: director{m.Director.Name, m.Director.LastName},
		Year:     m.Year,
		Created:  time.Now(),
	}
	query := "INSERT INTO movies (id, name, director_name, director_last_name, year, created) VALUES ($1, $2, $3 ,$4, $5) RETURNING id;"
	row := s.db.QueryRow(query, newMovie.ID, newMovie.Name, newMovie.Director.Name, newMovie.Director.LastName, newMovie.Year, newMovie.Created)
	
	errScan := row.Scan(id)
	if errScan != nil {
		return err
	}
	return nil
}
