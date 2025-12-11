package reportes

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ReporteHandler struct {
    Service *ReporteService
}

func NewReporteHandler(s *ReporteService) *ReporteHandler {
    return &ReporteHandler{Service: s}
}

func (h *ReporteHandler) ProductosPendientes(c *gin.Context) {
    data, err := h.Service.ProductosPendientes()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, data)
}

func (h *ReporteHandler) ProductosPorCliente(c *gin.Context) {
    docStr := c.Param("documento")
    doc, err := strconv.Atoi(docStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "documento inválido"})
        return
    }
    data, err := h.Service.ProductosPorCliente(doc)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, data)
}

func (h *ReporteHandler) ProductosDisponibles(c *gin.Context) {
    data, err := h.Service.ProductosDisponibles()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, data)
}

func (h *ReporteHandler) ColegiosConUniformes(c *gin.Context) {
    data, err := h.Service.ColegiosConUniformes()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, data)
}

func (h *ReporteHandler) UniformesPorColegio(c *gin.Context) {
    data, err := h.Service.UniformesPorColegio()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, data)
}

func (h *ReporteHandler) VentasPorColegio(c *gin.Context) {
    data, err := h.Service.VentasPorColegio()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, data)
}

func (h *ReporteHandler) TotalVentas(c *gin.Context) {
    total, err := h.Service.TotalVentas()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"total": total})
}

