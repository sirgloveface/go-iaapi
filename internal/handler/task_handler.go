package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirgloveface/go-iaapi/internal/service"
)

type TaskHandler struct {
	svc *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{svc: s}
}

func (h *TaskHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/tasks", h.GetAll)
	r.POST("/tasks", h.Create)
	r.GET("/tasks/:id", h.GetByID)
}

func (h *TaskHandler) GetAll(c *gin.Context) {
	tasks, err := h.svc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) Create(c *gin.Context) {
	var payload struct {
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if err := h.svc.Create(payload.Title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *TaskHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	task, err := h.svc.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}
