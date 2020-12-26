package json

import "time"

type Movie struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Director director  `json:"director"`
	Year     int       `json:"year"`
	Created  time.Time `json:"created"`
}

type director struct {
	Name     string
	LastName string
}
