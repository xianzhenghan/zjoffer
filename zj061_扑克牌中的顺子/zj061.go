package zj061_扑克牌中的顺子

import (
	"math"
	s "sort"
)

// 比较推荐的解法
func isStraight2(nums []int) bool {
	s.Ints(nums)
	joker := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[joker] < 5
	// 最大值-最小值 < 5
}

func isStraight(nums []int) bool {
	if len(nums) != 5 {
		return false
	}
	quees := getNumQuee(nums)
	nums = sort(nums)
	diffnum := diff(nums)
	if diffnum == 0 {
		return true
	}
	return quees >= diffnum

}

func getNumQuee(nums []int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			num++
		}
	}
	return num
}

func sort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func diff(nums []int) int {
	dsum := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > 0 {
			if nums[i+1] == (nums[i]) {
				dsum = 1 << 32
			} else {
				dsum = dsum + int(math.Abs(float64(nums[i+1]-(nums[i]+1))))
			}
		}
	}
	return dsum
}
