package zj011_旋转数组的最小数字

// 输入：numbers = [3,4,5,1,2]
// 输出：1
func minArray(numbers []int) int {
	sli := make([]int, 0)
	sli = numbers
	length := len(sli)
	idx := 0
	for idx < length-1 {
		sli = append(sli, sli[idx])
		if sli[idx+1] < sli[length+idx] {
			return sli[idx+1]
		}
		if sli[idx+1] == sli[length+idx] {
			idx++
			continue
		}
		idx++
	}
	if idx == length-1 {
		return sli[0]
	}
	return -1
}

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left, right, mid := 0, len(nums)-1, 0
	for left < right {
		mid = (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}

	}
	return nums[left]
}
