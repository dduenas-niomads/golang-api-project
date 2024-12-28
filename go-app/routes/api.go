package routes

import (
	"go-app/controllers"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	
	"go-app/middlewares"
)

type WelcomeMessage struct {
	Message string `json:"message"`
	Author  string `json:"author"`
}

func NewRouter(tagsController *controllers.TagsController) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		// message
		var welcomeMessage = WelcomeMessage{
			Message: "Tags management API | Golang | Gin | JWT | Docker",
			Author: "Daniel Dueñas - @github.com/dduenas-niomads"}
		// response
		ctx.JSON(http.StatusOK, welcomeMessage)
	})

	router.Group("/api")

	// tags
	// tagsRouter := baseRouter.Group("/tags")
	router.GET("/api/tags", middlewares.CheckAuth, tagsController.FindAll)
	router.GET("/api/tags/:tagId", middlewares.CheckAuth, tagsController.FindById)
	router.POST("/api/tags", middlewares.CheckAuth, tagsController.Create)
	router.PATCH("/api/tags/:tagId", middlewares.CheckAuth, tagsController.Update)
	router.DELETE("/api/tags/:tagId", middlewares.CheckAuth, tagsController.Delete)

	// auth
	// tagsRouter := baseRouter.Group("/auth")
	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)
	router.GET("/auth/profile", middlewares.CheckAuth, controllers.GetUserProfile)

	return router
}
