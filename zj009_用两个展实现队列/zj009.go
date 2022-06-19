package zj009_用两个展实现队列

type Stack struct {
	Values [] int
}

func (s *Stack) Pop() int {
	length := len(s.Values)
	if(length == 0){
		return -1
	}
	popValue := s.Values[length-1]
	popStack := s.Values[:length-1]
	s.Values = popStack
	return popValue
}

func (s *Stack) Push(v int) {
	s.Values = append(s.Values,v)
}

func NewStack() Stack  {
	return Stack{
		Values: make([]int,0,5),
	}
}
// CQueue
type CQueue struct {
	inStack  Stack
	outStack Stack
}

func Constructor() CQueue {
	return CQueue{
		inStack: NewStack(),
		outStack: NewStack(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack.Push(value)
}

func (this *CQueue) DeleteHead() int {
	// 需要从instack 进行pop
	if(len(this.outStack.Values) == 0) {
		for len(this.inStack.Values) !=0 {
			this.outStack.Push(this.inStack.Pop())
		}
	}
	value := this.outStack.Pop()
	return value;
}
