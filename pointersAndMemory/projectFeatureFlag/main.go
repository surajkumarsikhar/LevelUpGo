package main

import (
	"fmt"
	"slices"
)

type Flag struct {
	Name    string
	Enabled bool
	Rollout int
}

type Store struct {
	Flags []Flag
}

// Write IsEnabled: return whether the named flag is on for this user.
// Unknown flag or disabled flag means off. Otherwise roll out by bucket.
func (s *Store) IsEnabled(name string, userID int) bool {
	f := s.Find(name)
	if f == nil {
		return false
	} else if !f.Enabled {
		return false
	} else if userID%100 >= f.Rollout {
		return false
	}
	return true
}

// --- Store methods from earlier lessons. No need to change these. ---

func (s *Store) Set(name string, enabled bool) {
	for i := range s.Flags {
		if s.Flags[i].Name == name {
			s.Flags[i].Enabled = enabled
			return
		}
	}
	s.Flags = append(s.Flags, Flag{Name: name, Enabled: enabled})
}

func (s *Store) Find(name string) *Flag {
	for i := range s.Flags {
		if s.Flags[i].Name == name {
			return &s.Flags[i]
		}
	}
	return nil
}

func (s *Store) Remove(name string) {
	for i := range s.Flags {
		if s.Flags[i].Name == name {
			s.Flags = slices.Delete(s.Flags, i, i+1)
			return
		}
	}
}

func main() {
	store := &Store{}
	store.Set("new-checkout", true)
	if f := store.Find("new-checkout"); f != nil {
		f.Rollout = 30
	}

	for _, id := range []int{7, 42, 95} {
		fmt.Println(id, store.IsEnabled("new-checkout", id))
	}
}
