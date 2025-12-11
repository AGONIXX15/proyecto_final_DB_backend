package uniforme

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	ErrNotFound   = errors.New("uniforme no encontrado")
	ErrDBInternal = errors.New("error interno en la base de datos")
)

type UniformeRepository struct {
	db *gorm.DB
}

func NewUniformeRepository(db *gorm.DB) *UniformeRepository {
	return &UniformeRepository{db: db}
}

func (r *UniformeRepository) GetAll() ([]Uniforme, error) {
	var uniformes []Uniforme
	if err := r.db.Find(&uniformes).Error; err != nil {
		return nil, ErrDBInternal
	}
	return uniformes, nil
}

func (r *UniformeRepository) GetByID(id int) (*Uniforme, error) {
	var uniforme Uniforme
	if err := r.db.First(&uniforme, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrDBInternal
	}
	return &uniforme, nil
}

func (r *UniformeRepository) Create(u *Uniforme) error {
	if err := r.db.Create(u).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *UniformeRepository) Update(u *Uniforme) error {
	if err := r.db.Save(u).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *UniformeRepository) Delete(id int) error {
	if err := r.db.Delete(&Uniforme{}, id).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *UniformeRepository) UpdateUniformePartial(id int, updates map[string]interface{}) error {
    result := r.db.Model(&Uniforme{}).Where("id = ?", id).Updates(updates)

    if result.Error != nil {
        return fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
    }

    if result.RowsAffected == 0 {
        return fmt.Errorf("%w: uniforme con id %d no encontrado", ErrNotFound, id)
    }

    return nil
}

