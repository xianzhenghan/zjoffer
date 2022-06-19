package zj030_min_stack函数

import (
	"fmt"
	"testing"
	)
// go test -v zj030.go zj030_test.go
func TestMinStack_Min(t *testing.T) {
	obj := Constructor();
	obj.Push(-1);
	obj.Push(7);
	obj.Push(-10);
	obj.Push(90);
	obj.Pop();
	obj.Pop();
	top := obj.Top();
	fmt.Println(top)
	min := obj.Min();
	fmt.Println(min)
}