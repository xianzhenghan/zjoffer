package zj038_字符串的排列

import "strings"

//输入：s = "abc"
//输出：["abc","acb","bac","bca","cab","cba"]
//1 <= s 的长度 <= 8
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
