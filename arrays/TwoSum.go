package main

import "fmt"

func findtwosum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		if v, found := m[target-num]; found {
			return []int{v, index}
		}
		m[num] = index
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	result := []int{}
	target := 9
	result = findtwosum(nums, target)
	fmt.Println(result)

}
