package zj004_二维数组中查找

import "fmt"

//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//]
//给定 target = 5， 返回true。
//给定 target = 20，返回alse
func FindNumberIn2DArray(matrix [][]int, target int) (bool, int, int) {
	if len(matrix) == 0 {
		return false, -1, -1
	}
	yNum := len(matrix[0]) //列数
	xNum := len(matrix)    //行数
	fmt.Printf("xNum =%d,yNum =%d \n)", xNum, yNum)
	x, y := 0, yNum-1
	for x < xNum && y >= 0 {
		if matrix[x][y] > target {
			y--
		} else if matrix[x][y] < target {
			x++
		} else {
			return true, x, y
		}
	}
	return false, -1, -1
}
