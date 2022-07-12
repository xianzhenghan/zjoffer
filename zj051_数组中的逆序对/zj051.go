package zj051_数组中的逆序对

func reversePairs2(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		tmp := 0
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				tmp++
			}
		}
		dp[i] = dp[i-1] + tmp
	}
	return dp[len(nums)-1]
}

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}
	nums := quicksort(arr, 0, len(arr)-1)
	return nums[:k]
}

func quicksort(arr []int, left, right int) []int {
	if left < right {
		mid := pattition(arr, left, right)
		quicksort(arr, left, mid-1)
		quicksort(arr, mid+1, right)
	}
	return arr
}

// 包含前后
func pattition(arr []int, left, right int) int {
	piv := arr[left]
	for left < right {
		for left < right && arr[right] >= piv {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= piv {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = piv
	return left
}
