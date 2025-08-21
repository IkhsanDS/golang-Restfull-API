package router

import (
	"github.com/gin-gonic/gin"

	// PENTING: path harus persis sama dengan module di go.mod
	"github.com/IkhsanDS/golang-api/handlers"
	"github.com/IkhsanDS/golang-api/middlewares"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	r := gin.Default()

	// (opsional) CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/api/v1/auth")

	// Auth
	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)
	auth.GET("/me", middlewares.AuthRequired(), handlers.Me)

	// Todos
	api := r.Group("/api/v1")
	{
		api.GET("/todos", handlers.GetTodos)
		api.GET("/todos/:id", handlers.GetTodo)
		api.POST("/todos", middlewares.AuthRequired(), handlers.CreateTodo)
		api.PUT("/todos/:id", middlewares.AuthRequired(), handlers.UpdateTodo)
		api.DELETE("/todos/:id",
			middlewares.AuthRequired(),
			middlewares.RequireRoles("admin"),
			handlers.DeleteTodo,
		)
	}

	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	return r
}
