package listing

type Service interface {
	GetAllMovies() ([]Movie,error)
}

type Repository interface {
	GetAllMovies() ([]Movie, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAllMovies() ([]Movie, error){
	return s.r.GetAllMovies()
}

