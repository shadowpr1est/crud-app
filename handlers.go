package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	storage *TaskStorage
}

func newHandler(storage *TaskStorage) *Handler {
	return &Handler{storage: storage}
}

// GetAll
// @Summary      Get all tasks
// @Tags         tasks
// @Produce      json
// @Router       /tasks [get]
func (h *Handler) GetAll(c *gin.Context) {
	tasks := h.storage.GetAll()
	c.JSON(http.StatusOK, tasks)
}

// GetByID
// @Summary      Get task by ID
// @Tags         tasks
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Router       /tasks/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := h.storage.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Create
// @Summary      Create new task
// @Tags         tasks
// @Produce      json
// @Param        body  body      Task  true  "Task body"
// @Router       /tasks [post]
func (h *Handler) Create(c *gin.Context) {
	var input Task
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask := h.storage.Create(input)
	c.JSON(http.StatusOK, newTask)
}

// Update
// @Summary      Update task
// @Tags         tasks
// @Produce      json
// @Param        id    path      int   true "Task ID"
// @Param        body  body      Task  true "Updated task"
// @Router       /tasks/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input Task
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTask, err := h.storage.Update(id, input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.JSON(http.StatusOK, updatedTask)

}

// Delete
// @Summary      Delete task
// @Tags         tasks
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Router       /tasks/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.storage.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
