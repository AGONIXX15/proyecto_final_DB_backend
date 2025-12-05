package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MustParamUint(c *gin.Context, name string) uint {
    param := c.Param(name)
    id64, err := strconv.ParseUint(param, 10, 64)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error": "parámetro inválido: " + name,
        })
        c.Done()
        return 0
    }
    return uint(id64)
}





