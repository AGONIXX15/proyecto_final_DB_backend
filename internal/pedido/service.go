package pedido

type PedidoService struct {
	repo *PedidoRepository
}

func NewPedidoService(repo *PedidoRepository) *PedidoService {
	return &PedidoService{repo: repo}
}

func (s *PedidoService) GetAllPedidos() ([]Pedido, error) {
	return s.repo.GetAll()
}

func (s *PedidoService) GetPedido(id int) (*Pedido, error) {
	return s.repo.GetByID(id)
}

func (s *PedidoService) CreatePedido(p *Pedido) error {
	return s.repo.Create(p)
}

func (s *PedidoService) UpdatePedido(p *Pedido) error {
	return s.repo.Update(p)
}

func (s *PedidoService) DeletePedido(id int) error {
	return s.repo.Delete(id)
}

