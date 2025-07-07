package handlers

import (
	"Simple_http_api/models"
	"Simple_http_api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(task service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: task}
}

func (t *TaskHandler) CreateNewTaskHand(c *gin.Context) {
	task := t.taskService.CreateTask()
	c.JSON(http.StatusCreated, task)
}

func (t *TaskHandler) GetTaskByIdHand(c *gin.Context) {
	id := c.Param("id")
	task, ok := t.taskService.GetTask(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	responseTask := struct {
		models.Task
		DurationSec *float64 `json:"duration_sec,omitempty"`
	}{
		Task: *task,
	}
	if task.StartedAt != nil {
		var duration time.Duration
		if task.EndedAt != nil {
			duration = task.EndedAt.Sub(*task.StartedAt)
		} else if task.Status == models.Running {
			duration = time.Now().Sub(*task.StartedAt)
		}
		durationSec := duration.Seconds()
		responseTask.DurationSec = &durationSec
	}
	c.JSON(http.StatusOK, responseTask)
}
func (t *TaskHandler) GetAllTasksHand(c *gin.Context) {
	tasks := t.taskService.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

func (t *TaskHandler) DeleteTaskByIdHand(c *gin.Context) {
	id := c.Param("id")
	ok := t.taskService.DeleteTask(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
