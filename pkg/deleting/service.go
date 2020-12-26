package deleting

type Service interface {
	DeleteMovie(Movie) error
}

type Repository interface {
	DeleteMovie(Movie) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteMovie(m Movie) error {
	return s.r.DeleteMovie(m)
}
