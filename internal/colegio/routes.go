package colegio

import "github.com/gin-gonic/gin"

func RegisterColegioRoutes(r *gin.Engine, h *ColegioHandler) {
	router := r.Group("/colegios")
	{
	router.POST("", h.CreateColegio)
	router.GET("", h.GetAllColegios)

	router.GET("/:id", h.GetColegio)
	router.PUT("/:id", h.UpdateColegio)
	router.DELETE("/:id", h.DeleteColegio)
	}
}

