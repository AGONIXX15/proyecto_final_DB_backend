package detalle_pedido

import "github.com/gin-gonic/gin"

func RegisterDetallePedidoRoutes(r *gin.Engine, h *DetallePedidoHandler) {
	router := r.Group("/detalle_pedidos")
	{
		router.GET("", h.GetAllDetalles)
		router.GET("/:numPedido/:typePedido/:codProducto", h.GetDetalle)

		router.POST("", h.CreateDetalle)
		router.PUT("/:numPedido/:typePedido/:codProducto", h.UpdateDetalle)

		router.DELETE("/:numPedido/:codProducto", h.DeleteDetalle)
	}
}
