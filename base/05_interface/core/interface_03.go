package core

type Admin struct {
	User
	level int
	name  string
}

func Interface03() {
	u := Admin{
		User: User{
			name:  "Simeon",
			email: "qixiyi79@gmail.com",
		},
		level: 999,
	}
	SendNotifiction(&u)
}
