package zj004_二维数组中查找

import (
	"fmt"
	"testing"
)

// go test -v zj004.go zj004_test.go
func TestFindNumberIn2DArray(t *testing.T) {
	m := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	/*	m2 := [][]int{
		{1, 1},
	}*/

	in, x, y := FindNumberIn2DArray(m, 9)
	fmt.Printf("m in %t \n", in)
	fmt.Printf("m in (%d,%d)", x, y)
}
