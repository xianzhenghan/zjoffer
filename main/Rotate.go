package main

import "fmt"

/*
*

	矩阵旋转90度

1 水平翻转
2 对角线翻转
*/
func Rotate(matrix [][]int) {
	length := len(matrix)
	// 1 水平翻转
	for i := 0; i < length/2; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[length-i-1][j] = matrix[length-i-1][j], matrix[i][j]
		}
	}
	// 对角线翻转
	for i := 0; i < length; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	Rotate(matrix)
	fmt.Println(matrix)
}
