package reportes

import (
    "github.com/gin-gonic/gin"
)

func RegisterRoutesReportes(router *gin.Engine, h *ReporteHandler) {

    r := router.Group("/reportes")
    {
        // r.GET("/productos-pendientes", h.PedidosPendientes)
        r.GET("/productos-cliente/:documento", h.ProductosPorCliente)
        r.GET("/colegios-uniformes", h.ColegiosUniformes)
        r.GET("/ventas-colegio", h.VentasColegio)
        r.GET("/total-ventas", h.TotalVentas)
    }
}

