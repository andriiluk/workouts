package exercisesvc

import (
	"context"
	"fmt"
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

	if len(p.Tags) > 0 {
		exercises, err := s.searchByTags(p.Tags...)
		if err != nil {
			return nil, fmt.Errorf("[%w]: search by tags", err)
		}

		res = append(res, exercises...)
	}

	if len(p.Muscles) > 0 {
		exercises, err := s.searchByMuscles(p.Muscles...)
		if err != nil {
			return nil, fmt.Errorf("[%w]: search by tags", err)
		}

		res = append(res, exercises...)
	}

	return res, nil
}

func (s *InMemStore) searchByTags(tags ...string) ([]*internal.Exercise, error) {
	var res []*internal.Exercise

	for i, exercise := range s.data {
		for _, tag := range tags {
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

func (s *InMemStore) searchByMuscles(muscles ...string) ([]*internal.Exercise, error) {
	var res []*internal.Exercise

	for i, exercise := range s.data {
		for _, muscle := range muscles {
			hasMuscle := false

			for _, exMuscle := range exercise.Muscles {
				if exMuscle == muscle {
					hasMuscle = true

					break
				}
			}

			if hasMuscle {
				res = append(res, s.data[i])

				break
			}
		}
	}

	return res, nil
}
