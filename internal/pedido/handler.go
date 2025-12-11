package pedido

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AGONIXX15/db_proyecto_final/internal/detalle_pedido"
	"github.com/AGONIXX15/db_proyecto_final/internal/utils"
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

	pedido, err := h.Service.GetPedidoByID(id)
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, pedido)
}

// DTO para recibir JSON de creación/actualización
type PedidoRequest struct {
	Pedido  Pedido                     `json:"pedido"`
	Detalles []detalle_pedido.DetallePedido `json:"detalles"`
}

// POST /pedidos
func (h *PedidoHandler) CreatePedido(c *gin.Context) {
	var req PedidoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("datos invalidos")
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}

	if err := h.Service.CreatePedido(&req.Pedido, req.Detalles); err != nil {
		fmt.Println("no se pudo crear")
		HandleServiceError(c, err)
		return
	}
	fmt.Println("se pudo crear")
	c.JSON(http.StatusCreated, req.Pedido)
}

// PUT /pedidos/:id
func (h *PedidoHandler) UpdatePedido(c *gin.Context) {
	id := utils.MustParamUint(c,"id")
	var req PedidoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("datos actualizacion pedido invalidos")
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}
	if err := h.Service.UpdatePedido(int(id), &req.Pedido, req.Detalles); err != nil {
		fmt.Println("error realizando la actualizacion")
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

// /pedidos/:id/entregar
func (h *PedidoHandler) DeliverPedido(c *gin.Context) {
	id := utils.MustParamUint(c,"id")

	err := h.Service.repo.deliverPedidoFunction(int(id))
	if err != nil {
		fmt.Println("error en deliver ",err)
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message":"se ha hecho entrega"})
}

