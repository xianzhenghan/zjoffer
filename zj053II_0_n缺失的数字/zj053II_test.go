package zj053II_0_n缺失的数字

import (
	"fmt"
	"testing"
)

//go test -v zj053II.go zj053II_test.go
func TestMissingNumber(t *testing.T) {
	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	miss := MissingNumber(num)
	fmt.Println(miss)
}
