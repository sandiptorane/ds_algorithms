package stack

import (
	"errors"
	"sync"
)

// NewStack : returns new instance of stack
func NewStack() *Stack {
	return new(Stack)
}

// Stack holds data
// this stack build by using slice
type Stack struct {
	data []string
	lock sync.Mutex
}

// Push : push the data to top of stack
func (s *Stack) Push(value string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	// add data to top of stack
	s.data = append(s.data, value)
}

// Pop : pop the data from stack. It will remove data at top
func (s *Stack) Pop() (string, error) {

	// if stack is empty return error
	if s.Empty() {
		return "", errors.New("pop error : Stack is empty")
	}

	length := s.Size()

	popedData := s.data[length-1]

	s.lock.Lock()
	defer s.lock.Unlock()

	// pop top
	s.data = s.data[:length-1]

	return popedData, nil
}

// Size returns the size of stack
func (s *Stack) Size() int {
	return len(s.data)
}

// Empty : returns if stack is empaty or not
func (s *Stack) Empty() bool {
	return s.Size() == 0
}
