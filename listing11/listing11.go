package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@gmail.com"}
	bill.notify()

	lisa := &user{"Lisa", "lisa@gmail.com"}
	lisa.notify()

	bill.changeEmail("bill@yahoo.com")
	bill.notify()

	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
