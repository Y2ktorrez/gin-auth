package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/mutinho/src/handler"
)

func AuthRouter(r *gin.RouterGroup) {

	auth := r.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}

}

func ProtectedRoutes(r *gin.RouterGroup) {
	r.GET("/profile", func(c *gin.Context) {
		userID := c.GetUint("userID")
		c.JSON(200, gin.H{
			"user_id": userID,
			"message": "This is a protected route",
		})
	})
}
