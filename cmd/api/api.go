package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mutinho/middleware"
	"github.com/mutinho/src/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORS())

	api := r.Group("/api")
	{
		//Rutas
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})

		})

		auth := api.Group("/auth")
		{
			auth.POST("/register", handler.Register)
			auth.POST("/login", handler.Login)
		}

	}

	//Rutas Protegidas
	protected := r.Group("/api")
	protected.Use(middleware.Auth())
	{
		// Ejemplo de ruta protegida
		protected.GET("/profile", func(c *gin.Context) {
			userID := c.GetUint("userID")
			c.JSON(200, gin.H{
				"user_id": userID,
				"message": "This is a protected route",
			})
		})

	}

	return r

}
