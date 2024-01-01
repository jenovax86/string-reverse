package main

import (
	"fmt"
	"strings"
)

type Stack []string

func (stack *Stack) push(value string) {
	*stack = append(*stack, value)
}

func (stack *Stack) pop() string {
	length := len(*stack)
	if length == 0 {
		panic("Stack is empty")
	}
	result := (*stack)[length-1]
	*stack = (*stack)[0 : length-1]
	return result
}

func (stack *Stack) isEmpty() bool {
	return len(*stack) == 0
}

func stringReverser() []string {
	var result []string
	var value string
	fmt.Scan(&value)
	stack := make(Stack, 0)
	for i := 0; i < len(value); i++ {
		characters := strings.Split(value, "")
		stack.push(characters[i])
	}

	for !stack.isEmpty() {
		result = append(result, stack.pop())
	}
	return result
}

func main() {
	fmt.Println(stringReverser())
}
