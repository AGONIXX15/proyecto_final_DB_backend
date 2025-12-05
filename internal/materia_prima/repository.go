package materia_prima

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrNotFound   = errors.New("materia prima no encontrada")
	ErrDBInternal = errors.New("error interno en la base de datos")
)

type MateriaPrimaRepository struct {
	db *gorm.DB
}

func NewMateriaPrimaRepository(db *gorm.DB) *MateriaPrimaRepository {
	return &MateriaPrimaRepository{db: db}
}

func (r *MateriaPrimaRepository) GetAll() ([]MateriaPrima, error) {
	var materias []MateriaPrima
	if err := r.db.Find(&materias).Error; err != nil {
		return nil, ErrDBInternal
	}
	return materias, nil
}

func (r *MateriaPrimaRepository) GetByID(id int) (*MateriaPrima, error) {
	var materia MateriaPrima
	if err := r.db.First(&materia, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, ErrDBInternal
	}
	return &materia, nil
}

func (r *MateriaPrimaRepository) Create(m *MateriaPrima) error {
	if err := r.db.Create(m).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *MateriaPrimaRepository) Update(m *MateriaPrima) error {
	if err := r.db.Save(m).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

func (r *MateriaPrimaRepository) Delete(id int) error {
	if err := r.db.Delete(&MateriaPrima{}, id).Error; err != nil {
		return ErrDBInternal
	}
	return nil
}

