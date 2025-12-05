package detalle_pedido

type DetallePedido struct {
	NumPedido    int    `gorm:"primaryKey" json:"num_pedido"`
	CodProducto  int    `gorm:"primaryKey" json:"cod_producto"`
	Cantidad     int    `json:"cantidad"`
	Medidas      string `gorm:"size:200" json:"medidas"`
	Observaciones string `gorm:"size:200" json:"observaciones"`
}

