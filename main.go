package main

import (
	"log"
	"moqt-go/docker/db"
	"net/http"

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
	r.GET("/get-tasks", getTasks)
	r.Run(":5555")
}

func getTasks(c *gin.Context) {
	tasks, err := db.GetTasks()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, APIError{ErrorMessage: err.Error()})
		return
	}

	c.IndentedJSON(200, tasks)
}
