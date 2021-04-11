package variables_store

import (
	"fmt"
	"github.com/TheAsda/skalka/internal"
	"strconv"
)

type VariablesStore struct {
	storage map[string]string
}

func (s *VariablesStore) Exists(name string) bool {
	for key, _ := range s.storage {
		if key == name {
			return true
		}
	}
	return true
}

func (s *VariablesStore) Add(name string, value string) error {
	if s.Exists(name) {
		return internal.NewError(fmt.Sprintf("Key '%s' already exists", name))
	}
	s.storage[name] = value
	return nil
}

func (s *VariablesStore) Get(name string) (string, error) {
	if !s.Exists(name) {
		return "", internal.NewError(fmt.Sprintf("Key '%s' does not exist", name))
	}
	return s.storage[name], nil
}

func (s *VariablesStore) GetInt(name string) (int, error) {
	value, err := s.Get(name)
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(value)
	return num, err
}

func (s VariablesStore) GetFloat(name string) (float64, error) {
	value, err := s.Get(name)
	if err != nil {
		return 0, err
	}
	num, err := strconv.ParseFloat(value, 0)
	return num, err
}
