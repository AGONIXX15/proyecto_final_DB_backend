package factura

import "time"

type Factura struct {
	NumFactura int     `gorm:"primaryKey" json:"num_factura"`
	Fecha      time.Time `json:"fecha"`
	Total      float64 `json:"total"`
	NumPedido  int     `json:"num_pedido"`
}
