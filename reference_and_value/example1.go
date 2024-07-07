package main

import "fmt"

type User struct {
	Name  string
	Count int
}

func AddTo(u *User) {
	u.Count++
}

func main() {
	users := make([]User, 0, 4)

	users = append(users, User{Name: "Alice", Count: 0})
	fmt.Printf("%p\n", &users[0])

	users = append(users, User{Name: "Bob", Count: 0})
	fmt.Printf("%p\n", &users[0])

	alice := &users[0]
	amy := User{Name: "Amy", Count: 1}

	users = append(users, amy)

	AddTo(alice)
	fmt.Println(users)
}
