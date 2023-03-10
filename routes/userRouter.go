package routes

import (
	controller "github.com/KiranDesh/ecommece/controllers"
	"github.com/KiranDesh/ecommece/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
