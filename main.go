package main

import (
	"log"
	"moqt-go/docker/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	ErrorMessage string `json:"error"`
}

func main() {
	err := db.SetupDB()

	if err != nil {
		log.Fatalf("DB err: %v", err)
	}

	r := gin.Default()
	r.GET("/tasks", getTasks)
	r.GET("/task/:id", getTaskByID)
	r.POST("/task", addTask)
	r.DELETE("/task/:id", deleteTaskByID)
	r.Run(":5555")
}

func getTasks(c *gin.Context) {
	tasks, err := db.GetTasks()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, APIError{ErrorMessage: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskByID(c *gin.Context) {
	taskID := c.Param("id")
	taskIDInt, err := strconv.Atoi(taskID)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, APIError{ErrorMessage: err.Error()})
		return
	}

	task, err := db.GetTaskByID(int32(taskIDInt))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, APIError{ErrorMessage: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func deleteTaskByID(c *gin.Context) {
	taskID := c.Param("id")
	taskIDInt, err := strconv.Atoi(taskID)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, APIError{ErrorMessage: err.Error()})
		return
	}

	err = db.DeleteTaskByID(int32(taskIDInt))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, APIError{ErrorMessage: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

type taskInputModel struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func addTask(c *gin.Context) {
	var task taskInputModel
	err := c.Bind(&task)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, APIError{ErrorMessage: err.Error()})
		return
	}

	taskFromDB, err := db.AddTask(db.Task{Name: task.Name, Description: task.Description})

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, APIError{ErrorMessage: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, taskFromDB)
}
