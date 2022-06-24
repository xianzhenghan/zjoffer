package zj050_第一个只出现一次的字符

import (
	"fmt"
	"testing"
)

// go test -v zj050.go zj050_test.go
func TestFirstUniqChar(t *testing.T) {
	s := "leetcode"
	first := FirstUniqChar(s)
	fmt.Printf("first = %s	", string(first))
}
