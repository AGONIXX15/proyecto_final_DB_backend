package pedido

import "github.com/AGONIXX15/db_proyecto_final/internal/detalle_pedido"

type PedidoService struct {
	repo *PedidoRepository
}

func NewPedidoService(repo *PedidoRepository) *PedidoService {
	return &PedidoService{
		repo: repo,
	}
}

// CRUD básico usando repository
func (s *PedidoService) GetAllPedidos() ([]Pedido, error) {
	return s.repo.GetAll()
}

func (s *PedidoService) GetPedidoByID(id int) (*Pedido, error) {
	return s.repo.GetByID(id)
}

// Crear pedido usando función SQL
func (s *PedidoService) CreatePedido(pedido *Pedido, detalles []detalle_pedido.DetallePedido) error {
	return s.repo.CreateWithFunction(pedido, detalles)
}

// Actualizar pedido usando función SQL
func (s *PedidoService) UpdatePedido(numPedido int, updatedPedido *Pedido, nuevosDetalles []detalle_pedido.DetallePedido) error {
	return s.repo.UpdateWithFunction(numPedido, updatedPedido, nuevosDetalles)
}

// Eliminar/cancelar pedido
func (s *PedidoService) DeletePedido(numPedido int) error {
	return s.repo.CancelWithFunction(numPedido)
}

