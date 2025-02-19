package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mutinho/src/service"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := service.RegisterUser(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "User registered successfully",
		"id":       user.ID,
		"username": user.Name,
		"email":    user.Email,
	})

	// Crear un mapa en el orden específico
	/*response := map[string]interface{}{
		"message":  "User registered successfully",
		"id":       user.ID,
		"username": user.Name,
		"email":    user.Email,
	}
	// Enviar respuesta con el orden correcto
	c.JSON(http.StatusCreated, response)*/

}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := service.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"expires": 24 * 3600,
	})
}
