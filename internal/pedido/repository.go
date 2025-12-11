package pedido

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"github.com/AGONIXX15/db_proyecto_final/internal/detalle_pedido"
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

// CRUD estándar
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

// Métodos específicos que usan funciones SQL
func (r *PedidoRepository) CreateWithFunction(pedido *Pedido, detalles []detalle_pedido.DetallePedido) error {
	detallesJSON, err := json.Marshal(detalles)
	if err != nil {
		fmt.Println("error serializando datos")
		return fmt.Errorf("error serializando detalles: %w", err)
	}

	fmt.Println("ejecutando la vuelta")
	result := r.db.Exec(
		"SELECT crear_pedido($1::bigint, $2::date, $3::date, $4::numeric, $5::json)",
		pedido.DocCliente,
		pedido.FechaEncargo,
		pedido.FechaEntrega,
		pedido.Abono,
		string(detallesJSON),
	)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return result.Error
}

func (r *PedidoRepository) UpdateWithFunction(numPedido int, updatedPedido *Pedido, nuevosDetalles []detalle_pedido.DetallePedido) error {
	detallesJSON, err := json.Marshal(nuevosDetalles)
	if err != nil {
		return fmt.Errorf("error serializando detalles: %w", err)
	}

	query := `
        SELECT actualizar_pedido(
	$1::int, $2::date, $3::numeric, $4::json
        )
    `
	result := r.db.Exec(
		query,
		numPedido,
		updatedPedido.FechaEntrega,
		updatedPedido.Abono,
		detallesJSON,
	)

	if result.Error != nil {
		fmt.Println(result.Error)
	}


	return result.Error
}

func (r *PedidoRepository) CancelWithFunction(numPedido int) error {
	return r.db.Exec("SELECT cancelar_pedido($1::int)", numPedido).Error
}
func (r *PedidoRepository) deliverPedidoFunction(numPedido int) error {
	return r.db.Exec("SELECT entregar_pedido($1::int)", numPedido).Error

}

