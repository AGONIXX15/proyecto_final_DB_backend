package producto_terminado

import "github.com/gin-gonic/gin"

func RegisterProductoTerminadoRoutes(r *gin.Engine, h *ProductoTerminadoHandler) {
	router := r.Group("/producto-terminado")
	{
		router.POST("", h.Create)
		router.GET("", h.GetAll)

		router.GET("/:id", h.GetByID)
		router.PUT("/:id", h.Update)
		router.PATCH("/:id", h.Update)

		router.DELETE("/:id", h.Delete)
	}
}
