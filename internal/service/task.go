package service

import (
	"crud-app/internal/domain"
	"crud-app/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateTask(task domain.Task) domain.Task {
	return s.repo.Create(task)
}

func (s *TaskService) GetAllTasks() []domain.Task {
	return s.repo.GetAll()
}

func (s *TaskService) GetByIdTask(id int) (domain.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) UpdateTask(id int, task domain.Task) (domain.Task, error) {
	return s.repo.Update(id, task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}
