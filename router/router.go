package router

import (
	"github.com/IkhsanDS/golang-api/app"
	"github.com/IkhsanDS/golang-api/controller"
	"github.com/IkhsanDS/golang-api/middlewares"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gorm.io/gorm"
)

func Setup(r *gin.Engine, db *gorm.DB) {
	r.Use(middlewares.CORS())

	// wiring service & controller
	todoSvc := app.NewTodoService(db)
	todoCtl := controller.NewTodoController(todoSvc)

	// health
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	// api v1
	api := r.Group("/api/v1")
	{
		api.POST("/todos", todoCtl.Create)
		api.GET("/todos", todoCtl.List)
		api.GET("/todos/:id", todoCtl.Get)
		api.PATCH("/todos/:id", todoCtl.Update)
		api.DELETE("/todos/:id", todoCtl.Delete)
	}

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
