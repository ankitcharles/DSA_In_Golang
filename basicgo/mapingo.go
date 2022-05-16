package main

import "fmt"

func main() {
	cities := map[string]int{
		"Bangalore": 1234,
		"Hyedrabad": 5678,
		"Chennai":   91011,
	}
	for k, v := range cities {
		fmt.Printf("%s has %d inhabitants \n", k, v)
	}
}
