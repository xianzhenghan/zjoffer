package zj030_min_stack函数

import (
	"math"
)

type MinStack struct {
	Vlaues []int
	min []int
}

func Constructor() MinStack {
	return MinStack{
		Vlaues: make([]int,0,0),
		min: make([]int,0,0),
	}
}

func (this *MinStack) Pop()  {
	length := len(this.Vlaues)
	if(length > 0){
		this.Vlaues = this.Vlaues[:length-1]
		this.min = this.min[:length-1]
	}
}

func (this *MinStack) Push(x int)  {
	curMin := this.Min()
	this.Vlaues = append(this.Vlaues,x)
	if(x < curMin){
		curMin = x;
	}
	this.min = append(this.min,curMin)
}

func (this *MinStack) Top() int {
	length := len(this.Vlaues)
	if(length <= 0){
		return -1
	}
	return this.Vlaues[length-1]
}

func (this *MinStack) Min() int {
	length := len(this.Vlaues)
	if length == 0  {
		return math.MaxInt64
	}
	return this.min[length-1]
}