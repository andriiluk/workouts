package musclesvc

import (
	"context"
	"sync"

	"github.com/andriiluk/workouts/internal"
)

type InMemStore struct {
	index int
	mu    sync.RWMutex
	data  map[int]*internal.Muscle
}

func NewInMemStore() *InMemStore {
	return &InMemStore{
		data: make(map[int]*internal.Muscle),
	}
}

func (s *InMemStore) InsertOrUpdate(ctx context.Context, m *internal.Muscle) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if m.ID != 0 {
		s.data[m.ID] = m
		return nil
	}

	s.index++
	m.ID = s.index
	s.data[s.index] = m

	return nil
}

func (s *InMemStore) Get(ctx context.Context, id int) (*internal.Muscle, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data[id], nil
}

func (s *InMemStore) Delete(ctx context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, id)
	return nil
}

func (s *InMemStore) Search(ctx context.Context, p *internal.Params) ([]*internal.Muscle, error) {
	var res []*internal.Muscle
	for i, muscle := range s.data {
		for _, tag := range p.Tags {
			hasTag := false
			for _, mTag := range muscle.Tags {
				if mTag == tag {
					hasTag = true
					break
				}
			}
			if hasTag {
				res = append(res, s.data[i])
				break
			}
		}
	}

	return res, nil
}
