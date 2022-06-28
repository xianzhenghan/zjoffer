package zjoffer

//输入：nums = [2,7,11,15], target = 9
//输出：[2,7] 或者 [7,2]
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nums
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left]+nums[right] > target {
			right--
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			return []int{nums[left], nums[right]}
		}
	}
	return []int{}
}
