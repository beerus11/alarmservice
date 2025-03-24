package storage

import (
	"alarmservice/models"
	"sync"
)

type MemoryStore struct {
	mu     sync.RWMutex
	alarms map[string]models.Alarm
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		alarms: make(map[string]models.Alarm),
	}
}

func (s *MemoryStore) Create(alarm models.Alarm) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.alarms[alarm.ID] = alarm
}

func (s *MemoryStore) GetAll() []models.Alarm {
	s.mu.RLock()
	defer s.mu.RUnlock()
	alarms := []models.Alarm{}
	for _, a := range s.alarms {
		alarms = append(alarms, a)
	}
	return alarms
}

func (s *MemoryStore) GetByID(id string) (models.Alarm, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	a, ok := s.alarms[id]
	return a, ok
}

func (s *MemoryStore) Update(id string, alarm models.Alarm) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.alarms[id]; !ok {
		return false
	}
	s.alarms[id] = alarm
	return true
}

func (s *MemoryStore) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.alarms[id]; !ok {
		return false
	}
	delete(s.alarms, id)
	return true
}
