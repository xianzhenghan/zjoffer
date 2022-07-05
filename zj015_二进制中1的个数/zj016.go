package zj015_二进制中1的个数

func hammingWeight(num uint32) int {
	rn := 0
	for num > 0 {
		if num&1 == 1 {
			rn++
			num = num >> 1
		}
	}
	return rn
}
