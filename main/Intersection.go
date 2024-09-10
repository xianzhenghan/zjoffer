package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	sli := make([]int, 0)
	for _, value := range nums1 {
		m[value] = 1
	}

	for _, value := range nums2 {
		if m[value] == 1 {
			sli = append(sli, value)
			m[value] = m[value] + 1
		}
	}
	return sli
}
func main() {
	sli1 := []int{1, 2, 2, 1}
	sli2 := []int{2, 2}
	fmt.Println(intersection(sli1, sli2))
}
