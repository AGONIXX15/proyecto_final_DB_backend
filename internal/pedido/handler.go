package pedido

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PedidoHandler struct {
	Service *PedidoService
}

func NewPedidoHandler(service *PedidoService) *PedidoHandler {
	return &PedidoHandler{Service: service}
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

// GET /pedidos
func (h *PedidoHandler) GetAllPedidos(c *gin.Context) {
	pedidos, err := h.Service.GetAllPedidos()
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, pedidos)
}

// GET /pedidos/:id
func (h *PedidoHandler) GetPedido(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	pedido, err := h.Service.GetPedido(id)
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, pedido)
}

// POST /pedidos
func (h *PedidoHandler) CreatePedido(c *gin.Context) {
	var pedido Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}
	if err := h.Service.CreatePedido(&pedido); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, pedido)
}

// PUT /pedidos/:id
func (h *PedidoHandler) UpdatePedido(c *gin.Context) {
	var pedido Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}
	if err := h.Service.UpdatePedido(&pedido); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "pedido actualizado exitosamente"})
}

// DELETE /pedidos/:id
func (h *PedidoHandler) DeletePedido(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeletePedido(id); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "pedido eliminado exitosamente"})
}

