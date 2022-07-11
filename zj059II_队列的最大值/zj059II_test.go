package zj059II_队列的最大值

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor()
	obj.Push_back(2)
	obj.Push_back(3)
	pop := obj.Pop_front()
	fmt.Printf("pop=%d", pop)
}
