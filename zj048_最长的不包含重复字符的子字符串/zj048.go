package zj048_最长的不包含重复字符的子字符串

//输入: "abcabcbb"
//输出: 3
// 滑动窗口
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	set := make(map[byte]bool)
	res := 1
	left := 0
	set[s[0]] = true
	for right := 1; right < len(s); {
		if !set[s[right]] {
			set[s[right]] = true
			right++
		} else {
			delete(set, s[left])
			left++
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}
func lengthOfLongestSubstring(s string) int {
	//滑动窗口+hash
	if len(s) == 0 {
		return 0
	}
	res := 1
	set := make(map[byte]bool)
	set[s[0]] = true
	left := 0
	for right := 1; right < len(s); {
		if !set[s[right]] {
			set[s[right]] = true
			right++
		} else {
			//往右滑动一个元素
			delete(set, s[left])
			left++
		}
		if res < right-left {
			res = right - left
		}
	}
	return res
}
