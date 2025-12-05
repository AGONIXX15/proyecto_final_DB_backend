package uniforme

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UniformeHandler struct {
	Service *UniformeService
}

func NewUniformeHandler(service *UniformeService) *UniformeHandler {
	return &UniformeHandler{Service: service}
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

// GET /uniformes
func (h *UniformeHandler) GetAllUniformes(c *gin.Context) {
	uniformes, err := h.Service.GetAllUniformes()
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, uniformes)
}

// GET /uniformes/:id
func (h *UniformeHandler) GetUniforme(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	uniforme, err := h.Service.GetUniforme(id)
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, uniforme)
}

// POST /uniformes
func (h *UniformeHandler) CreateUniforme(c *gin.Context) {
	var uniforme Uniforme
	if err := c.ShouldBindJSON(&uniforme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}
	if err := h.Service.CreateUniforme(&uniforme); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, uniforme)
}

// PUT /uniformes/:id
func (h *UniformeHandler) UpdateUniforme(c *gin.Context) {
	var uniforme Uniforme
	if err := c.ShouldBindJSON(&uniforme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}
	if err := h.Service.UpdateUniforme(&uniforme); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "uniforme actualizado exitosamente"})
}

// DELETE /uniformes/:id
func (h *UniformeHandler) DeleteUniforme(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteUniforme(id); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "uniforme eliminado exitosamente"})
}

