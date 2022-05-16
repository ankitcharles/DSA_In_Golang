package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
	User
	GameId int
}

func main() {
	p := Player{}
	p.Id = 58
	p.Name = "Ankit"
	p.Location = "Bangalore"
	fmt.Println(p.Greetings())
	a := [...]string{"Hello", "Golang!"}
	fmt.Printf("%q", a)
	cities := make([]string, 3)
	cities[0] = "Bangalore"
	cities[1] = "Hyderabad"
	cities[2] = "Chennai"
	fmt.Printf("%q", cities)
}
