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

type UpdateDetalleDTO struct {
    Cantidad      *int    `json:"cantidad"`
    Medidas       *string `json:"medidas"`
    Observaciones *string `json:"observaciones"`
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

// GET /detalle_pedidos/:numPedido/:codProducto
func (h *DetallePedidoHandler) GetDetalle(c *gin.Context) {
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

	detalle, err := h.Service.GetDetalle(numPedido, codProducto)
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
    numPedido := utils.MustParamUint(c, "numPedido")
    typeItem := c.Param("typeItem")
    codItem := utils.MustParamUint(c, "codItem")

    var dto UpdateDetalleDTO
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

    if len(updates) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ningun campo para actualizar"})
        return
    }

    if err := h.Service.UpdateDetallePartial(int(numPedido), typeItem, int(codItem), updates); err != nil {
        HandleServiceError(c, err)
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "detalle pedido actualizado exitosamente"})
}


// DELETE /detalle_pedidos/:numPedido/:codProducto
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

