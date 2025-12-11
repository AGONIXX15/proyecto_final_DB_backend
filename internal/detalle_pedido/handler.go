package detalle_pedido

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AGONIXX15/db_proyecto_final/internal/utils"
	"github.com/gin-gonic/gin"
)

type DetallePedidoHandler struct {
	Service *DetallePedidoService
}

type DetallePedidoUpdateDTO struct {
	Cantidad       *int     `json:"cantidad"`
	Medidas        *string  `json:"medidas"`
	Observaciones  *string  `json:"observaciones"`
	PrecioUnitario *float64 `json:"precio_unitario"`
	SubTotal       *float64 `json:"subtotal"`
}



func NewDetallePedidoHandler(service *DetallePedidoService) *DetallePedidoHandler {
	return &DetallePedidoHandler{Service: service}
}

func HandleServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, ErrDBInternal):
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error desconocido"})
	}
}

// GET /detalle_pedidos
func (h *DetallePedidoHandler) GetAllDetalles(c *gin.Context) {
	detalles, err := h.Service.GetAllDetalles()
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, detalles)
}

// GET /detalle_pedidos/:numPedido/:typeItem/:codProducto
func (h *DetallePedidoHandler) GetDetalle(c *gin.Context) {
	numPedido := int(utils.MustParamUint(c, "numPedido"))
	typeItem := c.Param("typeItem")
	codProducto := utils.MustParamUint(c,"codProducto")

	detalle, err := h.Service.GetDetalle(numPedido, typeItem, codProducto)
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, detalle)
}

// POST /detalle_pedidos
func (h *DetallePedidoHandler) CreateDetalle(c *gin.Context) {
	var detalle DetallePedido
	if err := c.ShouldBindJSON(&detalle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma inválida"})
		return
	}
	if err := h.Service.CreateDetalle(&detalle); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusCreated, detalle)
}

// PUT /detalle_pedidos/:numPedido/:codProducto
func (h *DetallePedidoHandler) UpdateDetalle(c *gin.Context) {
    numPedido := int(utils.MustParamUint(c, "num_pedido"))
    typeItem := c.Param("type_item")
    codItem := utils.MustParamUint(c, "cod_item")

    // Verificar existencia
    _, err := h.Service.GetDetalle(numPedido, typeItem, codItem)
    if err != nil {
        HandleServiceError(c, err)
        return
    }

    // DTO parcial
    var dto DetallePedidoUpdateDTO
    if err := c.ShouldBindJSON(&dto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma inválida"})
        return
    }

    updates := make(map[string]interface{})

    if dto.Cantidad != nil {
        updates["cantidad"] = *dto.Cantidad
    }
    if dto.Medidas != nil {
        updates["medidas"] = *dto.Medidas
    }
    if dto.Observaciones != nil {
        updates["observaciones"] = *dto.Observaciones
    }
    if dto.PrecioUnitario != nil {
        updates["precio_unitario"] = *dto.PrecioUnitario
    }
    if dto.SubTotal != nil {
        updates["subtotal"] = *dto.SubTotal
    }

    if err := h.Service.UpdateDetallePartial(numPedido, typeItem, codItem, updates); err != nil {
        HandleServiceError(c, err)
        return
    }

    actualizado, _ := h.Service.GetDetalle(numPedido, typeItem, codItem)

    c.JSON(http.StatusOK, gin.H{
        "message":  "detalle pedido actualizado exitosamente",
        "detalle": actualizado,
    })
}



// DELETE /detalle_pedidos/:numPedido/:typePedido/:codProducto
func (h *DetallePedidoHandler) DeleteDetalle(c *gin.Context) {
	numPedido, err := strconv.Atoi(c.Param("numPedido"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "numPedido inválido"})
		return
	}
	codProducto, err := strconv.Atoi(c.Param("codProducto"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "codProducto inválido"})
		return
	}

	if err := h.Service.DeleteDetalle(numPedido, codProducto); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "detalle pedido eliminado exitosamente"})
}

