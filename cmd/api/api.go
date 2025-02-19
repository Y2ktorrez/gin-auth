package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mutinho/middleware"
	"github.com/mutinho/src/router"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.CORS())
	router.SetupRouter(r)
	return r

}
