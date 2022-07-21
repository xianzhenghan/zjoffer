package top_100_17_电话号码的字母组合

import "fmt"

/*0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。*/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := [][]string{
		{},
		{},
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}
	dp := make([][]string, len(digits))
	fmt.Printf("d=%d", (digits[0] - '0'))
	dp[0] = m[(digits[0] - '0')]
	for i := 1; i < len(digits); i++ {
		tmp := make([]string, 0)
		dgs := m[(digits[i] - '0')]
		for _, v := range dp[i-1] {
			for _, digit := range dgs {
				tmp = append(tmp, v+digit)
			}
		}
		dp[i] = tmp
	}
	return dp[len(digits)-1]
}
