package materia_prima

type MateriaPrimaService struct {
	repo *MateriaPrimaRepository
}

func NewMateriaPrimaService(repo *MateriaPrimaRepository) *MateriaPrimaService {
	return &MateriaPrimaService{repo: repo}
}

func (s *MateriaPrimaService) GetAllMaterias() ([]MateriaPrima, error) {
	return s.repo.GetAll()
}

func (s *MateriaPrimaService) GetMateria(id int) (*MateriaPrima, error) {
	return s.repo.GetByID(id)
}

func (s *MateriaPrimaService) CreateMateria(m *MateriaPrima) error {
	return s.repo.Create(m)
}

func (s *MateriaPrimaService) UpdateMateria(m *MateriaPrima) error {
	return s.repo.Update(m)
}

func (s *MateriaPrimaService) DeleteMateria(id int) error {
	return s.repo.Delete(id)
}

