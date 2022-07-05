package zj065_不用加减乘除做加法

import "fmt"

func Add(a int, b int) int {
	// 进位
	var carry int
	for b != 0 {
		// 进位
		fmt.Printf("1 b != 0 ;a=%d,b=%d\n", a, b)
		carry = (a & b) << 1
		fmt.Printf("2 cmd (a & b) << 1; carry=%d \n", carry)
		// 不加进位
		a ^= b
		fmt.Printf("3 cmd a ^= b ; a=%d \n", a)
		// 加进位
		b = carry
		fmt.Printf("4 cmd b = carry ;b=%d \n", b)
		fmt.Printf("\n\n\n")
	}
	return a
}

func add(a int, b int) int {
	for b != 0 {
		return add(a^b, (a&b)<<1)
	}
	return a
}
