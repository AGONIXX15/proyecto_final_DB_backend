package detalle_pedido

type DetallePedidoService struct {
	repo *DetallePedidoRepository
}

func NewDetallePedidoService(repo *DetallePedidoRepository) *DetallePedidoService {
	return &DetallePedidoService{repo: repo}
}

func (s *DetallePedidoService) GetAllDetalles() ([]DetallePedido, error) {
	return s.repo.GetAll()
}

func (s *DetallePedidoService) GetDetalle(numPedido int, codProducto int) (*DetallePedido, error) {
	return s.repo.GetByID(numPedido, codProducto)
}

func (s *DetallePedidoService) CreateDetalle(d *DetallePedido) error {
	return s.repo.Create(d)
}

func (s *DetallePedidoService) UpdateDetalle(d *DetallePedido) error {
	return s.repo.Update(d)
}

func (s *DetallePedidoService) DeleteDetalle(numPedido int, codProducto int) error {
	return s.repo.Delete(numPedido, codProducto)
}

