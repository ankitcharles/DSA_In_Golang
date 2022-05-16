package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		if i%2 == 0 {
			continue
		}
		pow[i] = 1 << uint(i)

	}
	fmt.Println(pow)
}
