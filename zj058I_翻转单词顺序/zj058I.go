package zj058I_翻转单词顺序

import "strings"

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
