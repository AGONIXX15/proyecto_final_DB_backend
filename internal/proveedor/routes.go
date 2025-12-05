package proveedor

import "github.com/gin-gonic/gin"

func RegisterProveedorRoutes(r *gin.Engine, h *ProveedorHandler) {
	router := r.Group("/proveedores")
	{
		router.GET("", h.GetAllProveedores)
		router.GET("/:nit", h.GetProveedor)

		router.POST("", h.CreateProveedor)
		router.PUT("/:nit", h.UpdateProveedor)

		router.DELETE("/:nit", h.DeleteProveedor)
	}
}
