package admin

import (
	"errors"
	"net/http"
	"fmt"

	"github.com/AGONIXX15/db_proyecto_final/internal/utils"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	service *AdminService
}

func NewAdminHandler(service *AdminService) *AdminHandler {
	return &AdminHandler{service}

}

type UpdateAdminDTO struct {
    Username *string `json:"username"`
    Password *string `json:"password"`
    Role     *string `json:"role"`
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

type AdminDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role string `json:"role"`
}

// POST /admin
func (h *AdminHandler) CreateAdmin(c *gin.Context) {
	var json AdminDTO
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos invalidos"})
		return
	}

	admin := Admin{Username:json.Username, Password: json.Password, Role: json.Role}
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

    var dto UpdateAdminDTO
    if err := c.BindJSON(&dto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
        return
    }


	fmt.Println("llego al handler")

    // Crear map dinámico con solo los campos que sí llegaron
    updates := make(map[string]interface{})

    if dto.Username != nil {
        updates["username"] = *dto.Username
    }
    if dto.Password != nil {
        updates["password"] = *dto.Password
    }
    if dto.Role != nil {
        updates["role"] = *dto.Role
    }

    // Si no hay cambios:
    if len(updates) == 0 {
		fmt.Println("tiro error en el handler")
        c.JSON(http.StatusBadRequest, gin.H{"error": "No se enviaron campos para actualizar"})
        return
    }

    err := h.service.UpdateAdminPartial(id, updates)
    if err != nil {
		fmt.Println("tiro error en el handler")
        HandleServiceError(c, err)
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Admin actualizado"})
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
	var json AdminDTO
	fmt.Println("login admin")
		if err := c.BindJSON(&json); err != nil {
		fmt.Println("fail bindingjson")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err := h.service.LoginAdmin(json.Username, json.Password)
	if err != nil {
		fmt.Println("fallo en login admin")
		c.JSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
		return
	}
	
	// enviar jwts
	jwtToken, err := utils.GenerateJWT(int(admin.ID), admin.Username, admin.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login exitoso","token":jwtToken})
}
