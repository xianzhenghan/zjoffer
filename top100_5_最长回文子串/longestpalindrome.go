package top100_5_最长回文子串

import "fmt"

func LongestPalindrome(s string) string {
	max := new(int)
	res := new(string)
	rec(s, max, res)
	return *res
}

func LongestPalindrome2(s string) string {
	max := 0
	maxstr := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			fmt.Printf("substr=%s ", s[i:j])
			if isPalindrome(s[i:j]) {
				fmt.Printf("is=%s ", s[i:j])
				if len(s[i:j]) > max {
					maxstr = s[i:j]
					max = len(s[i:j])
				}
			}
			fmt.Printf("\n")
		}
	}
	return maxstr
}

func LongestPalindrome3(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	start, end := 0, 0
	max := 0
	for i := 0; i < n; i++ {
		left := i - 1
		right := i + 1
		length := 1
		//向左扩张
		for left >= 0 && s[left] == s[i] {
			left--
			length++
		}
		//向右扩张
		for right < len(s) && s[right] == s[i] {
			right++
			length++
		}
		//向两侧扩张
		for right < len(s) && left >= 0 && s[right] == s[left] {
			right++
			left--
			length += 2
		}
		if length > max {
			max = length
			start, end = left, right
		}
	}
	fmt.Printf("start=%d;end=%d max=%d\n", start, end, max)
	return s[start+1 : start+1+max]
}

func rec(s string, max *int, res *string) {
	if isPalindrome(s) {
		if len(s) > *max {
			*max = len(s)
			*res = s
		}
	} else {
		if s[0] == s[len(s)-2] {
			rec(s[:len(s)-1], max, res)
		}
		if s[1] == s[len(s)-1] {
			rec(s[1:], max, res)
		}
	}
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
