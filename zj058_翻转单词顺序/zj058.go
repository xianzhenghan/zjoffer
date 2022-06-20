package zj058_翻转单词顺序

import "strings"

//输入: "the sky is blue"
//输出: "blue is sky the"

//输入: "  hello world!  "
//输出: "world! hello"
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

// reverseWords2
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
