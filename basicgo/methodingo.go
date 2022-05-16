package main

import "fmt"

type User1 struct {
	FirstName, LastName string
}

func (u User1) Greetings() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {
	u := User1{"Ankit", "Charles"}
	fmt.Println(u.Greetings())
}
