package controller

import (
	"net/http"

	"github.com/IkhsanDS/golang-api/app"
	"github.com/gin-gonic/gin"
)

type TodoController struct{ svc *app.TodoService }

func NewTodoController(s *app.TodoService) *TodoController { return &TodoController{svc: s} }

type createTodoReq struct {
	Title string `json:"title" binding:"required,min=3"`
}
type updateTodoReq struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

// @Summary Create todo
// @Tags todos
// @Accept json
// @Produce json
// @Param payload body createTodoReq true "payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/v1/todos [post]
func (h *TodoController) Create(c *gin.Context) {
	var req createTodoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t, err := h.svc.Create(req.Title)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, t)
}

// @Summary List todos
// @Tags todos
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /api/v1/todos [get]
func (h *TodoController) List(c *gin.Context) {
	items, err := h.svc.List()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, items)
}

// @Summary Get todo by ID
// @Tags todos
// @Param id path string true "Todo ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /api/v1/todos/{id} [get]
func (h *TodoController) Get(c *gin.Context) {
	id := c.Param("id")
	t, err := h.svc.Get(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, t)
}

// @Summary Update todo
// @Tags todos
// @Accept json
// @Param id path string true "Todo ID"
// @Param payload body updateTodoReq true "payload"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/todos/{id} [patch]
func (h *TodoController) Update(c *gin.Context) {
	id := c.Param("id")
	var req updateTodoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	t, err := h.svc.Update(id, req.Title, req.Completed)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, t)
}

// @Summary Delete todo
// @Tags todos
// @Param id path string true "Todo ID"
// @Success 204 "No Content"
// @Router /api/v1/todos/{id} [delete]
func (h *TodoController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.Delete(id); err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
