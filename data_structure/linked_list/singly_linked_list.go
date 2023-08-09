package linked_list

import (
	"errors"
	"fmt"
)

// Node holds data and next node
type Node struct {
	Data int
	Next *Node
}

// SingleList holds head of list and length of list
type SingleList struct {
	length int
	Head   *Node
}

// NewList returns new instance of single list
func NewList() *SingleList {
	return &SingleList{}
}

// AddFront add data to the head of list
func (s *SingleList) AddFront(data int) {
	newNode := &Node{
		Data: data,
		Next: s.Head, // add head to next of new node
	}

	// update head add new node to head
	s.Head = newNode

	// update length
	s.length++
}

// AddBack : add data at the end of list
func (s *SingleList) AddBack(data int) {
	defer func() {
		// increment length
		s.length++
	}()

	newNode := &Node{
		Data: data,
	}

	// if list is empty add to head
	if s.Head == nil {
		s.Head = newNode
		return
	}

	// consider head as current
	current := s.Head

	// parse util last node
	for current.Next != nil {
		current = current.Next
	}

	// add new node
	current.Next = newNode
}

func (s *SingleList) RemoveFront() error {
	if s.Head == nil {
		return errors.New("list is empty")
	}

	// set next of next to the head
	s.Head = s.Head.Next

	s.length--

	return nil
}

func (s *SingleList) RemoveBack() error {
	if s.Head == nil || s.Size() == 0 {
		return errors.New("removeBack : List is empty")
	}

	defer func() {
		// decrement  length
		s.length--
	}()

	// if only one item in list set head as nil
	if s.Head.Next == nil {
		s.Head = nil
		return nil
	}

	current := s.Head

	// parse util next of current
	for current.Next.Next != nil {
		// set current to next
		current = current.Next
	}

	current.Next = nil

	return nil
}

// Remove data from list
func (s *SingleList) Remove(value int) error {
	if s.Head == nil {
		return errors.New("remove : List is empty")
	}

	defer func() {
		// decrement  length
		s.length--
	}()

	// if only one item in list set head as next head
	if s.Head.Data == value {
		s.Head = s.Head.Next
		return nil
	}

	current := s.Head

	// parse util next of current
	for current.Next != nil {
		if current.Next.Data == value {
			current.Next = current.Next.Next
			return nil
		}

		current = current.Next
	}

	return nil
}

// Traverse visit each element
func (s *SingleList) Traverse() []int {
	var listData []int

	current := s.Head

	for current != nil {
		fmt.Print(current.Data, " ")
		listData = append(listData, current.Data)
		current = current.Next
	}

	fmt.Println()

	return listData
}

// Front return head value
func (s *SingleList) Front() (int, error) {
	if s.Head == nil {
		return 0, fmt.Errorf("front error: List is empty")
	}

	return s.Head.Data, nil
}

// Size returns length of linked list
func (s *SingleList) Size() int {
	return s.length
}
