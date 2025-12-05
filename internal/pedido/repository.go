package pedido

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrNotFound   = errors.New("pedido no encontrado")
	ErrDBInternal = errors.New("error interno en la base de datos")
)

type PedidoRepository struct {
	db *gorm.DB
}

func NewPedidoRepository(db *gorm.DB) *PedidoRepository {
	return &PedidoRepository{db: db}
}

func (r *PedidoRepository) GetAll() ([]Pedido, error) {
	var pedidos []Pedido
	if err := r.db.Find(&pedidos).Error; err != nil {
		return nil, ErrDBInternal
	}
	return pedidos, nil
}

func (r *PedidoRepository) GetByID(id int) (*Pedido, error) {
	var pedido Pedido
	if err := r.db.First(&pedido, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrDBInternal
	}
	return &pedido, nil
}

func (r *PedidoRepository) Create(p *Pedido) error {
	if err := r.db.Create(p).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *PedidoRepository) Update(p *Pedido) error {
	if err := r.db.Save(p).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *PedidoRepository) Delete(id int) error {
	if err := r.db.Delete(&Pedido{}, id).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

