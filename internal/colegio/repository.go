package colegio

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	ErrNotFound   = errors.New("colegio no encontrado")
	ErrDBInternal = errors.New("error interno en la base de datos")
)

type ColegioRepository struct {
	db *gorm.DB
}

func NewColegioRepository(db *gorm.DB) *ColegioRepository {
	return &ColegioRepository{db: db}
}

func (r *ColegioRepository) GetAll() ([]Colegio, error) {
	var colegios []Colegio
	if err := r.db.Find(&colegios).Error; err != nil {
		return nil, ErrDBInternal
	}
	return colegios, nil
}

func (r *ColegioRepository) GetByID(id int) (*Colegio, error) {
	var colegio Colegio
	if err := r.db.First(&colegio, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrDBInternal
	}
	return &colegio, nil
}

func (r *ColegioRepository) Create(c *Colegio) error {
	if err := r.db.Create(c).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *ColegioRepository) Update(c *Colegio) error {
	if err := r.db.Save(c).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *ColegioRepository) Delete(id int) error {
	if err := r.db.Delete(&Colegio{}, id).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *ColegioRepository) UpdatePartial(id uint, updates map[string]interface{}) error {
    result := r.db.Model(&Colegio{}).
        Where("id = ?", id).
        Updates(updates)

    if result.Error != nil {
        return fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
    }

    if result.RowsAffected == 0 {
        return fmt.Errorf("%w: colegio con id %d no encontrado", ErrNotFound, id)
    }

    return nil
}


