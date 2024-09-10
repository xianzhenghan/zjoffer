package zj038_字符串的排列

import (
	"fmt"
	"strings"
)

// 输入：s = "abc"
// 输出：["abc","acb","bac","bca","cab","cba"]
// 1 <= s 的长度 <= 8
func Permutation(s string) []string {
	dp := rec(s)
	return strings.Split(dp[len(s)-1], ",")
}

func rec(s string) []string {
	dp := make([]string, len(s))
	dp[0] = string(s[0])
	for i := 1; i < len(s); i++ {
		add := string(s[i])
		tmp := make([]string, 0)
		sli := strings.Split(dp[i-1], ",")
		for _, str := range sli {
			for j := 0; j < len(str)+1; j++ {
				if j == 0 {
					tmp = append(tmp, add+str)
				} else if j == len(str) {
					tmp = append(tmp, str+add)
				} else {
					tmp = append(tmp, str[:j]+add+str[j:])
				}
			}
		}
		tmp = removeDuplicateElement(tmp)
		dp[i] = strings.Join(tmp, ",")
	}
	return dp
}

func removeDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
//
// 换句话说，s1 的排列之一是 s2 的 子串 。
//
// 示例 1：
//
// 输入：s1 = "ab" s2 = "eidbaooo"
// 输出：true
// 解释：s2 包含 s1 的排列之一 ("ba").
// 示例 2：
//
// 输入：s1= "ab" s2 = "eidboaoo"
// 输出：false
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	strSlices := getStrSlices(s1)
	for _, str := range strSlices {
		if strings.Contains(s2, str) {
			return true
		}
	}
	return false
}

func getStrSlices(str string) []string {
	dp := make([][]string, 0)
	str0 := []string{string(str[0])}
	dp = append(dp, str0)
	for i := 1; i < len(str); i++ {
		char := str[i]
		strSlice := dp[i-1]
		sub_ans := make([]string, 0)
		for _, s := range strSlice {
			tmpans := addCharToStrSlice(s, char)
			sub_ans = append(sub_ans, tmpans...)
		}
		dp = append(dp, sub_ans)
	}
	return dp[len(str)-1]
}

func addCharToStrSlice(s string, c byte) []string {
	if len(s) == 0 {
		return []string{string(c)}
	}
	ans := make([]string, 0)
	ans = append(ans, string(c)+s)
	for i := 1; i < len(s); i++ {
		ans = append(ans, s[:i]+string(c)+s[i:])
	}
	ans = append(ans, s+string(c))
	return ans
}

func main() {
	fmt.Println(checkInclusion("abc", "a"))
}

func checkInclusionPlus(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := range s1 {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if cnt1 == cnt2 {
			return true
		}
		cnt2[s2[i]-'a']--
		cnt2[s2[i+len(s1)]-'a']++
	}
	return cnt1 == cnt2
}
