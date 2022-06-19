package zjoffer
import (
	"testing"
	"fmt"
)

func TestStack(t *testing.T) {
	obj := Constructor();
	obj.AppendTail(2222);
	popValue := obj.DeleteHead();
	t.Log("Pass")
	fmt.Println(popValue)
}

