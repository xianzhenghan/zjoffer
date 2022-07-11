package zj019_正则表达式匹配

import (
	"fmt"
	"testing"
)

func TestIsMatch(t *testing.T) {
	str := ""
	pattern := ""
	b := IsMatch(str, pattern)
	fmt.Printf("str=%s;pattern=%s result=%t\n", str, pattern, b)
}
