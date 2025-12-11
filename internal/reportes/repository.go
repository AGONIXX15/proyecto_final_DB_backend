package reportes

import (
    "gorm.io/gorm"
)

type ReporteRepository interface {
    GetProductosPendientes() ([]ProductoPendiente, error)
    GetProductosPorCliente(documento int) ([]ProductoCliente, error)
    GetProductosDisponibles() ([]ProductoDisponible, error)
    GetColegiosConUniformes() ([]ColegioUniforme, error)
    GetUniformesPorColegio() ([]UniformeColegio, error)
    GetVentasPorColegio() ([]VentasPorColegio, error)
    GetTotalVentas() (float64, error)
}

type reporteRepo struct {
    db *gorm.DB
}

func NewReporteRepo(db *gorm.DB) ReporteRepository {
    return &reporteRepo{db}
}

func (r *reporteRepo) GetProductosPendientes() ([]ProductoPendiente, error) {
    var list []ProductoPendiente
    err := r.db.Raw("SELECT * FROM productos_pendientes").Scan(&list).Error
    return list, err
}

func (r *reporteRepo) GetProductosPorCliente(documento int) ([]ProductoCliente, error) {
    var list []ProductoCliente
    err := r.db.Raw("SELECT * FROM productos_cliente WHERE documento = ?", documento).Scan(&list).Error
    return list, err
}

func (r *reporteRepo) GetProductosDisponibles() ([]ProductoDisponible, error) {
    var list []ProductoDisponible
    err := r.db.Raw("SELECT * FROM productos_disponibles").Scan(&list).Error
    return list, err
}

func (r *reporteRepo) GetColegiosConUniformes() ([]ColegioUniforme, error) {
    var list []ColegioUniforme
    err := r.db.Raw("SELECT * FROM colegios_uniformes").Scan(&list).Error
    return list, err
}

func (r *reporteRepo) GetUniformesPorColegio() ([]UniformeColegio, error) {
    var list []UniformeColegio
    err := r.db.Raw("SELECT * FROM uniforme_colegio").Scan(&list).Error
    return list, err
}

func (r *reporteRepo) GetVentasPorColegio() ([]VentasPorColegio, error) {
    var list []VentasPorColegio
    err := r.db.Raw("SELECT * FROM ventas_por_colegio").Scan(&list).Error
    return list, err
}

func (r *reporteRepo) GetTotalVentas() (float64, error) {
    var tv TotalVentas
    err := r.db.Raw("SELECT * FROM total_ventas").Scan(&tv).Error
    return tv.Total, err
}

