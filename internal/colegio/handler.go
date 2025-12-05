package colegio

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ColegioHandler struct {
	Service *ColegioService
}

func NewColegioHandler(service *ColegioService) *ColegioHandler {
	return &ColegioHandler{Service: service}
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

// GET /colegios
func (h *ColegioHandler) GetAllColegios(c *gin.Context) {
	colegios, err := h.Service.GetAllColegios()
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, colegios)
}

// GET /colegios/:id
func (h *ColegioHandler) GetColegio(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	colegio, err := h.Service.GetColegio(id)
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, colegio)
}

// POST /colegios
func (h *ColegioHandler) CreateColegio(c *gin.Context) {
	var colegio Colegio
	if err := c.ShouldBindJSON(&colegio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}
	if err := h.Service.CreateColegio(&colegio); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, colegio)
}

// PUT /colegios/:id
func (h *ColegioHandler) UpdateColegio(c *gin.Context) {
	var colegio Colegio
	if err := c.ShouldBindJSON(&colegio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos de forma invalida"})
		return
	}
	if err := h.Service.UpdateColegio(&colegio); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "colegio actualizado exitosamente"})
}

// DELETE /colegios/:id
func (h *ColegioHandler) DeleteColegio(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteColegio(id); err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "colegio eliminado exitosamente"})
}

