package main

import (
	"errors"
	"fmt"
	"sync"
)

// Queue object
type Queue struct {
	elements []int
	lock     *sync.Mutex
}

func newQueue() *Queue {
	return &Queue{make([]int, 0), &sync.Mutex{}}
}

func (q *Queue) enqueue(value int) {

	q.lock.Lock()
	defer q.lock.Unlock()

	q.elements = append(q.elements, value)
}

func (q Queue) isEmpty() bool {
	return len(q.elements) == 0
}

func (q Queue) len() int {
	return len(q.elements)
}

func (q *Queue) dequeue() (int, error) {

	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.elements) == 0 {
		return -1, errors.New("underflow: the stack is empty")
	}

	value := q.elements[0]
	q.elements = q.elements[1:]

	return value, nil
}

func (q Queue) peek() (int, error) {

	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.elements) == 0 {
		return -1, errors.New("the stack is empty")
	}

	return q.elements[0], nil
}

func main() {

	queue := newQueue()

	fmt.Println(queue.len())
	fmt.Println(queue.dequeue())
	fmt.Println(queue.peek())
	fmt.Println(queue.isEmpty())

	for i := 0; i < 10; i++ {
		queue.enqueue(i)
	}

	fmt.Println(queue.len())
	fmt.Println(queue.peek())
	fmt.Println(queue.isEmpty())

	for !queue.isEmpty() {
		fmt.Println(queue.dequeue())
	}

}
