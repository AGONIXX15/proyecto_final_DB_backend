package colegio

type ColegioService struct {
	repo *ColegioRepository
}

func NewColegioService(repo *ColegioRepository) *ColegioService {
	return &ColegioService{repo: repo}
}

func (s *ColegioService) GetAllColegios() ([]Colegio, error) {
	return s.repo.GetAll()
}

func (s *ColegioService) GetColegio(id int) (*Colegio, error) {
	return s.repo.GetByID(id)
}

func (s *ColegioService) CreateColegio(c *Colegio) error {
	return s.repo.Create(c)
}

func (s *ColegioService) UpdateColegio(c *Colegio) error {
	return s.repo.Update(c)
}

func (s *ColegioService) DeleteColegio(id int) error {
	return s.repo.Delete(id)
}

func (s *ColegioService) UpdateColegioPartial(id uint, updates map[string]interface{}) error {
    return s.repo.UpdatePartial(id, updates)
}

