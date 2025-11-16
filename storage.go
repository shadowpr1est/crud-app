package main

import "errors"

type TaskStorage struct {
	data   map[int]Task
	nextID int
}

func NewTaskStorage() *TaskStorage {
	return &TaskStorage{
		data:   make(map[int]Task),
		nextID: 1,
	}
}

// POST

func (s *TaskStorage) Create(t Task) Task {
	t.ID = s.nextID
	s.nextID++
	s.data[t.ID] = t
	return t
}

// GET

func (s *TaskStorage) GetAll() []Task {
	res := make([]Task, 0, len(s.data))
	for _, t := range s.data {
		res = append(res, t)
	}
	return res
}

func (s *TaskStorage) GetByID(id int) (Task, error) {
	t, ok := s.data[id]
	if !ok {
		return Task{}, errors.New("Not found")
	}
	return t, nil
}

// UPDATE

func (s *TaskStorage) Update(id int, t Task) (Task, error) {
	_, ok := s.data[id]
	if !ok {
		return Task{}, errors.New("Couldn't update")
	}
	t.ID = id
	s.data[id] = t
	return t, nil
}

// DELETE

func (s *TaskStorage) Delete(id int) error {
	_, ok := s.data[id]
	if !ok {
		return errors.New("Not found")
	}
	delete(s.data, id)
	return nil
}
