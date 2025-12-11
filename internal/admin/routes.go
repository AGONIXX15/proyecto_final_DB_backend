package admin

import "github.com/gin-gonic/gin"

func RegisterAdminRoutes(r *gin.Engine, h *AdminHandler) {
	router := r.Group("/admin")
	{
		router.GET("", h.GetAllAdmins)

		router.POST("/login",h.LoginAdmin)
	
		router.GET("/:id",h.GetAdminByID)

		router.POST("",h.CreateAdmin)

		router.PATCH("/:id",h.UpdateAdmin)
		router.PUT("/:id", h.UpdateAdmin)
		router.DELETE("/:id", h.DeleteAdmin)
	}
}
