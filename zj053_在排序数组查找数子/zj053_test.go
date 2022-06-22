package zj053_在排序数组查找数子

import (
	"fmt"
	"testing"
)

// go test -v zj053.go zj053_test.go
func TestSearch(t *testing.T) {
	nums := []int{1, 2, 3, 5, 5, 5, 8, 9}
	index := Search(nums, 5)
	fmt.Println(index)
}
