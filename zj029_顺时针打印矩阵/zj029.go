package zj029_顺时针打印矩阵

import "fmt"

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	nums := make([]int, 0)
	fmt.Printf("right=%d;bottom=%d\n", len(matrix[0])-1, len(matrix)-1)
	left, top, right, bottom := 0, 0, len(matrix[0])-1, len(matrix)-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			fmt.Printf("(%d,%d)\n", i, top)
			nums = append(nums, matrix[top][i])
		}
		top++
		if top > bottom {
			continue
		}
		fmt.Printf("\n")
		for j := top; j <= bottom; j++ {
			fmt.Printf("(%d,%d)\n", right, j)
			nums = append(nums, matrix[j][right])
		}
		right--
		if right < left {
			continue
		}
		fmt.Printf("\n")
		for m := right; m >= left; m-- {
			fmt.Printf("(%d,%d)\n", m, bottom)
			nums = append(nums, matrix[bottom][m])
		}
		bottom--
		if bottom < top {
			continue
		}
		fmt.Printf("\n")
		for n := bottom; n >= top; n-- {
			fmt.Printf("(%d,%d)\n", left, n)
			nums = append(nums, matrix[n][left])
		}
		left++
		if left > right {
			continue
		}
		fmt.Printf("\n")
	}
	return nums
}
