package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

// le chemin d'accès qu'on a définis pour accéder aux controlleurs
func UserRoute(router *gin.Engine) {
	router.GET("/", controllers.UserController)
}
