package solutions

import models "github.com/nsuthison/dojo-go/problems/leetcode/medium/validate-stack-sequences/models"

func validateStackSequences(pushed []int, popped []int) bool {

	stack := make(models.Stack, 0)
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
