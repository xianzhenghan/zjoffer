package zj009_用两个展实现队列
import (
	"testing"
	"fmt"
)
// TestStack   go test -v zj009.go zj009_test.go
func TestStack(t *testing.T) {
	obj := Constructor();
	obj.AppendTail(2222);
	popValue := obj.DeleteHead();
	t.Log("Pass")
	fmt.Println(popValue)
}

