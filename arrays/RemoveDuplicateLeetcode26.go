package main

import "fmt"

func removeduplicates1(nums []int) int {
	size := len(nums)
	if size == 0 {
		return size
	}
	i := 0
	for j := 1; j < size; j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
func main() {
	i := []int{1, 1, 2, 2, 2, 3, 3}
	//fmt.Println(removeDuplicates(i))
	fmt.Println(removeduplicates1(i))
}
