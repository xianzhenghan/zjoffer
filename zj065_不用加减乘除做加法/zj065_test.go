package zj065_不用加减乘除做加法

import (
	"fmt"
	"testing"
)

// go run -v zj065.go zj065_test.go
func TestAdd(t *testing.T) {
	numa := 127
	numb := 15
	sum := Add(numa, numb)
	fmt.Printf("sum=%d\n", sum)
}
