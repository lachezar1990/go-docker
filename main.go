package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedDate time.Time `json:"created_date"`
}

func main() {
	r := gin.Default()
	r.GET("/get-tasks", getTasks)
	r.Run(":5555")
}

func getTasks(c *gin.Context) {
	tasks := make([]Task, 0)

	tasks = append(tasks, Task{ID: 1, Name: "pyrvi test", Description: "pyrvo opisanie", CreatedDate: time.Now()})

	c.IndentedJSON(200, tasks)
}
