package interfaces

import (
	"log"
)

type User struct {
	name  string
	email string
}

func (u *User) Notify() error {
	log.Println(u)
	log.Printf("notify %s:%s\n", u.name, u.email)
	return nil
}

type Notifier interface {
	Notify() error
}

func SendNotifiction(notifier Notifier) error {
	return notifier.Notify()
}

func Interface02() {
	u := User{
		name:  "Simeon",
		email: "qixiyi@yeah.net",
	}
	u.Notify()

	// SendNotifiction(u) // 这个方法会报错!
	SendNotifiction(&u)
}
