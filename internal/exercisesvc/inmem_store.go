package exercisesvc

import (
	"context"
	"sync"

	"github.com/andriiluk/workouts/internal"
)

type InMemStore struct {
	data  map[int]*internal.Exercise
	mu    sync.RWMutex
	index int
}

func NewInMemStore() *InMemStore {
	return &InMemStore{
		data: make(map[int]*internal.Exercise),
	}
}

func (s *InMemStore) InsertOrUpdate(ctx context.Context, m *internal.Exercise) error {
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

func (s *InMemStore) Get(ctx context.Context, id int) (*internal.Exercise, error) {
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

func (s *InMemStore) Search(ctx context.Context, p *internal.Params) ([]*internal.Exercise, error) {
	var res []*internal.Exercise

	for i, exercise := range s.data {
		for _, tag := range p.Tags {
			hasTag := false

			for _, mTag := range exercise.Tags {
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
