package detalle_pedido

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrNotFound   = errors.New("detalle pedido no encontrado")
	ErrDBInternal = errors.New("error interno en la base de datos")
)

type DetallePedidoRepository struct {
	db *gorm.DB
}

func NewDetallePedidoRepository(db *gorm.DB) *DetallePedidoRepository {
	return &DetallePedidoRepository{db: db}
}

func (r *DetallePedidoRepository) GetAll() ([]DetallePedido, error) {
	var detalles []DetallePedido
	if err := r.db.Find(&detalles).Error; err != nil {
		return nil, ErrDBInternal
	}
	return detalles, nil
}

func (r *DetallePedidoRepository) GetByID(numPedido int, codProducto int) (*DetallePedido, error) {
	var detalle DetallePedido
	if err := r.db.First(&detalle, "num_pedido = ? AND cod_producto = ?", numPedido, codProducto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrDBInternal
	}
	return &detalle, nil
}

func (r *DetallePedidoRepository) Create(d *DetallePedido) error {
	if err := r.db.Create(d).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *DetallePedidoRepository) Update(d *DetallePedido) error {
	if err := r.db.Save(d).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *DetallePedidoRepository) Delete(numPedido int, codProducto int) error {
	if err := r.db.Delete(&DetallePedido{}, "num_pedido = ? AND cod_producto = ?", numPedido, codProducto).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

