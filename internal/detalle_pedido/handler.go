package detalle_pedido

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DetallePedidoHandler struct {
	Service *DetallePedidoService
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
	numPedido, _ := strconv.Atoi(c.Param("numPedido"))
	codProducto, _ := strconv.Atoi(c.Param("codProducto"))

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}
	if err := h.Service.CreateDetalle(&detalle); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, detalle)
}

// PUT /detalle_pedidos/:numPedido/:codProducto
func (h *DetallePedidoHandler) UpdateDetalle(c *gin.Context) {
	var detalle DetallePedido
	if err := c.ShouldBindJSON(&detalle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}
	if err := h.Service.UpdateDetalle(&detalle); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "detalle pedido actualizado exitosamente"})
}

// DELETE /detalle_pedidos/:numPedido/:codProducto
func (h *DetallePedidoHandler) DeleteDetalle(c *gin.Context) {
	numPedido, _ := strconv.Atoi(c.Param("numPedido"))
	codProducto, _ := strconv.Atoi(c.Param("codProducto"))

	if err := h.Service.DeleteDetalle(numPedido, codProducto); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "detalle pedido eliminado exitosamente"})
}
