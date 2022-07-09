package zjffer020_表示数值的字符串

import (
	"fmt"
	"strings"
)

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	number := ""
	poiternumer := ""
	enumner := ""
	point := false
	etype := false
	frontfalg := ""
	eflag := ""
	s = strings.TrimSpace(s)
	strings.TrimLeft(s, "0")
	for _, v := range s {
		switch {
		case v == '+' || v == '-':
			if len(number) == 0 && len(frontfalg) == 0 {
				frontfalg = string(v)
			} else if len(number) > 0 && len(enumner) == 0 && etype == true && len(eflag) == 0 {
				eflag = string(v)
			} else if len(enumner) > 0 && etype == true {
				return false
			} else {
				return false
			}
			break
		case v == ' ':
			return false
		case v == '.':
			if point == true {
				return false
			} else if len(number) == 0 && etype == false {
				point = true
				number = "00"
			} else if len(number) > 0 && etype == false {
				point = true
			} else {
				return false
			}
			break
		case v == 'E' || v == 'e':
			if len(number) == 0 || len(enumner) > 0 {
				return false
			} else if etype == true {
				return false
			} else {
				etype = true
			}
			break
		case '0' <= v && v <= '9':
			if point == false && etype == false {
				number = number + string(v)
			} else if point == true && etype == false {
				poiternumer = poiternumer + string(v)
			} else {
				enumner = enumner + string(v)
			}
			break
		default:
			return false
		} //switch
	} //for
	fmt.Printf("number=%s\n", number)
	fmt.Printf("poiternumer=%s\n", poiternumer)
	fmt.Printf("enumner=%s\n", enumner)

	if point == true && len(poiternumer) == 0 && number == "00" {
		return false
	}
	if len(number) == 0 {
		return false
	}
	if etype == true && len(enumner) == 0 {
		return false
	}
	return true
}
