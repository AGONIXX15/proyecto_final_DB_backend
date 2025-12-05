package admin

import (
	"errors"
	"net/http"

	"github.com/AGONIXX15/db_proyecto_final/internal/utils"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	service *AdminService
}

func HandleServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, ErrDBInternal):
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	case errors.Is(err, ErrPasswordWrong):
		c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error desconocido"})
	default:
		// No hace nada el caller se encarga
	}
}

// POST /admin
func (h *AdminHandler) CreateAdmin(c *gin.Context) {
	var admin Admin
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos invalidos"})
		return
	}

	err := h.service.CreateAdmin(&admin)
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusCreated, admin)
}

// GET /admin
func (h *AdminHandler) GetAllAdmins(c *gin.Context) {
	admins, err := h.service.GetAllAdmins()
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, admins)
}

// GET /admin/:id
func (h *AdminHandler) GetAdminByID(c *gin.Context) {
	id := utils.MustParamUint(c, "id")
	admin, err := h.service.GetAdmin(id)
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.JSON(http.StatusOK, admin)
}

// PATCH /admin/:id || UPDATE /admin/:id
func (h *AdminHandler) UpdateAdmin(c *gin.Context) {
	id := utils.MustParamUint(c, "id")
	var admin Admin
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	admin.ID = id
	err := h.service.UpdateAdmin(&admin)
	if err != nil {
		HandleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, admin)
}

// DELETE /admin/:id
func (h *AdminHandler) DeleteAdmin(c *gin.Context) {
	id := utils.MustParamUint(c, "id")
	err := h.service.DeleteAdmin(id)
	if err != nil {
		HandleServiceError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// POST /admins/login
func (h *AdminHandler) LoginAdmin(c *gin.Context) {
	var admin Admin
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.LoginAdmin(admin.Username, admin.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
		return
	}
	
	// enviar jwts
	c.JSON(http.StatusOK, gin.H{"message": "Login exitoso"})
}
