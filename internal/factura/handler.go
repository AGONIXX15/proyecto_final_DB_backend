package factura

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FacturaHandler struct {
	Service *FacturaService
}

func NewFacturaHandler(service *FacturaService) *FacturaHandler {
	return &FacturaHandler{Service: service}
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

// GET /facturas
func (h *FacturaHandler) GetAllFacturas(c *gin.Context) {
	facturas, err := h.Service.GetAllFacturas()
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, facturas)
}

// GET /facturas/:id
func (h *FacturaHandler) GetFactura(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	factura, err := h.Service.GetFactura(id)
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, factura)
}

// POST /facturas
func (h *FacturaHandler) CreateFactura(c *gin.Context) {
	var factura Factura
	if err := c.ShouldBindJSON(&factura); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}
	if err := h.Service.CreateFactura(&factura); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, factura)
}

// PUT /facturas/:id
func (h *FacturaHandler) UpdateFactura(c *gin.Context) {
	var factura Factura
	if err := c.ShouldBindJSON(&factura); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}
	if err := h.Service.UpdateFactura(&factura); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "factura actualizada exitosamente"})
}

// DELETE /facturas/:id
func (h *FacturaHandler) DeleteFactura(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteFactura(id); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "factura eliminada exitosamente"})
}

