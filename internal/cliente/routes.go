package cliente

import "github.com/gin-gonic/gin"


func RegisterClienteRoutes(r *gin.Engine, h *ClienteHandler) {
	router := r.Group("/cliente")
	{
		router.GET("",h.GetAllClientes)
		router.POST("",h.CreateCliente)

		router.GET("/:id",h.GetCliente)
		router.DELETE("/:id",h.DeleteCliente)
		router.PATCH("/:id", h.UpdateCliente)
		router.PUT("/:id", h.UpdateCliente)
	}
}
