package domain

import "time"

type Task struct {
	ID          string
	Description string
	Status      string
	Project     string
	Entered     time.Time
	Due         time.Time
	UUID        string
	Urgency     float32
}

type TaskRepository interface {
	Store(task Task) error
	GetTasks() ([]Task, error)
	Count() (int, error)
}

type TaskManager struct {
	TaskRepo TaskRepository
}

func (tm *TaskManager) AddTask(task Task) error {
	return tm.TaskRepo.Store(task)
}

func (tm *TaskManager) GetTasks() ([]Task, error) {
	return tm.TaskRepo.GetTasks()
}

func (tm *TaskManager) Count() (int, error) {
	return tm.TaskRepo.Count()
}
