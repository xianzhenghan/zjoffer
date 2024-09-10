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

func minArrayDemo(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		mid := left + (right-left)>>1
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] <= numbers[right] {
			right = right - 1
		}
	}
	return numbers[left]
}
