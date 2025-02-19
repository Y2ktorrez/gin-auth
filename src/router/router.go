package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mutinho/middleware"
	"github.com/mutinho/src/router/endpoints"
)

func SetupRouter(r *gin.Engine) {

	api := r.Group("/api")

	//Las Demas Rutas
	endpoints.AuthRouter(api)

	/*Rutas Protegidas*/
	protected := api.Group("")
	protected.Use(middleware.Auth())

	//Las Demas Rutas Protegidas
	endpoints.ProtectedRoutes(protected)

}
