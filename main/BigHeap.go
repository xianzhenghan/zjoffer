package main

import "fmt"

/*
给定整数数组nums和k，
请返回数组中第k个最大元素，
请注意，你需要找的是数组排序后的第k个最大元素，
而不是第k个不同的元素
*/
func swap(a, b *int) {
	*a, *b = *b, *a
}
func less(a, b int) bool {
	return a < b
}

func HeapSort(nums []int) []int {
	// 堆排序,只能确认第一次个数是最大或最小的
	// 调换第一个元素和最后一个元素位置、从0倒数第二个继续堆排序
	i := len(nums)
	for i > 1 {
		buildHeep(nums, i)
		swap(&nums[0], &nums[i-1])
		i--
	}
	return nums
}

func buildHeep(nums []int, len int) {
	// 找到最后一个节点的父节点
	parent := len/2 - 1
	for parent >= 0 {
		heapify(nums, parent, len)
		parent--
	}
	fmt.Println(nums[0:len])
}

func heapify(nums []int, parent, len int) {
	// 判断两个子节点是否比父节点大，如果是的话替换
	max := parent
	lson := parent*2 + 1
	rson := parent*2 + 2
	if lson < len && less(nums[lson], nums[max]) {
		// 左节点是否大于父节点
		max = lson
	}
	if rson < len && less(nums[rson], nums[max]) {
		// 右节点是否大于父节点
		max = rson
	}
	if parent != max {
		swap(&nums[max], &nums[parent])
		heapify(nums, max, len)
	}
}

func main() {
	fmt.Println(HeapSort([]int{
		3, 5, 3, 0, 8, 6,
	}))
}
