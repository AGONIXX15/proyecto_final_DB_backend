package reportes

type ReporteService struct {
    repo ReporteRepository
}

func NewReporteService(r ReporteRepository) *ReporteService {
    return &ReporteService{repo: r}
}

func (s *ReporteService) ProductosPendientes() ([]ProductoPendiente, error) {
    return s.repo.GetProductosPendientes()
}

func (s *ReporteService) ProductosPorCliente(documento int) ([]ProductoCliente, error) {
    return s.repo.GetProductosPorCliente(documento)
}

func (s *ReporteService) ProductosDisponibles() ([]ProductoDisponible, error) {
    return s.repo.GetProductosDisponibles()
}

func (s *ReporteService) ColegiosConUniformes() ([]ColegioUniforme, error) {
    return s.repo.GetColegiosConUniformes()
}

func (s *ReporteService) UniformesPorColegio() ([]UniformeColegio, error) {
    return s.repo.GetUniformesPorColegio()
}

func (s *ReporteService) VentasPorColegio() ([]VentasPorColegio, error) {
    return s.repo.GetVentasPorColegio()
}

func (s *ReporteService) TotalVentas() (float64, error) {
    return s.repo.GetTotalVentas()
}

