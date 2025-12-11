package producto_terminado

import (
	"gorm.io/gorm"
)

type ProductoTerminadoRepository interface {
	Create(producto *ProductoTerminado) error
	GetAll() ([]ProductoTerminado, error)
	GetByID(id uint) (*ProductoTerminado, error)
	Update(producto *ProductoTerminado) error
	UpdatePartial(id uint, fields map[string]interface{}) error
	Delete(id uint) error
}

type productoTerminadoRepository struct {
	db *gorm.DB
}

func NewProductoTerminadoRepository(db *gorm.DB) ProductoTerminadoRepository {
	return &productoTerminadoRepository{db}
}

func (r *productoTerminadoRepository) Create(producto *ProductoTerminado) error {
	return r.db.Create(producto).Error
}

func (r *productoTerminadoRepository) GetAll() ([]ProductoTerminado, error) {
	var productos []ProductoTerminado
	err := r.db.Find(&productos).Error
	return productos, err
}

func (r *productoTerminadoRepository) GetByID(id uint) (*ProductoTerminado, error) {
	var producto ProductoTerminado
	err := r.db.First(&producto, id).Error
	if err != nil {
		return nil, err
	}
	return &producto, nil
}

func (r *productoTerminadoRepository) Update(producto *ProductoTerminado) error {
	return r.db.Save(producto).Error
}

func (r *productoTerminadoRepository) Delete(id uint) error {
	return r.db.Delete(&ProductoTerminado{}, id).Error
}

func (r *productoTerminadoRepository) UpdatePartial(id uint, fields map[string]interface{}) error {
	return r.db.Model(&ProductoTerminado{}).
		Where("cod_producto = ?", id).
		Updates(fields).Error
}
