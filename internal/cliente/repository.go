package cliente

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var (
	ErrNotFound   = errors.New("cliente no encontrado")
	ErrDBInternal = errors.New("error interno de la base de datos")
)

type ClienteRepository interface {
	FindByID(id uint) (*Cliente, error)
	FindAll() ([]Cliente, error)
	Create(cliente *Cliente) error
	Update(cliente *Cliente) error
	Delete(id uint) error
	UpdatePartial(id uint, updates map[string]interface{}) error
}

type clienteRepository struct {
	db *gorm.DB
}

func NewClienteRepository(db *gorm.DB) ClienteRepository {
	return &clienteRepository{db: db}
}

func (r *clienteRepository) Create(cliente *Cliente) error {
	result := r.db.Create(cliente)
	if result.Error != nil {
		return fmt.Errorf("%w: no se pudo crear el cliente con documento %d: %s", ErrDBInternal, cliente.Documento, result.Error)
	}
	return nil
}

func (c *clienteRepository) Delete(documento uint) error {
	result := c.db.Delete(&Cliente{}, documento)

	if result.Error != nil {
		return fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("%w: cliente con documento %d", ErrNotFound, documento)
	}

	return nil
}

// FindAll obtiene todos los clientes
func (c *clienteRepository) FindAll() ([]Cliente, error) {
	var clientes []Cliente
	result := c.db.Find(&clientes)
	if result.Error != nil {
		return nil, fmt.Errorf("%w: error al obtener todos los clientes: %s", ErrDBInternal, result.Error)
	}
	return clientes, nil
}

// FindByID obtiene un cliente por documento
func (c *clienteRepository) FindByID(documento uint) (*Cliente, error) {
	var cliente Cliente
	result := c.db.First(&cliente, documento)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("%w: cliente con documento %d", ErrNotFound, documento)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
	}

	return &cliente, nil
}

// Update actualiza un cliente
func (c *clienteRepository) Update(cliente *Cliente) error {
	result := c.db.Model(&Cliente{}).Where("documento = ?", cliente.Documento).Updates(cliente)

	if result.Error != nil {
		return fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("%w: cliente con documento %d", ErrNotFound, cliente.Documento)
	}

	return nil
}

func (r *clienteRepository) UpdatePartial(id uint, updates map[string]interface{}) error {
    result := r.db.Model(&Cliente{}).
        Where("documento = ?", id).
        Updates(updates)

    if result.Error != nil {
        return fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
    }

    if result.RowsAffected == 0 {
        return fmt.Errorf("%w: cliente con documento %d no encontrado", ErrNotFound, id)
    }

    return nil
}


