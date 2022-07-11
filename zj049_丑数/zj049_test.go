package zj049_丑数

import (
	"fmt"
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	num := 1000
	ugly := NthUglyNumber(num)
	fmt.Printf("ugly=%d", ugly)
}
