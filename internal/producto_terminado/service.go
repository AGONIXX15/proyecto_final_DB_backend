package producto_terminado


type ProductoTerminadoService interface {
    Create(producto *ProductoTerminado) error
    GetAll() ([]ProductoTerminado, error)
    GetByID(id uint) (*ProductoTerminado, error)
    Update(producto *ProductoTerminado) error
	UpdatePartial(id uint, fields map[string]interface{}) error
    Delete(id uint) error
}

type productoTerminadoService struct {
    repo ProductoTerminadoRepository
}

func NewProductoTerminadoService(repo ProductoTerminadoRepository) ProductoTerminadoService {
    return &productoTerminadoService{repo}
}

func (s *productoTerminadoService) Create(producto *ProductoTerminado) error {
    return s.repo.Create(producto)
}

func (s *productoTerminadoService) GetAll() ([]ProductoTerminado, error) {
    return s.repo.GetAll()
}

func (s *productoTerminadoService) GetByID(id uint) (*ProductoTerminado, error) {
    return s.repo.GetByID(id)
}

func (s *productoTerminadoService) Update(producto *ProductoTerminado) error {
    return s.repo.Update(producto)
}

func (s *productoTerminadoService) Delete(id uint) error {
    return s.repo.Delete(id)
}

func (s *productoTerminadoService) UpdatePartial(id uint, fields map[string]interface{}) error {
    return s.repo.UpdatePartial(id, fields)
}
