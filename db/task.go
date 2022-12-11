package db

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string
	Description string
}

func GetTasks() ([]Task, error) {
	var tasks []Task

	err := db.Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
