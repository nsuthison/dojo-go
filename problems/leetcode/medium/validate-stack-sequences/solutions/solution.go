package solutions

func validateStackSequences(pushed []int, popped []int) bool {

	stack := make(Stack, 0)
	idxPop := 0

	for idxPush := 0; idxPush < len(pushed); idxPush++ {
		stack.Push(pushed[idxPush])

		topElementInStack, _ := stack.Peek()
		for idxPop < len(popped) && popped[idxPop] == topElementInStack {
			stack.Pop()
			idxPop++

			if len(stack) > 0 {
				topElementInStack, _ = stack.Peek()
			} else {
				break
			}
		}
	}

	if idxPop == len(popped) {
		return true
	}

	return false
}

// Stack is stack
type Stack []int

// Push element to stack
func (stack *Stack) Push(toPush int) {
	*stack = append(*stack, toPush)
}

// Pop element from stack
func (stack *Stack) Pop() (result int, canPop bool) {

	if len(*stack) <= 0 {
		return 0, false
	}

	lastIndex := len(*stack) - 1

	result = (*stack)[lastIndex]
	*stack = (*stack)[:lastIndex]

	return result, true
}

// Peek top element in stack
func (stack *Stack) Peek() (result int, canPeek bool) {
	if len(*stack) <= 0 {
		return -1, false
	}

	return (*stack)[len(*stack)-1], true
}
