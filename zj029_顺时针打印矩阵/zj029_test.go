package zj029_顺时针打印矩阵

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	res := SpiralOrder(nums)
	fmt.Printf("res=%v", res)
}
