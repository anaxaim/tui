package utils

import (
	"sort"
)

type Empty struct{}

type String map[string]Empty

func NewString(items ...string) String {
	ss := String{}
	ss.Insert(items...)

	return ss
}

// Insert adds items to the set.
func (s String) Insert(items ...string) String {
	for _, item := range items {
		s[item] = Empty{}
	}

	return s
}

// Has returns true if and only if item is contained in the set.
func (s String) Has(item string) bool {
	_, contained := s[item]
	return contained
}

func (s String) Slice() []string {
	slice := make([]string, 0, len(s))
	for item := range s {
		slice = append(slice, item)
	}

	sort.Strings(slice)

	return slice
}
