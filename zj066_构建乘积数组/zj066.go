package zj066_构建乘积数组

//输入: [1,2,3,4,5]
//输出: [120,60,40,30,24]

func constructArr(a []int) []int {
	if len(a) == 0 {
		return a
	}
	length := len(a)

	lr := make([]int, length)
	rl := make([]int, length)
	lr[0] = 1
	for i := 1; i < length; i++ {
		lr[i] = a[i-1] * lr[i-1]
	}
	r := make([]int, length)
	rl[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		rl[i] = a[i+1] * rl[i+1]

	}
	for i := 0; i < length; i++ {
		r[i] = lr[i] * rl[i]
	}
	return r
}
