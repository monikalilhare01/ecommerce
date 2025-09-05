package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/monikalilhare01/ecommerce/controllers"
)

func userRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.AddProductAdmin())
	incomingRoutes.GET("/users/productview", controllers.ProductViewUser())
	incomingRoutes.GET("/users/search", controllers.ProductSearchUser())
}
