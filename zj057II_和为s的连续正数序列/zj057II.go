package zj057II_和为s的连续正数序列

func findContinuousSequence(target int) [][]int {
	if target == 1 || target == 2 {
		return [][]int{}
	}
	left, right, sum := 1, 1, 0
	res := make([][]int, 0)
	for left <= (target >> 1) {
		if sum < target {
			sum += right
			right++
		} else if sum > target {
			sum -= left
			left++
		} else {
			tmp := make([]int, 0)
			for i := left; i < right; i++ {
				tmp = append(tmp, i)
			}
			res = append(res, tmp)
			sum -= left
			left++
		}
	}
	return res
}

func findContinuousSequence2(target int) [][]int {
	if target == 1 || target == 2 {
		return [][]int{}
	}
	seqs := make([][]int, 0)
	lengthMax := target / 2
	for i := 1; i <= lengthMax; i++ {
		sum := 0
		t := make([]int, 0)
		for j := i; j <= lengthMax+1 && sum <= target; j++ {
			sum = sum + j
			t = append(t, j)
			if sum == target && len(t) > 0 {
				seqs = append(seqs, t)
			}
		}
	}
	return seqs
}
