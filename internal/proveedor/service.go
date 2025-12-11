package proveedor

type ProveedorService struct {
	repo *ProveedorRepository
}

func NewProveedorService(repo *ProveedorRepository) *ProveedorService {
	return &ProveedorService{repo: repo}
}

func (s *ProveedorService) GetAllProveedores() ([]Proveedor, error) {
	return s.repo.GetAll()
}

func (s *ProveedorService) GetProveedor(nit int) (*Proveedor, error) {
	return s.repo.GetByNIT(nit)
}

func (s *ProveedorService) CreateProveedor(p *Proveedor) error {
	return s.repo.Create(p)
}

func (s *ProveedorService) UpdateProveedor(p *Proveedor) error {
	return s.repo.Update(p)
}

func (s *ProveedorService) DeleteProveedor(nit int) error {
	return s.repo.Delete(nit)
}

func (s *ProveedorService) UpdateProveedorPartial(nit int, updates map[string]interface{}) error {
	return s.repo.UpdateProveedorPartial(nit, updates)
}

