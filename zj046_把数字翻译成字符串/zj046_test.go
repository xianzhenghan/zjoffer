package zj046_把数字翻译成字符串

import (
	"fmt"
	"testing"
)

// go test -v zj046.go zj046_test.go
func TestTranslateNum(t *testing.T) {
	num := 0
	count := TranslateNum(num)
	fmt.Printf("count=%d", count)
}
