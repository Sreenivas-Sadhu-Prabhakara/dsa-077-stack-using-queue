package main

import "fmt"

// Stack implements LIFO behavior on top of a single FIFO queue: after each
// push, the queue is rotated so the newest element stands at the front.
type Stack struct {
	q []int // used strictly as a queue: append at back, remove from front
}

// Push enqueues v, then rotates the previous elements behind it.
func (s *Stack) Push(v int) {
	s.q = append(s.q, v)
	for i := 0; i < len(s.q)-1; i++ {
		s.q = append(s.q[1:], s.q[0]) // dequeue front, enqueue it at the back
	}
}

// Pop removes and returns the most recently pushed value.
func (s *Stack) Pop() (int, bool) {
	if len(s.q) == 0 {
		return 0, false
	}
	v := s.q[0]
	s.q = s.q[1:]
	return v, true
}

// Top returns the most recently pushed value without removing it.
func (s *Stack) Top() (int, bool) {
	if len(s.q) == 0 {
		return 0, false
	}
	return s.q[0], true
}

// Size reports how many values are stored.
func (s *Stack) Size() int { return len(s.q) }

func main() {
	var s Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for s.Size() > 0 {
		v, _ := s.Pop()
		fmt.Print(v, " ")
	}
	fmt.Println()
}
