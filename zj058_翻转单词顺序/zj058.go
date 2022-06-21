package zj058_翻转单词顺序

import (
	"strings"
)

//输入: "the sky is blue"
//输出: "blue is sky the"

//输入: "  hello world!  "
//输出: "world! hello"
// 58-1
func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	sbyte := []byte(s)
	sli := make([]string, 0)
	tmp := ""
	for _, v := range sbyte {
		if v == ' ' {
			if len(tmp) > 0 {
				sli = append(sli, tmp)
				tmp = ""
			}
		} else {
			tmp = tmp + string(v)
		}
	}
	sli = append(sli, tmp)
	length := len(sli)
	reslut := ""
	for length > 0 {
		reslut = reslut + " " + sli[length-1]
		length--
	}
	return strings.TrimSpace(reslut)
}

// reverseWords2  58-1
func reverseWords2(s string) string {
	strSlice := strings.Fields(s)
	left, right := 0, len(strSlice)-1
	for left < right {
		strSlice[left], strSlice[right] = strSlice[right], strSlice[left]
		left++
		right--
	}
	return strings.Join(strSlice, " ")
}

// reverseWords2  58-II

//输入: s = "abcdefg", k = 2
//输出: "cdefgab"

//输入: s = "lrloseumgh", k = 6
//输出: "umghlrlose"
func reverseLeftWords(s string, n int) string {
	length := len(s)
	sbyte := []byte(s)
	return string(sbyte[n-1:length-1]) + string(sbyte[0:n-1])
}

// 原地翻转 ，3次反转
func reverseLeftWords2(s string, n int) string {
	sbyte := []byte(s)
	reverseString2(sbyte, 0, n-1)
	reverseString2(sbyte, n, len(sbyte)-1)
	reverseString2(sbyte, 0, len(sbyte)-1)
	return string(sbyte)
}

func reverseString2(s []byte, left, right int) []byte {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}

// 原地翻转 ，3次反转
func reverseLeftWords3(s string, n int) string {
	sbyte := []byte(s)
	reverseString3(sbyte[0:n])
	reverseString3(sbyte[n:len(sbyte)])
	reverseString3(sbyte[0:len(sbyte)])
	return string(sbyte)
}

func reverseString3(s []byte) []byte {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}
