package zj053_在排序数组查找数子

//输入: nums = [5,7,7,8,8,10], target = 8
//输出: 2
func Search(nums []int, target int) int {
	l := BinarySearch(nums, target, func(mid, target int) bool { return nums[mid] < target })
	r := BinarySearch(nums, target, func(mid, target int) bool { return nums[mid] <= target })
	return r - l
}

// BinarySearch 二分查找
func BinarySearch(nums []int, target int, cond func(int, int) bool) int {
	l, r := -1, len(nums)
	for l+1 != r {
		m := (l + r) / 2
		if cond(m, target) {
			l = m
		} else {
			r = m
		}
	}
	return l
}

//
func IsBluel(mid, target int) bool { return mid < target }
func IsBluer(mid, target int) bool { return mid <= target }

func search(nums []int, t int) int {
	N := len(nums)
	bisearch := func(cond func(int, int) bool) int {
		l, r := -1, N
		for l+1 != r {
			m := (l + r) >> 1
			if cond(m, t) {
				l = m
			} else {
				r = m
			}
		}
		return l
	}
	r := bisearch(func(m, t int) bool { return nums[m] <= t })
	l := bisearch(func(m, t int) bool { return nums[m] < t })
	return r - l
}
