package main

import "fmt"

func mySwap(a, b *int) {
	*a, *b = *b, *a
}

/*
*
待排序数组
*/
func mySort(nums []int) []int {
	length := len(nums)
	for length > 1 {
		myBuildHeap(nums, length)
		mySwap(&nums[0], &nums[length-1])
		length--
	}
	return nums
}

/*
进行对 nums 数组；长度为length 进行排
*/
func myBuildHeap(nums []int, length int) {

	parent := length/2 - 1
	for parent >= 0 {
		myHeapify(nums, parent, length)
		parent--
	}
	fmt.Println(nums[:])
}

func myHeapify(nums []int, parent, length int) {
	max := parent
	leftSon := 2*parent + 1
	rightSon := 2*parent + 2
	if leftSon < length && (nums[leftSon] > nums[max]) {
		max = leftSon
	}

	if rightSon < length && (nums[rightSon] > nums[max]) {
		max = rightSon
	}

	if max != parent {
		mySwap(&nums[max], &nums[parent])
		myHeapify(nums, max, length)
	}
}

func main() {
	fmt.Println(findKthLargest([]int{7, 8, 1, 2, 5, 6, 6}, 1))
}

func findKthLargest(nums []int, k int) int {
	length := len(nums)
	times := k
	for length > 1 && times > 0 {
		myBuildHeap(nums, length)
		mySwap(&nums[0], &nums[length-1])
		length--
		times--
	}
	return nums[len(nums)-k]
}
