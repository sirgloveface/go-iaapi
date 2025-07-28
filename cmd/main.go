// cmd/main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirgloveface/go-iaapi/internal/auth"
	"github.com/sirgloveface/go-iaapi/internal/config"
	"github.com/sirgloveface/go-iaapi/internal/handler"
	"github.com/sirgloveface/go-iaapi/internal/middleware"
	"github.com/sirgloveface/go-iaapi/internal/repository"
	"github.com/sirgloveface/go-iaapi/internal/service"
)

func main() {
	r := gin.Default()
	var repo repository.TaskRepository

	// Middlewares
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.LoggerMiddleware())

	// Public routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// PostgreSQL connection
	// db := config.ConnectPostgres()
	//repo := repository.NewInMemoryTaskRepository()

	// switch os.Getenv("DB")
	switch "mongo" {
	case "mongo":
		mongo := config.ConnectMongo()
		repo = repository.NewMongoTaskRepository(mongo)
	default:
		db := config.ConnectPostgres()
		repo = repository.NewPostgresTaskRepository(db)
	}

	//repo := repository.NewPostgresTaskRepository(db)
	svc := service.NewTaskService(repo)
	h := handler.NewTaskHandler(svc)

	//h.RegisterRoutes(r)

	r.POST("/login", handler.LoginHandler)
	// Protected routes (requieren token)
	protected := r.Group("/tasks")
	//protected.Use(middleware.AuthMiddleware())
	protected.Use(auth.JWTAuthMiddleware())
	{
		protected.POST("", h.Create)
		protected.GET("", h.GetAll)
		protected.GET("/:id", h.GetByID)
	}

	r.Run() // :8080
}
