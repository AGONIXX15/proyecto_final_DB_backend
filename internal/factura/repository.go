package factura

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	ErrNotFound   = errors.New("factura no encontrada")
	ErrDBInternal = errors.New("error interno en la base de datos")
)

type FacturaRepository struct {
	db *gorm.DB
}

func NewFacturaRepository(db *gorm.DB) *FacturaRepository {
	return &FacturaRepository{db: db}
}

func (r *FacturaRepository) GetAll() ([]Factura, error) {
	var facturas []Factura
	if err := r.db.Find(&facturas).Error; err != nil {
		fmt.Println(err)
		return nil, ErrDBInternal
	}
	return facturas, nil
}

func (r *FacturaRepository) GetByID(id int) (*Factura, error) {
	var factura Factura
	if err := r.db.First(&factura, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrDBInternal
	}
	return &factura, nil
}

func (r *FacturaRepository) Create(f *Factura) error {
	if err := r.db.Create(f).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *FacturaRepository) Update(f *Factura) error {
	if err := r.db.Save(f).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *FacturaRepository) Delete(id int) error {
	if err := r.db.Delete(&Factura{}, id).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}
func (r *FacturaRepository) UpdatePartial(numFactura int, updates map[string]interface{}) error {
    result := r.db.
        Model(&Factura{}).
        Where("num_factura = ?", numFactura).
        Updates(updates)

    if result.Error != nil {
        return fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
    }

    if result.RowsAffected == 0 {
        return fmt.Errorf("%w: factura con num_factura %d no encontrada", ErrNotFound, numFactura)
    }

    return nil
}

