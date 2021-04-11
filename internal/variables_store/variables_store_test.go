package variables_store

import (
	"fmt"
	"testing"
)

func TestVariablesStore_Add(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		store := NewVariablesStore()
		err := store.Add("test", "test")
		if err != nil {
			t.Errorf(fmt.Sprintf("Add error: %s", err.Error()))
		}
	})

	t.Run("add exiting", func(t *testing.T) {
		store := NewVariablesStore()
		_ = store.Add("test", "test")
		err := store.Add("test", "test")
		if err == nil {
			t.Errorf("Add exiting key did not return error")
		}
	})
}

func TestVariablesStore_Exists(t *testing.T) {
	store := NewVariablesStore()
	_ = store.Add("test", "test")

	t.Run("exists", func(t *testing.T) {
		exists := store.Exists("test")
		if !exists {
			t.Errorf("Name must exist")
		}
	})

	t.Run("not exists", func(t *testing.T) {
		exists := store.Exists("random")
		if exists {
			t.Errorf("Name must noy exist")
		}
	})
}

func TestVariablesStore_Get(t *testing.T) {
	store := NewVariablesStore()
	_ = store.Add("test", "test")

	t.Run("get", func(t *testing.T) {
		value, err := store.Get("test")
		if err != nil {
			t.Errorf("Error getting: %s", err.Error())
		}
		if value != "test" {
			t.Errorf("Invalid value '%s'", value)
		}
	})

	t.Run("not get", func(t *testing.T) {
		_, err := store.Get("random")
		if err == nil {
			t.Errorf("No error with wrong name")
		}
	})
}

func TestVariablesStore_GetFloat(t *testing.T) {
	store := NewVariablesStore()
	_ = store.Add("test", "123.123")
	_ = store.Add("test2", "abcd")

	t.Run("get float", func(t *testing.T) {
		value, err := store.GetFloat("test")
		if err != nil {
			t.Errorf("Error getting: %s", err.Error())
		}
		if value != 123.123 {
			t.Errorf("Invalid value '%f'", value)
		}
	})

	t.Run("not get float", func(t *testing.T) {
		_, err := store.GetFloat("random")
		if err == nil {
			t.Errorf("No error with wrong name")
		}
	})

	t.Run("not get float with wrong format", func(t *testing.T) {
		_, err := store.GetFloat("test2")
		if err == nil {
			t.Errorf("No error with wrong format")
		}
	})
}

func TestVariablesStore_GetInt(t *testing.T) {
	store := NewVariablesStore()
	_ = store.Add("test", "123")
	_ = store.Add("test2", "abcd")

	t.Run("get int", func(t *testing.T) {
		value, err := store.GetInt("test")
		if err != nil {
			t.Errorf("Error getting: %s", err.Error())
		}
		if value != 123 {
			t.Errorf("Invalid value '%d'", value)
		}
	})

	t.Run("not get int", func(t *testing.T) {
		_, err := store.GetInt("random")
		if err == nil {
			t.Errorf("No error with wrong name")
		}
	})

	t.Run("not get int with wrong format", func(t *testing.T) {
		_, err := store.GetInt("test2")
		if err == nil {
			t.Errorf("No error with wrong format")
		}
	})
}
