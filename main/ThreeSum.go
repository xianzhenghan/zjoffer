package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		ansTmps := twoSum(nums, i+1, 0-nums[i])
		for j := range ansTmps {
			ansTmps[j] = append(ansTmps[j], nums[i])
			ans = append(ans, ansTmps[j])
		}
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}

func twoSum(nums []int, start, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	left, right := start, len(nums)-1

	for left < right {
		rightValue := nums[right]
		leftValue := nums[left]
		if target == nums[left]+nums[right] {
			ans = append(ans, []int{nums[left], nums[right]})
			for left < right && nums[right] == rightValue {
				right--
			}
			for left < right && nums[left] == leftValue {
				left++
			}
		} else if target > nums[left]+nums[right] {
			for left < right && nums[left] == leftValue {
				left++
			}
		} else {
			for left < right && nums[right] == rightValue {
				right--
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
