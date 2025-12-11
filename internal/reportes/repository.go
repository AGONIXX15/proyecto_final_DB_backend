package reportes

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Producto struct {
	CodProducto   int     `json:"cod_producto"`
	Nombre        string  `json:"nombre"`
	TipoProducto  string  `json:"tipo_producto"`
	Cantidad      int     `json:"cantidad"`
	Medidas       string  `json:"medidas"`
	Observaciones string  `json:"observacion"`
	PrecioUnidad  float64 `json:"precio_unidad"`
	Subtotal      float64 `json:"subtotal"`
}

type PedidoCliente struct {
	NumPedido    int        `json:"num_pedido"`
	FechaEncargo time.Time  `json:"fecha_encargo"`
	FechaEntrega time.Time  `json:"fecha_entrega"`
	Estado       string     `json:"estado"`
	Total        float64    `json:"total"`
	Productos    []Producto `json:"productos"`
}

// Struct para mapear la fila plana de la vista
type ProductoRow struct {
	NumPedido    int
	FechaEncargo time.Time
	FechaEntrega time.Time
	Estado       string
	CodProducto  int
	Nombre       string
	TipoProducto string
	Cantidad     int
	Medidas      string
	PrecioUnidad float64
}

type ReporteRepository struct {
	db *gorm.DB
}

type Venta struct {
	NumPedido    int     `json:"num_pedido"`
	FechaEncargo string  `json:"fecha_encargo"`
	FechaEntrega string  `json:"fecha_entrega"`
	Estado       string  `json:"estado"`
	Total        float64 `json:"total"`
}

type ColegioVenta struct {
	IDColegio     int     `json:"id_colegio"`
	NombreColegio string  `json:"nombre_colegio"`
	TotalPedidos  int     `json:"total_pedidos"`
	TotalVentas   float64 `json:"total_ventas"`
	Pedidos       []Venta `json:"pedidos"`
}

type UniformeEsencial struct {
    IDUniforme    int    `json:"id_uniforme"`
    TipoUniforme  string `json:"tipo_uniforme"`
    Color         string `json:"color"`
    TipoTela      string `json:"tipo_tela"`
    Bordado       string `json:"bordado"`
    Estampado     string `json:"estampado"`
    Detalles      string `json:"detalles"`
}

type ColegioUniformesEsenciales struct {
    IDColegio     int                `json:"id_colegio"`
    NombreColegio string             `json:"nombre_colegio"`
    Uniformes     []UniformeEsencial `json:"uniformes"`
}

func NewReporteRepository(db *gorm.DB) *ReporteRepository {
	return &ReporteRepository{db}
}

func (r *ReporteRepository) GetProductosPorCliente(documento int) ([]PedidoCliente, error) {
	var rows []ProductoRow

	// Consulta a la vista
	err := r.db.Raw(`
        SELECT 
            num_pedido, fecha_encargo, fecha_entrega, estado,
            cod_producto, nombre, tipo_producto, cantidad, medidas, precio_unidad
        FROM v_reportes_productos_cliente
        WHERE cliente_documento = ?`, documento).Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	// Map para agrupar productos por pedido
	pedidosMap := make(map[int]*PedidoCliente)
	for _, row := range rows {
		if _, exists := pedidosMap[row.NumPedido]; !exists {
			pedidosMap[row.NumPedido] = &PedidoCliente{
				NumPedido:    row.NumPedido,
				FechaEncargo: row.FechaEncargo,
				FechaEntrega: row.FechaEntrega,
				Estado:       row.Estado,
				Total:        0,
				Productos:    []Producto{},
			}
		}

		producto := Producto{
			CodProducto:  row.CodProducto,
			Nombre:       row.Nombre,
			TipoProducto: row.TipoProducto,
			Cantidad:     row.Cantidad,
			Medidas:      row.Medidas,
			PrecioUnidad: row.PrecioUnidad,
			Subtotal:     float64(row.Cantidad) * row.PrecioUnidad,
		}

		pedidosMap[row.NumPedido].Productos = append(pedidosMap[row.NumPedido].Productos, producto)
		pedidosMap[row.NumPedido].Total += producto.Subtotal
	}

	// Convertir map a slice
	var pedidos []PedidoCliente
	for _, p := range pedidosMap {
		pedidos = append(pedidos, *p)
	}

	return pedidos, nil
}

