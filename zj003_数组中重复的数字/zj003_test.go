package zj003_数组中重复的数字

import (
	"fmt"
	"testing"
)

// go test -v zj004.go zj004_test.go
func TestFindRepeatNumber2(t *testing.T) {
	nums := []int{1, 2, 3, 3, 4, 5}
	repeat := FindRepeatNumber2(nums)
	fmt.Println(repeat)
}
