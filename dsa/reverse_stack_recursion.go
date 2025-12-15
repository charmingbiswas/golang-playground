package dsa

import "github.com/charmingbiswas/golang-stl/stack"

func ReverseStack(st *stack.Stack[int]) {
	if st.Size() == 1 {
		return
	} else {
		val := st.Top()
		st.Pop()
		ReverseStack(st)
		insertAtCorrectPosition(st, val)
	}
}

func insertAtCorrectPosition(st *stack.Stack[int], val int) {
	if st.IsEmpty() {
		st.Push(val)
		return
	} else {
		curr := st.Top()
		st.Pop()
		insertAtCorrectPosition(st, val)
		st.Push(curr)
	}
}
