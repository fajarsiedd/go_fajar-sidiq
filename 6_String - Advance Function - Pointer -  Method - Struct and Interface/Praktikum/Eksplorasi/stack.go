package main

import (
	"fmt"
	"slices"
)

type Stack struct {
	ids    map[string]struct{}
	values []string
}

func (s *Stack) push(value string) {
	_, exists := s.ids[value]

	if !exists {
		s.ids[value] = struct{}{}
		s.values = append(s.values, value)
	}
}

func (s *Stack) pop() {
	if s.isEmpty() {
		fmt.Println("Stack kosong!")
	} else {
		s.values = slices.Delete(s.values, len(s.values)-1, len(s.values))
		delete(s.ids, s.values[len(s.values)-1])
	}
}

func (s *Stack) isEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) value() []string {
	return s.values
}

func main() {
	stack := Stack{
		ids:    make(map[string]struct{}),
		values: []string{},
	}
	fmt.Println("isEmpty", stack.isEmpty())
	stack.pop()

	stack.push("First")
	stack.push("Second")
	stack.push("Third")
	stack.push("Fourth")

	fmt.Println(stack.value())
	stack.pop()
	fmt.Println(stack.value())
	fmt.Println("isEmpty", stack.isEmpty())
}
