package producto_terminado

import (
	"fmt"
	"net/http"

	"github.com/AGONIXX15/db_proyecto_final/internal/utils"
	"github.com/gin-gonic/gin"
)

type ProductoTerminadoHandler struct {
    service ProductoTerminadoService
}


type ProductoTerminadoUpdateDTO struct {
	Nombre             *string  `json:"nombre"`
	Categoria          *string  `json:"categoria"`
	Descripcion        *string  `json:"descripcion"`
	Talla              *string  `json:"talla"`
	Sexo               *string  `json:"sexo"`
	Precio             *float64 `json:"precio"`
	CantidadExistencia *int     `json:"cantidad_existencia"`
}


func NewProductoTerminadoHandler(service ProductoTerminadoService) *ProductoTerminadoHandler {
    return &ProductoTerminadoHandler{service}
}

func (h *ProductoTerminadoHandler) Create(c *gin.Context) {
    var producto ProductoTerminado

    if err := c.ShouldBindJSON(&producto); err != nil {
			fmt.Println("fallo haciendo el binding")
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }	

    if err := h.service.Create(&producto); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, producto)
}

func (h *ProductoTerminadoHandler) GetAll(c *gin.Context) {
    productos, err := h.service.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, productos)
}

func (h *ProductoTerminadoHandler) GetByID(c *gin.Context) {
    id := utils.MustParamUint(c,"id")
    producto, err := h.service.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
        return
    }

    c.JSON(http.StatusOK, producto)
}

func (h *ProductoTerminadoHandler) Update(c *gin.Context) {
    id := utils.MustParamUint(c, "id")

    _, err := h.service.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
        return
    }

    var dto ProductoTerminadoUpdateDTO
    if err := c.ShouldBindJSON(&dto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
        return
    }

    updates := make(map[string]interface{})

    if dto.Nombre != nil {
        updates["nombre"] = *dto.Nombre
    }
    if dto.Categoria != nil {
        updates["categoria"] = *dto.Categoria
    }
    if dto.Descripcion != nil {
        updates["descripcion"] = *dto.Descripcion
    }
    if dto.Talla != nil {
        updates["talla"] = *dto.Talla
    }
    if dto.Sexo != nil {
        updates["sexo"] = *dto.Sexo
    }
    if dto.Precio != nil {
        updates["precio"] = *dto.Precio
    }
    if dto.CantidadExistencia != nil {
        updates["cantidad_existencia"] = *dto.CantidadExistencia
    }

    if err := h.service.UpdatePartial(id, updates); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    updated, _ := h.service.GetByID(id)

    c.JSON(http.StatusOK, updated)
}


func (h *ProductoTerminadoHandler) Delete(c *gin.Context) {
    id := utils.MustParamUint(c,"id")

    if err := h.service.Delete(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
}

