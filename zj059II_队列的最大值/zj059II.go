package zj059II_队列的最大值

import (
	"fmt"
	"math"
)

type MaxQueue struct {
	StackIn     []int
	StackInMax  []int
	Stackout    []int
	StackoutMax []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		StackIn:     make([]int, 0, 0),
		StackInMax:  make([]int, 0, 0),
		Stackout:    make([]int, 0, 0),
		StackoutMax: make([]int, 0, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	li := len(this.StackInMax)
	lo := len(this.StackoutMax)
	if lo > 0 && li > 0 {
		vo := this.StackoutMax[len(this.StackoutMax)-1]
		vi := this.StackInMax[len(this.StackInMax)-1]
		return int(math.Max(float64(vo), float64(vi)))
	}
	if li > 0 {
		vi := this.StackInMax[len(this.StackInMax)-1]
		return vi
	}

	if lo > 0 {
		vo := this.StackoutMax[len(this.StackoutMax)-1]
		return vo
	}
	return -1
}

func (this *MaxQueue) Push_back(value int) {
	this.StackIn = append(this.StackIn, value)
	if len(this.StackInMax) == 0 {
		this.StackInMax = append(this.StackInMax, value)
	} else {
		currMax := this.StackInMax[len(this.StackInMax)-1]
		if value > currMax {
			this.StackInMax = append(this.StackInMax, value)
		} else {
			this.StackInMax = append(this.StackInMax, currMax)
		}
	}
}

func (this *MaxQueue) Pop_front() int {
	if len(this.Stackout) == 0 {
		if len(this.StackIn) == 0 { //没有元素
			return -1
		} else {
			for length := len(this.StackIn) - 1; length >= 0; length-- {
				fmt.Printf("length=%d; this.StackIn=%v\n", length, this.StackIn)
				value := this.StackIn[length]
				//stack in 减少
				this.StackIn = this.StackIn[:length]
				this.StackInMax = this.StackInMax[:length]
				// stack out 增加
				this.Stackout = append(this.Stackout, value)
				// stack max out 增加
				if len(this.StackoutMax) == 0 {
					this.StackoutMax = append(this.StackoutMax, value)
				} else {
					currMax := this.StackoutMax[len(this.StackoutMax)-1]
					if value > currMax {
						this.StackoutMax = append(this.StackoutMax, value)
					} else {
						this.StackoutMax = append(this.StackoutMax, currMax)
					}
				}
			}
		}
	}
	//stack in 减少
	v := this.Stackout[len(this.Stackout)-1]
	this.Stackout = this.Stackout[:len(this.Stackout)-1]
	this.StackoutMax = this.StackoutMax[:len(this.StackoutMax)-1]
	return v
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
