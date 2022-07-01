package zj045_把数组排成最小的数

import (
	"fmt"
	"testing"
)

func TestMinNumber(t *testing.T) {
	nums := []int{121, 12}
	str := MinNumber(nums)
	fmt.Printf("str = %s", str)
}
