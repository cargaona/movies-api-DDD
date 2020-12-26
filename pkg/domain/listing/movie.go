package listing

type Movie struct {
	Name     string
	Year     int
	Director director
}

type director struct {
	Name     string
	LastName string
}
