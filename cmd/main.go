// cmd/main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirgloveface/go-iaapi/internal/config"
	"github.com/sirgloveface/go-iaapi/internal/handler"
	"github.com/sirgloveface/go-iaapi/internal/repository"
	"github.com/sirgloveface/go-iaapi/internal/service"
)

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
	// PostgreSQL connection
	// db := config.ConnectPostgres()
	//repo := repository.NewInMemoryTaskRepository()

	var repo repository.TaskRepository

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

	h.RegisterRoutes(r)

	r.Run() // :8080
}
