package pedido

import "github.com/gin-gonic/gin"

func RegisterPedidoRoutes(r *gin.Engine, h *PedidoHandler) {
	router := r.Group("/pedidos")
	{

	router.GET("", h.GetAllPedidos)
	router.GET("/:id", h.GetPedido)

	router.POST("", h.CreatePedido)

	router.PUT("/:id", h.UpdatePedido)
	router.PUT("/:id/entregar", h.DeliverPedido)

	router.DELETE("/:id", h.DeletePedido)
	}
}

