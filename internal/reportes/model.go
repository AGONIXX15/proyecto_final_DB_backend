package reportes


import "time"

type ProductoPendiente struct {
    NumPedido    int        `gorm:"column:num_pedido"`
    CodItem      int        `gorm:"column:cod_item"`
    TypeItem     string     `gorm:"column:type_item"`
    Cantidad     int        `gorm:"column:cantidad"`
    FechaEntrega *time.Time `gorm:"column:fecha_entrega"`
}

type ProductoCliente struct {
    Documento      int    `gorm:"column:documento"`
    NombreCompleto string `gorm:"column:nombre_completo"`
    CodItem        int    `gorm:"column:cod_item"`
    TypeItem       string `gorm:"column:type_item"`
    Cantidad       int    `gorm:"column:cantidad"`
    NumPedido      int    `gorm:"column:num_pedido"`
}

type ProductoDisponible struct {
    CodProducto        int    `gorm:"column:cod_producto"`
    Nombre             string `gorm:"column:nombre"`
    CantidadDisponible int    `gorm:"column:cantidad_disponible"`
}

type ColegioUniforme struct {
    ID     int    `gorm:"column:id"`
    Nombre string `gorm:"column:nombre"`
}

type UniformeColegio struct {
    IDColegio int    `gorm:"column:id_colegio"`
    Tipo      string `gorm:"column:tipo_uniforme"`
    Color     string `gorm:"column:color"`
    TipoTela  string `gorm:"column:tipo_tela"`
    Bordado   string `gorm:"column:bordado"`
    Estampado string `gorm:"column:estampado"`
    Detalles  string `gorm:"column:detalles"`
}

type VentasPorColegio struct {
    IDColegio     int    `gorm:"column:id_colegio"`
    ColegioNombre string `gorm:"column:colegio_nombre"`
    TotalVendido  int    `gorm:"column:total_vendido"`
}

type TotalVentas struct {
    Total float64 `gorm:"column:total_ventas"`
}

