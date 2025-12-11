package reportes

type ReporteService struct {
    repo *ReporteRepository
}

func NewReporteService(repo *ReporteRepository) *ReporteService {
    return &ReporteService{repo: repo}
}

// func (s *ReporteService) PedidosPendientes() ([]PedidoPendiente, error) {
//     return s.repo.GetPedidosPendientes()
// }

func (s *ReporteService) ProductosPorCliente(documento int) ([]PedidoCliente, error) {
    return s.repo.GetProductosPorCliente(documento)
}

func (s *ReporteService) ColegiosUniformes() ([]ColegioUniforme, error) {
    return s.repo.GetColegiosUniformes()
}

func (s *ReporteService) VentasColegio() ([]ColegioVenta, error) {
    return s.repo.GetVentasColegio()
}

func (s *ReporteService) TotalVentas() (TotalVentas, error) {
    return s.repo.GetTotalVentas()
}

