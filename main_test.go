package main

import "testing"

func TestStackUsingQueue(t *testing.T) {
	var s Stack
	if _, ok := s.Pop(); ok {
		t.Error("pop on empty must fail")
	}
	s.Push(1)
	s.Push(2)
	if v, _ := s.Top(); v != 2 {
		t.Errorf("Top = %d, want 2", v)
	}
	s.Push(3)
	want := []int{3, 2, 1}
	for _, w := range want {
		if v, _ := s.Pop(); v != w {
			t.Errorf("Pop = %d, want %d (LIFO)", v, w)
		}
	}
	// Interleaved use keeps LIFO order.
	s.Push(5)
	s.Push(6)
	s.Pop()
	s.Push(7)
	if v, _ := s.Pop(); v != 7 {
		t.Errorf("Pop = %d, want 7", v)
	}
	if v, _ := s.Pop(); v != 5 {
		t.Errorf("Pop = %d, want 5", v)
	}
}
