package factura

type FacturaService struct {
	repo *FacturaRepository
}

func NewFacturaService(repo *FacturaRepository) *FacturaService {
	return &FacturaService{repo: repo}
}

func (s *FacturaService) GetAllFacturas() ([]Factura, error) {
	return s.repo.GetAll()
}

func (s *FacturaService) GetFactura(id int) (*Factura, error) {
	return s.repo.GetByID(id)
}

func (s *FacturaService) CreateFactura(f *Factura) error {
	return s.repo.Create(f)
}

func (s *FacturaService) UpdateFactura(f *Factura) error {
	return s.repo.Update(f)
}

func (s *FacturaService) DeleteFactura(id int) error {
	return s.repo.Delete(id)
}

