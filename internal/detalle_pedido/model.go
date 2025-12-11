package detalle_pedido

type DetallePedido struct {
	NumPedido     int    `gorm:"primaryKey" json:"num_pedido"`
	TypeItem      string `gorm:"primaryKey;check: type_item IN ('uniforme', 'producto')" json:"type_item"`
	CodItem       int    `gorm:"primaryKey" json:"cod_item"`
	Cantidad      int    `json:"cantidad"`
	Medidas       string `gorm:"size:200" json:"medidas"`
	Observaciones string `gorm:"size:200" json:"observaciones"`
}
