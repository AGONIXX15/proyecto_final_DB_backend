package reportes

import "github.com/gin-gonic/gin"


func RegisterRoutesReportes(r *gin.Engine, h *ReporteHandler) {
	v := r.Group("/reportes")
{
    v.GET("/productos-pendientes", h.ProductosPendientes)
    v.GET("/productos-cliente/:documento", h.ProductosPorCliente)
    v.GET("/productos-disponibles", h.ProductosDisponibles)
    v.GET("/colegios-uniformes", h.ColegiosConUniformes)
    v.GET("/uniformes-colegio", h.UniformesPorColegio)
    v.GET("/ventas-colegio", h.VentasPorColegio)
    v.GET("/total-ventas", h.TotalVentas)
}
}
