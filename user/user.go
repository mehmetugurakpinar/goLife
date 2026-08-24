package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func NewUser(firstName string, lastName string) User {
	return User{FirstName: firstName, LastName: lastName}
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func rename(u User) {
	u.FirstName = "Changed" // only changes the local copy
	fmt.Println(u.FirstName)
}

func renamePtr(u *User) {
	u.FirstName = "Changed" // changes the original
}

func main() {
	u := NewUser("Ugur", "Akpinar")
	fullName := u.FullName()
	fmt.Println(fullName)
	rename(u)
	fmt.Println(u.FirstName) // "Jane" — unchanged
	renamePtr(&u)
	fmt.Println(u.FirstName)
}
