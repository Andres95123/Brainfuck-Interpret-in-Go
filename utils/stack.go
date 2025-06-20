package utils

type Stack struct {
	Stack   []int
	Pointer int
}

func Make_stack() Stack {
	return Stack{Stack: []int{}, Pointer: -1}
}

func (stack *Stack) Push(Value int) {
	stack.Pointer++
	stack.Stack = append(stack.Stack, Value)
}

func (stack *Stack) Pop() int {

	if stack.Pointer < 0 {
		panic("Stack underflow!")
	}

	value := stack.Stack[stack.Pointer]
	stack.Pointer--
	stack.Stack = stack.Stack[:len(stack.Stack)-1]
	return value
}

func (stack *Stack) Peak() int {
	return stack.Stack[stack.Pointer]
}

func (stack *Stack) IsEmpty() bool {
	return stack.Pointer < 0
}
