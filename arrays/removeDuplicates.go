package main

import "fmt"

//https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
func removeDuplicates(nums []int) []int {
	size := len(nums)
	i := 2
	for j := i; j < size; j++ {
		if nums[i-2] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	return nums
}
func main() {
	i := []int{1, 1, 2, 2, 2, 3, 3}
	//fmt.Println(removeDuplicates(i))
	fmt.Println(removeDuplicates(i))
}
