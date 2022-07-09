package zj067_把字符串转换成整数

import (
	"fmt"
	"strings"
)

const (
	INT_MAX = int32(^uint32((0)) >> 1)
	INT_MIN = ^INT_MAX
)

func strToInt(str string) int {
	str = strings.TrimSpace(str)
	strings.TrimLeft(str, "0")
	number := ""
	falg := ""
	for _, v := range str {
		switch {
		case v == '+' || v == '-':
			if len(number) == 0 && len(falg) == 0 {
				falg = string(v)
			} else {
				return getnumber(falg, number)
			}
			break
		case '0' <= v && v <= '9':
			number = number + string(v)
			break
		default:
			return getnumber(falg, number)
		} //switch
	} //for
	return getnumber(falg, number)
}

func getnumber(flag, number string) int {
	fmt.Printf("falg=%s;number=%s", flag, number)
	num := int64(0)
	for _, v := range number {
		tmp := int64(v - '0')
		if flag == "-" {
			tmp = -tmp
		}
		num = tmp + num*10

		if num >= int64(INT_MAX) {
			return int(INT_MAX)
		}
		if num <= int64(INT_MIN) {
			return int(INT_MIN)
		}
	}
	fmt.Printf("num=%d\n", num)
	if flag == "-" {
		num = -num
	}
	return int(num)
}
