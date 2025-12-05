package factura

import "github.com/gin-gonic/gin"

func RegisterFacturaRoutes(r *gin.Engine, h *FacturaHandler) {
	r.GET("/facturas", h.GetAllFacturas)
	r.GET("/facturas/:id", h.GetFactura)
	r.POST("/facturas", h.CreateFactura)
	r.PUT("/facturas/:id", h.UpdateFactura)
	r.DELETE("/facturas/:id", h.DeleteFactura)
}

