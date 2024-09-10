package main

import "fmt"

// 包括start 包括end
func partition(nums []int, start, end int) int {
	povit := nums[start]
	for start < end {
		for start < end && povit <= nums[end] {
			end--
		}
		nums[start] = nums[end]
		for start < end && povit >= nums[start] {
			start++
		}
		nums[end] = nums[start]
	}
	nums[start] = povit
	return start
}

func quickSort(nums []int, start, end int) {
	if end > start {
		povit := partition(nums, start, end)
		fmt.Println(povit)
		quickSort(nums, start, povit-1)
		quickSort(nums, povit+1, end)
	}
}

func main() {
	nums := []int{4, 5, 2, 3, 8, 9, 0}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
