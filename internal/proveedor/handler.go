package proveedor

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AGONIXX15/db_proyecto_final/internal/utils"
	"github.com/gin-gonic/gin"
)

type ProveedorHandler struct {
	Service *ProveedorService
}

type ProveedorUpdateDTO struct {
	Nombre         *string `json:"nombre"`
	Direccion      *string `json:"direccion"`
	Telefono       *string `json:"telefono"`
	NombreContacto *string `json:"nombre_contacto"`
}


func NewProveedorHandler(service *ProveedorService) *ProveedorHandler {
	return &ProveedorHandler{Service: service}
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

// GET /proveedores
func (h *ProveedorHandler) GetAllProveedores(c *gin.Context) {
	proveedores, err := h.Service.GetAllProveedores()
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, proveedores)
}

// GET /proveedores/:nit
func (h *ProveedorHandler) GetProveedor(c *gin.Context) {
	nit, _ := strconv.Atoi(c.Param("nit"))

	proveedor, err := h.Service.GetProveedor(nit)
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, proveedor)
}

// POST /proveedores
func (h *ProveedorHandler) CreateProveedor(c *gin.Context) {
	var proveedor Proveedor
	if err := c.ShouldBindJSON(&proveedor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}
	if err := h.Service.CreateProveedor(&proveedor); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, proveedor)
}

// PUT /proveedores/:nit
func (h *ProveedorHandler) UpdateProveedor(c *gin.Context) {
    nit := int(utils.MustParamUint(c, "nit"))

    _, err := h.Service.repo.GetByNIT(nit)
    if err != nil {
        HandleServiceError(c, err)
        return
    }

    // DTO parcial
    var dto ProveedorUpdateDTO
    if err := c.ShouldBindJSON(&dto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
        return
    }

    updates := make(map[string]interface{})

    if dto.Nombre != nil {
        updates["nombre"] = *dto.Nombre
    }
    if dto.Direccion != nil {
        updates["direccion"] = *dto.Direccion
    }
    if dto.Telefono != nil {
        updates["telefono"] = *dto.Telefono
    }
    if dto.NombreContacto != nil {
        updates["nombre_contacto"] = *dto.NombreContacto
    }

    if err := h.Service.UpdateProveedorPartial(nit, updates); err != nil {
        HandleServiceError(c, err)
        return
    }

    updated, _ := h.Service.repo.GetByNIT(nit)

    c.JSON(http.StatusOK, gin.H{
        "message":  "proveedor actualizado exitosamente",
        "proveedor": updated,
    })
}


// DELETE /proveedores/:nit
func (h *ProveedorHandler) DeleteProveedor(c *gin.Context) {
	nit, _ := strconv.Atoi(c.Param("nit"))
	if err := h.Service.DeleteProveedor(nit); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "proveedor eliminado exitosamente"})
}

