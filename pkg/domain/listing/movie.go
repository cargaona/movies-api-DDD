package listing

type Movie struct {
	Id       string
	Name     string
	Year     int
	Director Director
}

type Director struct {
	Name     string
	LastName string
}
