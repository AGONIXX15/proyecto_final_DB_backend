package uniforme

type UniformeService struct {
	repo *UniformeRepository
}

func NewUniformeService(repo *UniformeRepository) *UniformeService {
	return &UniformeService{repo: repo}
}

func (s *UniformeService) GetAllUniformes() ([]Uniforme, error) {
	return s.repo.GetAll()
}

func (s *UniformeService) GetUniforme(id int) (*Uniforme, error) {
	return s.repo.GetByID(id)
}

func (s *UniformeService) CreateUniforme(u *Uniforme) error {
	return s.repo.Create(u)
}

func (s *UniformeService) UpdateUniforme(u *Uniforme) error {
	return s.repo.Update(u)
}

func (s *UniformeService) DeleteUniforme(id int) error {
	return s.repo.Delete(id)
}

