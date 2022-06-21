package zj003_数组中重复的数字

// 使用哈希表
func findRepeatNumber(nums []int) int {
	m := make(map[int]int, 0)
	for k, v := range nums {
		_, ok := m[v]
		if ok {
			return v
		} else {
			m[v] = k
		}
	}
	return -1
}

// 原地交换， 时间O(n), 空间O(1) 比较理想的解法
func findRepeatNumber2(nums []int) int {
	// 原地交换， 时间O(n), 空间O(1)
	length := len(nums)
	i := 0
	for i < length {
		// 遍历的索引对应的值等于索引，不动，继续遍历
		if i == nums[i] {
			i++
			continue
		}
		// 即nums[i] == num[x]
		// 如果i下标对应的值与下标i的元素值做索引的值相等，说明遇到重复值，返回该元素
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		// 不等，交换两值
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}
	return -1
}
