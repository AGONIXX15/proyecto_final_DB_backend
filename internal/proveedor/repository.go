package proveedor

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	ErrNotFound   = errors.New("proveedor no encontrado")
	ErrDBInternal = errors.New("error interno en la base de datos")
)

type ProveedorRepository struct {
	db *gorm.DB
}

func NewProveedorRepository(db *gorm.DB) *ProveedorRepository {
	return &ProveedorRepository{db: db}
}

func (r *ProveedorRepository) GetAll() ([]Proveedor, error) {
	var proveedores []Proveedor
	if err := r.db.Find(&proveedores).Error; err != nil {
		return nil, ErrDBInternal
	}
	return proveedores, nil
}

func (r *ProveedorRepository) GetByNIT(nit int) (*Proveedor, error) {
	var proveedor Proveedor
	if err := r.db.First(&proveedor, nit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrDBInternal
	}
	return &proveedor, nil
}

func (r *ProveedorRepository) Create(p *Proveedor) error {
	if err := r.db.Create(p).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *ProveedorRepository) Update(p *Proveedor) error {
	if err := r.db.Save(p).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *ProveedorRepository) Delete(nit int) error {
	if err := r.db.Delete(&Proveedor{}, nit).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *ProveedorRepository) UpdateProveedorPartial(nit int, updates map[string]interface{}) error {
    result := r.db.Model(&Proveedor{}).Where("nit = ?", nit).Updates(updates)

    if result.Error != nil {
        return fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("%w: proveedor con NIT %d no encontrado", ErrNotFound, nit)
    }
    return nil
}


