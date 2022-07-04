package mst040_最小的k个数

func getLeastNumbers2(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr[:k]
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
