package stack

import "sync"

type StackItem struct {
	item []int
	next *StackItem
}

type Stack struct {
	sp    *StackItem
	depth int
	lock  sync.Mutex
}

func New() *Stack {
	stack := new(Stack)

	stack.depth = 0
	return stack
}

func (s *Stack) Count() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.depth
}

func (s *Stack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.depth == 0
}

func (stack *Stack) Push(item []int) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.sp = &StackItem{item: item, next: stack.sp}
	stack.depth++
}

func (stack *Stack) Pop() []int {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.depth > 0 {
		item := stack.sp.item
		stack.sp = stack.sp.next
		stack.depth--
		return item
	}
	return nil
}

func (stack *Stack) Peek() []int {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.depth > 0 {
		return stack.sp.item
	}
	return nil
}
