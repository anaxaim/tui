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

func (s String) Slice() []string {
	slice := make([]string, 0, len(s))
	for item := range s {
		slice = append(slice, item)
	}
	sort.Strings(slice)
	return slice
}
