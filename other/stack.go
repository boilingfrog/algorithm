package main

import "fmt"

type stack struct {
	Sil   []interface{}
	Count int
	N     int
}

func Init(n int) *stack {
	return &stack{
		N:   n,
		Sil: make([]interface{}, n),
	}
}

func (s *stack) Push(data int) {
	if len(s.Sil) < (s.Count + 1) {
		s.Sil = append(s.Sil, data)
	} else {
		s.Sil[s.Count] = data

	}
	s.Count++
	fmt.Println(s.Sil)
}

func (s *stack) Pop() interface{} {
	if s.Count == 0 {
		return nil
	}
	s.Count--
	return s.Sil[s.Count]
}
