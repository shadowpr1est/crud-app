package handler

import (
	"crud-app/internal/domain"
	"crud-app/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

// GetAll handles GET /tasks
// @Summary      Get all tasks
// @Tags         tasks
// @Produce      json
// @Success      200  {array}   domain.Task
// @Router       /tasks [get]
func (h *TaskHandler) GetAll(c *gin.Context) {
	tasks := h.taskService.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

// GetByID handles GET /tasks/:id
// @Summary      Get task by ID
// @Tags         tasks
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  domain.Task
// @Failure      404  {object}  map[string]string
// @Router       /tasks/{id} [get]
func (h *TaskHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := h.taskService.GetByIdTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Create handles POST /tasks
// @Summary      Create new task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        body  body      domain.Task  true  "Task body"
// @Success      201   {object}  domain.Task
// @Failure      400   {object}  map[string]string
// @Router       /tasks [post]
func (h *TaskHandler) Create(c *gin.Context) {
	var input domain.Task
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask := h.taskService.CreateTask(input)
	c.JSON(http.StatusOK, newTask)
}

// Update handles PUT /tasks/:id
// @Summary      Update task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id    path      int          true  "Task ID"
// @Param        body  body      domain.Task  true  "Updated task"
// @Success      200   {object}  domain.Task
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /tasks/{id} [put]
func (h *TaskHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[Update] Invalid task ID in URL: %s", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	log.Printf("[Update] Updating task with ID: %d", id)

	var input domain.Task
	err = c.ShouldBindJSON(&input)
	if err != nil {
		log.Printf("[Update] Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[Update] Input received: %+v", input)

	updatedTask, err := h.taskService.UpdateTask(id, input)
	if err != nil {
		log.Printf("[Update] Task with ID %d not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	log.Printf("[Update] Task updated successfully: %+v", updatedTask)
	c.JSON(http.StatusOK, updatedTask)
}

// Delete handles DELETE /tasks/:id
// @Summary      Delete task
// @Tags         tasks
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /tasks/{id} [delete]
func (h *TaskHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.taskService.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