// ------------------- Colegios con Uniformes -------------------
func (r *ReporteRepository) GetColegiosUniformes() ([]ColegioUniforme, error) {
	var rows []struct {
		IDColegio     int
		NombreColegio string
		Direccion     string
		Telefono      string
		CodUniforme   int
		TipoUniforme  string
		Categoria     string
		Detalles      string
	}

	err := r.db.Raw("SELECT * FROM v_reportes_colegios_uniformes").Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	colegiosMap := make(map[int]*ColegioUniforme)
	for _, row := range rows {
		if _, exists := colegiosMap[row.IDColegio]; !exists {
			colegiosMap[row.IDColegio] = &ColegioUniforme{
				IDColegio:     row.IDColegio,
				NombreColegio: row.NombreColegio,
				Direccion:     row.Direccion,
				Telefono:      row.Telefono,
				Uniformes:     []Producto{},
			}
		}

		colegiosMap[row.IDColegio].Uniformes = append(colegiosMap[row.IDColegio].Uniformes, Producto{
			CodProducto:   row.CodUniforme,
			Nombre:        row.TipoUniforme,
			TipoProducto:  row.Categoria,
			Observaciones: row.Detalles,
		})
	}

	var colegios []ColegioUniforme
	for _, c := range colegiosMap {
		colegios = append(colegios, *c)
	}

	return colegios, nil
}

// ------------------- Ventas por Colegio -------------------
func (r *ReporteRepository) GetVentasColegio() ([]ColegioVenta, error) {
	var rows []struct {
		IDColegio     int
		NombreColegio string
		NumPedido     int
		FechaEncargo  string
		FechaEntrega  string
		Estado        string
		TotalPedido   float64
	}

	err := r.db.Raw("SELECT * FROM v_reportes_ventas_colegio").Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	colegiosMap := make(map[int]*ColegioVenta)
	for _, row := range rows {
		if _, exists := colegiosMap[row.IDColegio]; !exists {
			colegiosMap[row.IDColegio] = &ColegioVenta{
				IDColegio:     row.IDColegio,
				NombreColegio: row.NombreColegio,
				TotalPedidos:  0,
				TotalVentas:   0,
				Pedidos:       []Venta{},
			}
		}

		colegiosMap[row.IDColegio].Pedidos = append(colegiosMap[row.IDColegio].Pedidos, Venta{
			NumPedido:    row.NumPedido,
			FechaEncargo: row.FechaEncargo,
			FechaEntrega: row.FechaEntrega,
			Estado:       row.Estado,
			Total:        row.TotalPedido,
		})

		colegiosMap[row.IDColegio].TotalPedidos++
		colegiosMap[row.IDColegio].TotalVentas += row.TotalPedido
	}

	var colegios []ColegioVenta
	for _, c := range colegiosMap {
		colegios = append(colegios, *c)
	}

	return colegios, nil
}

// ------------------- Total Ventas -------------------
func (r *ReporteRepository) GetTotalVentas() (TotalVentas, error) {
	var t TotalVentas
	err := r.db.Raw("SELECT * FROM v_reportes_total_ventas").Scan(&t).Error
	return t, err
}

func (r *ReporteRepository) GetColegiosConUniformes() ([]ColegioUniformesEsenciales, error) {
    // 1. Obtenemos los datos de la vista
    type ColegioRaw struct {
        IDColegio     int             `json:"id_colegio"`
        NombreColegio string          `json:"nombre_colegio"`
        Uniformes     json.RawMessage `json:"uniformes"`
    }

    var colegiosRaw []ColegioRaw
    if err := r.db.Raw("SELECT * FROM colegios_con_uniformes_esenciales").Scan(&colegiosRaw).Error; err != nil {
        return nil, err
    }

    // 2. Convertimos el JSON crudo en []UniformeEsencial
    var colegios []ColegioUniformesEsenciales
    for _, c := range colegiosRaw {
        var uniformes []UniformeEsencial
        if err := json.Unmarshal(c.Uniformes, &uniformes); err != nil {
            return nil, err
        }
        colegios = append(colegios, ColegioUniformesEsenciales{
            IDColegio:     c.IDColegio,
            NombreColegio: c.NombreColegio,
            Uniformes:     uniformes,
        })
    }

    return colegios, nil
}

