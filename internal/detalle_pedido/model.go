package detalle_pedido

type DetallePedido struct {
	NumPedido      int     `gorm:"primaryKey;column:num_pedido" json:"num_pedido"`
	TypeItem       string  `gorm:"primaryKey;column:type_item;check: type_item IN ('producto','uniforme')" json:"tipo_producto"` // ahora coincide con frontend
	CodItem        int     `gorm:"primaryKey;column:cod_item" json:"cod_producto"` // ahora coincide con frontend
	Cantidad       int     `gorm:"column:cantidad" json:"cantidad"`
	Medidas        string  `gorm:"column:medidas" json:"medidas"`
	Observaciones  string  `gorm:"column:observaciones" json:"observaciones"`
	PrecioUnitario float64 `gorm:"column:precio_unitario" json:"precio_unidad"` // ahora coincide con frontend
	Subtotal       float64 `gorm:"column:subtotal" json:"subtotal"`
}

