package reportes


type PedidoPendiente struct {
    NumPedido       int       `json:"num_pedido"`
    FechaEncargo    string    `json:"fecha_encargo"`
    FechaEntrega    string    `json:"fecha_entrega"`
    Estado          string    `json:"estado"`
    ClienteNombre   string    `json:"cliente_nombre"`
    ClienteDocumento int      `json:"cliente_documento"`
    Productos       []Producto `json:"productos"`
}


type ColegioUniforme struct {
    IDColegio     int       `json:"id_colegio"`
    NombreColegio string    `json:"nombre_colegio"`
    Direccion     string    `json:"direccion"`
    Telefono      string    `json:"telefono"`
    Uniformes     []Producto `json:"uniformes"`
}


type TotalVentas struct {
    Total float64 `json:"total"`
}

