package uniforme

import "github.com/gin-gonic/gin"

func RegisterUniformeRoutes(r *gin.Engine, h *UniformeHandler) {
	router := r.Group("/uniformes")
	{
		router.GET("", h.GetAllUniformes)
		router.GET("/:id", h.GetUniforme)

		router.POST("", h.CreateUniforme)
		router.PUT("/:id", h.UpdateUniforme)
		router.DELETE("/:id", h.DeleteUniforme)
	}
}
