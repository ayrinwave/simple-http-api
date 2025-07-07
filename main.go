package main

import (
	"Simple_http_api/service"
	"github.com/gin-gonic/gin"
	"log"
	"test-git/handlers"
)

func main() {
	r := gin.Default()

	taskService := service.NewTaskService()

	taskHandler := handlers.NewTaskHandler(taskService)

	r.POST("/task", taskHandler.CreateNewTaskHand)
	r.DELETE("/task/:id", taskHandler.DeleteTaskByIdHand)
	r.GET("/task/:id", taskHandler.GetTaskByIdHand)
	r.GET("/tasks", taskHandler.GetAllTasksHand)

	log.Println("Listening on port 8080")
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
