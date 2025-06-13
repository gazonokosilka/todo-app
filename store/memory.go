package store

import (
	"errors"
	"sync"
	"todo-app/model"
)

type MemoryStore struct {
	tasks  map[int]model.Task
	mu     sync.RWMutex
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks:  make(map[int]model.Task),
		nextID: 1,
	}
}
func (s *MemoryStore) Create(task model.Task) model.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	task.ID = s.nextID
	s.nextID++
	s.tasks[task.ID] = task
	return task
}
func (s *MemoryStore) List() []model.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := make([]model.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		result = append(result, task)
	}
	return result
}
func (s *MemoryStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tasks[id]; ok {
		delete(s.tasks, id)
		return true
	}
	return false
}
func (s *MemoryStore) MarkDone(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if task, ok := s.tasks[id]; ok {
		task.Done = true
		s.tasks[id] = task
		return nil
	}
	return errors.New("task not found")
}

//func check(err error) {
//	if err != nil {
//		log.Fatal(err)
//	}
//}
