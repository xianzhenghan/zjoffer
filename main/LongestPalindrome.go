package main

func longestPalndrome(s string) (res string) {
	res, length := "", len(s)
	for i := 0; i < length; i++ {
		// 偶数
		l, r := i, i
		for l > 0 && r < length && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > len(res) {
			res = s[l+1 : r]
		}
		// 奇数
		l, r = i, i+1
		for l > 0 && r < length && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > len(res) {
			res = s[l+1 : r]
		}
	}
	return
}

func main() {
	println(longestPalndrome("babad"))
}
