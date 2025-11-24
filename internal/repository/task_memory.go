package repository

import (
	"crud-app/internal/domain"
	"errors"
)

type TaskRepository interface {
	Create(task domain.Task) domain.Task
	GetAll() []domain.Task
	GetByID(id int) (domain.Task, error)
	Update(id int, task domain.Task) (domain.Task, error)
	Delete(id int) error
}

type TaskMemoryRepository struct {
	data   map[int]domain.Task
	nextID int
}

func NewTaskMemoryRepository() *TaskMemoryRepository {
	return &TaskMemoryRepository{
		data:   make(map[int]domain.Task),
		nextID: 1,
	}
}

func (r *TaskMemoryRepository) Create(task domain.Task) domain.Task {
	task.ID = r.nextID
	r.nextID++
	r.data[task.ID] = task
	return task
}

func (r *TaskMemoryRepository) GetAll() []domain.Task {
	tasks := make([]domain.Task, 0, len(r.data))
	for _, task := range r.data {
		tasks = append(tasks, task)
	}
	return tasks
}

func (r *TaskMemoryRepository) GetByID(id int) (domain.Task, error) {
	task, ok := r.data[id]
	if !ok {
		return domain.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (r *TaskMemoryRepository) Update(id int, task domain.Task) (domain.Task, error) {
	_, ok := r.data[id]
	if !ok {
		return domain.Task{}, errors.New("task not found")
	}
	task.ID = id
	r.data[id] = task
	return task, nil
}

func (r *TaskMemoryRepository) Delete(id int) error {
	_, ok := r.data[id]
	if !ok {
		return errors.New("task not found")
	}
	delete(r.data, id)
	return nil
}
