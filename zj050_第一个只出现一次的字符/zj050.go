package zj050_第一个只出现一次的字符

import "fmt"

// 输入：s = "abaccdeff"
// 输出：'b'
func FirstUniqChar(s string) byte {
	sli := [26]int{}
	sbtes := []byte(s)
	for _, value := range sbtes {
		idx := value - 'a'
		sli[idx]++
	}
	fmt.Printf("sli=%v", sli)
	for _, value := range sbtes {
		if sli[value-'a'] == 1 {
			return value
		}
	}
	return ' '
}
