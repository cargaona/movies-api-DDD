package adding

type Service interface {
	AddMovie(Movie) error
}

type Repository interface {
	AddMovie(Movie) error
//	GetAllMovies() []listing.Movie
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddMovie(m Movie) error{
	//Some validation must be done here
	s.r.AddMovie(m)
	return nil
}

