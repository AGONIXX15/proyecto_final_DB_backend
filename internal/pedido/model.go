package pedido

import (
	"time"
)

type Pedido struct {
	NumPedido    int       `gorm:"primaryKey" json:"num_pedido"`
	FechaEncargo time.Time `json:"fecha_encargo"`
	FechaEntrega time.Time `json:"fecha_entrega"`
	Abono        float64   `json:"abono"`
	Estado       string    `gorm:"size:20" json:"estado"`
	DocCliente   int       `json:"doc_cliente"`
}

