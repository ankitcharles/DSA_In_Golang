package main

import "fmt"

func twosum(nums []int, target int) []int {
	low := 0
	high := len(nums) - 1
	for low < high {
		currentSum := nums[low] + nums[high]
		if currentSum == target {
			return []int{low + 1, high + 1}
		}
		if currentSum > target {
			high--
		}
		if currentSum < target {
			low++
		}
	}
	return nil
}

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twosum(arr, target))
}
