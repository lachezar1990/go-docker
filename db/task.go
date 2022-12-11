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

func GetTaskByID(id int32) (Task, error) {
	var task Task

	err := db.First(&task, id).Error

	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func DeleteTaskByID(id int32) error {
	err := db.Where("id = ?", id).Delete(&Task{}).Error

	if err != nil {
		return err
	}

	return nil
}

func AddTask(task Task) (Task, error) {
	err := db.Create(&task).Error

	if err != nil {
		return Task{}, err
	}

	return task, nil
}
