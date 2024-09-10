package main

import (
	"fmt"
)

func shellSort1(nums []int) []int {
	for gap := len(nums) / 2; gap > 0; gap = gap / 2 {
		fmt.Println(gap)
		for i := gap; i < len(nums); i++ {
			for j := i - gap; j >= 0 && nums[j] > nums[j+gap]; j = j - gap {
				nums[j+gap], nums[j] = nums[j], nums[j+gap]
			}
		}
	}
	return nums
}

func shellSort2(nums []int) []int {
	for gap := len(nums) / 2; gap > 0; gap = gap / 2 {
		fmt.Println(gap)
		for i := gap; i < len(nums); i++ {
			for j := i; j-gap >= 0 && nums[j-gap] > nums[j]; j = j - gap {
				nums[j-gap], nums[j] = nums[j], nums[j-gap]
			}
		}
	}

	return nums
}

func main() {
	fmt.Println(shellSort2([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}))
}
