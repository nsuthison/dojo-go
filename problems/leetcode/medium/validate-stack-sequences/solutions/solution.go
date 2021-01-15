package solutions

func validateStackSequences(pushed []int, popped []int) bool {
	//stack := make(Stack, 0)

	// for idxPush := 0; idxPush < len(pushed); idxPush++ {
	// 	for idxPop := 0; idxPop < len(popped); idxPop++ {

	// 	}
	// }

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
