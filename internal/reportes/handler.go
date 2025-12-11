package reportes

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReporteHandler struct {
    service *ReporteService
}

func NewReporteHandler(service *ReporteService) *ReporteHandler {
    return &ReporteHandler{service: service}
}

// func (h *ReporteHandler) PedidosPendientes(c *gin.Context) {
//     pedidos, err := h.service.PedidosPendientes()
//     if err != nil {
//         c.JSON(500, gin.H{"error": err.Error()})
//         return
//     }
//     c.JSON(200, pedidos)
// }

func (h *ReporteHandler) ProductosPorCliente(c *gin.Context) {
    documento := c.Param("documento")
    pedidos, err := h.service.ProductosPorCliente(stringToInt(documento))
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, pedidos)
}

func (h *ReporteHandler) ColegiosUniformes(c *gin.Context) {
    colegios, err := h.service.repo.GetColegiosConUniformes()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, colegios)
}

func (h *ReporteHandler) VentasColegio(c *gin.Context) {
    ventas, err := h.service.VentasColegio()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, ventas)
}

func (h *ReporteHandler) TotalVentas(c *gin.Context) {
    t, err := h.service.TotalVentas()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, t)
}

func stringToInt(s string) int {
    // simple helper para convertir string a int
    i, _ := strconv.Atoi(s)
    return i
}

