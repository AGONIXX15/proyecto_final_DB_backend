package detalle_pedido

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	ErrNotFound   = errors.New("detalle de pedido no encontrado")
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
	err := r.db.First(
		&detalle,
		"num_pedido = ? AND cod_producto = ?",
		numPedido, codProducto,
	).Error

	if err != nil {
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

func (r *DetallePedidoRepository) UpdatePartial(numPedido int, typeItem string, codItem int, updates map[string]interface{}) error {
    result := r.db.
        Model(&DetallePedido{}).
        Where("num_pedido = ? AND type_item = ? AND cod_item = ?", numPedido, typeItem, codItem).
        Updates(updates)

    if result.Error != nil {
        return fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
    }

    if result.RowsAffected == 0 {
        return fmt.Errorf("%w: detalle pedido (%d, %s, %d) no encontrado", ErrNotFound, numPedido, typeItem, codItem)
    }

    return nil
}


