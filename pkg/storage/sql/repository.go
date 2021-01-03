package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"os"

	"github.com/cargaona/movies-api-DDD/pkg/domain/adding"
	"github.com/cargaona/movies-api-DDD/pkg/domain/deleting"
	"github.com/cargaona/movies-api-DDD/pkg/domain/listing"
	"github.com/cargaona/movies-api-DDD/pkg/storage"

	"log"
	"time"
)

type Storage struct {
	db *sql.DB
}

type config struct {
	User     string
	Password string
	Host     string
	Database string
}

func NewStorage() (*Storage, error) {
	config := configDB()
	uri := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", config.User, config.Password, config.Host, config.Database)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	s := new(Storage)
	s.db = db
	return s, nil
}

func (s *Storage) DeleteMovie(movie deleting.Movie) error {
	query := "DELETE FROM MOVIES WHERE id = $1"
	_, err := s.db.Exec(query, movie.Id)
	if err != nil {
		return errors.New("Error")
	}
	return nil
}

func (s *Storage) GetAllMovies() ([]listing.Movie, error) {
	query := "SELECT * FROM movies"
	row, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	var movies []listing.Movie
	for row.Next() {
		var movie Movie
		errScan := row.Scan(&movie.ID, &movie.Name, &movie.Director.Name, &movie.Director.LastName, &movie.Year, &movie.Created)
		if errScan != nil {
			return nil, errScan
		}
		movies = append(movies, listing.Movie{
			Id:   movie.ID,
			Name: movie.Name,
			Year: movie.Year,
			Director: listing.Director{
				Name:     movie.Director.Name,
				LastName: movie.Director.LastName,
			},
		})
	}
	return movies, nil
}

func (s *Storage) AddMovie(m adding.Movie) error {
	id, err := storage.GetID("movie")
	if err != nil {
		log.Fatal(err)
	}

	query := "INSERT INTO movies (id, name, director_name, director_last_name, year, created) VALUES ($1, $2, $3 ,$4, $5, $6) RETURNING id;"
	row := s.db.QueryRow(query, id, m.Name, m.Director.Name, m.Director.LastName, m.Year, time.Now())

	errScan := row.Scan(&id)
	if errScan != nil {
		return errScan
	}
	return nil
}

func configDB() config {
	config := config{
		User:     os.Getenv("MOVIESAPI_POSTGRES_USER"),
		Password: os.Getenv("MOVIESAPI_POSTGRES_PASSWORD"),
		Host:     os.Getenv("MOVIESAPI_POSTGRES_HOST"),
		Database: os.Getenv("MOVIESAPI_POSTGRES_DATABASE"),
	}
	if config.User == "" {
		config.User = "postgres"
	}
	if config.Password == "" {
		config.Password = "postgres"
	}
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}
	if config.Database == "" {
		config.Database = "postgres"
	}
	return config
}
