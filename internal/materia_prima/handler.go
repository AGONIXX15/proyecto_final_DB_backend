package materia_prima

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MateriaPrimaHandler struct {
	Service *MateriaPrimaService
}

func NewMateriaPrimaHandler(service *MateriaPrimaService) *MateriaPrimaHandler {
	return &MateriaPrimaHandler{Service: service}
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

// GET /materias_primas
func (h *MateriaPrimaHandler) GetAllMaterias(c *gin.Context) {
	materias, err := h.Service.GetAllMaterias()
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, materias)
}

// GET /materias_primas/:id
func (h *MateriaPrimaHandler) GetMateria(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	materia, err := h.Service.GetMateria(id)
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, materia)
}

// POST /materias_primas
func (h *MateriaPrimaHandler) CreateMateria(c *gin.Context) {
	var materia MateriaPrima
	if err := c.ShouldBindJSON(&materia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}
	if err := h.Service.CreateMateria(&materia); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, materia)
}

// PUT /materias_primas/:id
func (h *MateriaPrimaHandler) UpdateMateria(c *gin.Context) {
	var materia MateriaPrima
	if err := c.ShouldBindJSON(&materia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}
	if err := h.Service.UpdateMateria(&materia); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "materia prima actualizada exitosamente"})
}

// DELETE /materias_primas/:id
func (h *MateriaPrimaHandler) DeleteMateria(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteMateria(id); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "materia prima eliminada exitosamente"})
}

