package zj039_数组中出现次数超过一半的数字

import "fmt"

/*输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2*/
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	count := make(map[int]int, 0)
	length := len(nums)
	for i := 0; i < length; i++ {
		times, ok := count[nums[i]]
		if ok {
			count[nums[i]] = times + 1
		} else {
			count[nums[i]] = 1
		}
		fmt.Printf("coutn=%v\n", count)
		if times > length/2 {
			fmt.Printf("times=%d;leng=%d\n", times, length/2)
			return nums[i]
		}
	}
	return -1
}
