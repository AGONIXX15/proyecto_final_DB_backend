package factura

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/AGONIXX15/db_proyecto_final/internal/utils"
	"github.com/gin-gonic/gin"
)

type FacturaHandler struct {
	Service *FacturaService
}

type UpdateFacturaDTO struct {
    Fecha     *string  `json:"fecha"`       // string ISO (ej: "2025-12-10T15:04:05Z")
    Total     *float64 `json:"total"`
    NumPedido *int     `json:"num_pedido"`
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
    numFactura := utils.MustParamUint(c, "numFactura")

    var dto UpdateFacturaDTO
    if err := c.ShouldBindJSON(&dto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
        return
    }

    updates := make(map[string]interface{})

    if dto.Fecha != nil {
        // Parseamos la fecha ISO si llega
        t, err := time.Parse(time.RFC3339, *dto.Fecha)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "fecha en formato inválido, usa ISO 8601"})
            return
        }
        updates["fecha"] = t
    }

    if dto.Total != nil {
        updates["total"] = *dto.Total
    }

    if dto.NumPedido != nil {
        updates["num_pedido"] = *dto.NumPedido
    }

    if len(updates) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ningun campo para actualizar"})
        return
    }

    if err := h.Service.UpdateFacturaPartial(int(numFactura), updates); err != nil {
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

