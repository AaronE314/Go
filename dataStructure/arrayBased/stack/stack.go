package main

import (
	"errors"
	"fmt"
	"sync"
)

// Stack object
type Stack struct {
	elements []int
	lock     *sync.Mutex
}

func newStack() *Stack {
	return &Stack{make([]int, 0), &sync.Mutex{}}
}

func (s *Stack) push(value int) {

	s.lock.Lock()
	defer s.lock.Unlock()

	s.elements = append(s.elements, value)
}

func (s Stack) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack) len() int {
	return len(s.elements)
}

func (s *Stack) pop() (int, error) {

	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.elements) == 0 {
		return -1, errors.New("underflow: the stack is empty")
	}

	l := len(s.elements)

	value := s.elements[l-1]
	s.elements = s.elements[:l-1]

	return value, nil
}

func (s Stack) peek() (int, error) {

	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.elements) == 0 {
		return -1, errors.New("the stack is empty")
	}

	return s.elements[len(s.elements)-1], nil
}

func main() {

	stack := newStack()

	fmt.Println(stack.len())
	fmt.Println(stack.pop())
	fmt.Println(stack.peek())
	fmt.Println(stack.isEmpty())

	for i := 0; i < 10; i++ {
		stack.push(i)
	}

	fmt.Println(stack.len())
	fmt.Println(stack.peek())
	fmt.Println(stack.isEmpty())

	for !stack.isEmpty() {
		fmt.Println(stack.pop())
	}

}
