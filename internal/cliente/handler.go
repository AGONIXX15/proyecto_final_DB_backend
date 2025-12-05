package cliente

import (
	"errors"
	"net/http"

	"github.com/AGONIXX15/db_proyecto_final/internal/utils"
	"github.com/gin-gonic/gin"
)

type ClienteHandler struct {
	Service *ClienteService
}

func NewClienteHandler(service *ClienteService) *ClienteHandler {
	return &ClienteHandler{Service: service}
}

// GET /cliente
func (h *ClienteHandler) GetAllClientes(c *gin.Context) {
	clientes, err := h.Service.GetAllClientes()
	if err != nil {
		HandleServiceError(c,err)
		return
	}
	c.JSON(http.StatusOK, clientes)
}

// GET /cliente/:id
func (h *ClienteHandler) GetCliente(c *gin.Context) {
	documento := utils.MustParamUint(c, "documento")

	cliente, err := h.Service.GetCliente(documento)
	if err != nil {
		HandleServiceError(c,err)
		return
	}
	
	c.JSON(http.StatusOK, cliente)
}

// POST /cliente
func (h *ClienteHandler) CreateCliente(c *gin.Context) {
	var cliente Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
	}
	err := h.Service.CreateCliente(&cliente)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cliente)
}

func HandleServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, ErrDBInternal):
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error desconocido"})
	default:
		// No hace nada el caller se encarga
	}
}


// DELETE /clientes/:documento
func (h *ClienteHandler) DeleteCliente(c *gin.Context) {
	documento := utils.MustParamUint(c, "documento")
	err := h.Service.DeleteCliente(documento)
	if err != nil {
		HandleServiceError(c,err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cliente eliminado exitosamente"})
}
// UPDATE /clientes/:documento || PATCH /clientes/:documento
func (h *ClienteHandler) UpdateCliente(c *gin.Context) {
	var cliente Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"datos de forma invalida"})
		return 
	}

	err := h.Service.UpdateCliente(&cliente)
	if err != nil {
		HandleServiceError(c,err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"cliente actualizado exitosamente"})
}
