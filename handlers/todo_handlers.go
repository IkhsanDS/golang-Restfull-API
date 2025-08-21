package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/IkhsanDS/golang-api/database"
	"github.com/IkhsanDS/golang-api/models"
	"github.com/gin-gonic/gin"
)

// GetTodos godoc
// @Summary      List todos
// @Tags         todos
// @Produce      json
// @Router       /todos [get]
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	// pagination dll (boleh yang versi kamu)
	q := database.DB.Model(&models.Todo{})
	if s := c.Query("q"); s != "" {
		q = q.Where("title LIKE ?", "%"+s+"%")
	}
	if comp := c.Query("completed"); comp != "" {
		if v, err := strconv.ParseBool(comp); err == nil {
			q = q.Where("completed = ?", v)
		}
	}
	if sort := c.Query("sort"); sort != "" {
		if strings.HasPrefix(sort, "-") {
			q = q.Order(strings.TrimPrefix(sort, "-") + " DESC")
		} else {
			q = q.Order(sort + " ASC")
		}
	} else {
		q = q.Order("created_at DESC")
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	var total int64
	q.Count(&total)
	q.Offset((page - 1) * limit).Limit(limit).Find(&todos)

	c.JSON(http.StatusOK, gin.H{"page": page, "limit": limit, "total": total, "items": todos})
}

// GetTodo godoc
// @Summary      Get todo by ID
// @Tags         todos
// @Produce      json
// @Param        id path int true "Todo ID"
// @Router       /todos/{id} [get]
func GetTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

type CreateTodoInput struct {
	Title string `json:"title" binding:"required"`
}

// CreateTodo godoc
// @Summary      Create todo
// @Tags         todos
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        payload body CreateTodoInput true "Create payload"
// @Success      201 {object} models.Todo
// @Failure      400 {object} gin.H
// @Router       /todos [post]
func CreateTodo(c *gin.Context) {
	var in CreateTodoInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t := models.Todo{Title: in.Title}
	if err := database.DB.Create(&t).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot create"})
		return
	}
	c.JSON(http.StatusCreated, t)
}

type UpdateTodoInput struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

// UpdateTodo godoc
// @Summary      Update todo
// @Tags         todos
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id path int true "Todo ID"
// @Param        payload body UpdateTodoInput true "Update payload"
// @Router       /todos/{id} [put]
func UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	var in UpdateTodoInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if in.Title != nil {
		todo.Title = *in.Title
	}
	if in.Completed != nil {
		todo.Completed = *in.Completed
	}
	database.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

// DeleteTodo godoc
// @Summary      Delete todo
// @Tags         todos
// @Security     BearerAuth
// @Produce      json
// @Param        id path int true "Todo ID"
// @Router       /todos/{id} [delete]
func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := database.DB.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
