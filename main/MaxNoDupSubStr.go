package main

import "fmt"

func maxNoDupSubStr(str string) int {
	if len(str) < 2 {
		return len(str)
	}
	left, right, length := 0, 1, len(str)
	m := make(map[byte]bool, len(str))
	max := 1
	m[str[left]] = true
	for left < length && right < length {
		ok, _ := m[str[right]]
		if ok {
			delete(m, str[left])
			left++
		} else {
			m[str[right]] = true
			right++
		}
		if len(m) > max {
			max = len(m)
		}
	}
	return max
}

func main() {
	str := "abcabcbb"
	fmt.Println(maxNoDupSubStr(str))
}
