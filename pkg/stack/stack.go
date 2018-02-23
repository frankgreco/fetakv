package stack

import (
	"sync"

	"github.com/frankgreco/fetakv/pkg/transaction"
)

// Stack respresents a LIFO data structure
type Stack struct {
	mutex sync.RWMutex
	data  []transaction.Interface
}

// New creates a new instance of a stack.
func New() *Stack {
	return &Stack{
		mutex: sync.RWMutex{},
		data:  []transaction.Interface{},
	}
}

// Pop removes and returns the topmost element on the stack.
// If the stack if empty, nil is returned.
func (s *Stack) Pop() transaction.Interface {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.pop()
}

func (s *Stack) pop() transaction.Interface {
	if s.data == nil || len(s.data) < 1 {
		return nil
	}
	trans := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return trans
}

// Peek returns the topmost element on the stack.
// If the stack if empty, nil is returned.
func (s *Stack) Peek() transaction.Interface {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.peek()
}

func (s *Stack) peek() transaction.Interface {
	if s.data == nil || len(s.data) < 1 {
		return nil
	}
	return s.data[len(s.data)-1]
}

// Push adds an element onto the top of the stack.
// If the stack if empty, nil is returned.
// Nil elements are not allowed.
func (s *Stack) Push(t transaction.Interface) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.push(t)
}

func (s *Stack) push(t transaction.Interface) {
	if t != nil {
		s.data = append(s.data, t)
	}
}

// Size returns the current stack size.
func (s *Stack) Size() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.size()
}

func (s *Stack) size() int {
	return len(s.data)
}
