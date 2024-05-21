package routes

import (
	"github.com/EvertAnderson/controllers"
	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine) {

	routes := router.Group("api/v1/books")
	routes.POST("", controllers.BookCreate)
	routes.GET("", controllers.BookGet)
	routes.GET("/:id", controllers.BookGetById)
	routes.PUT("/:id", controllers.BookUpdate)
	routes.DELETE("/:id", controllers.BookDelete)

}
