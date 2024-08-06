package data

import (
	"time"

	model "github.com/Task-Management-go/models"
)

// Mock data for tasks
var tasks = []model.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func GetTasks() []model.Task {
	return tasks
}

func GetTaskById(ID string) *model.Task {

	for i, val := range tasks {
		if val.ID == ID {
			return &tasks[i]
		}
	}
	return nil
}

func UpdateItem(ID string, updatedTask model.Task) *model.Task {
	for i, val := range tasks {
		if val.ID == ID {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}

			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}

			if updatedTask.Status == "In Progress" || updatedTask.Status == "Completed" || updatedTask.Status == "Pending" {
				tasks[i].Status = updatedTask.Status
			}

			return &tasks[i]
		}
	}
	return nil

}

func DeleteTask(ID string) *model.Task {
	for i, val := range tasks {
		if val.ID == ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return &val
		}

	}
	return nil
}
