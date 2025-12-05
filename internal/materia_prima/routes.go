package materia_prima

import "github.com/gin-gonic/gin"

func RegisterMateriaPrimaRoutes(r *gin.Engine, h *MateriaPrimaHandler) {
	router := r.Group("/materias-primas")
	{
		router.GET("", h.GetAllMaterias)
		router.GET("/:id", h.GetMateria)

		router.POST("", h.CreateMateria)

		router.PUT("/:id", h.UpdateMateria)

		router.DELETE("/:id", h.DeleteMateria)
	}
}
