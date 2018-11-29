package structures

import (
	"encoding/json"
	"log"
)

type user struct {
	Name     string `json:"name"`
	Password string `json:"pwd"`
}

type customer struct {
	Location string `json:"location"`
	user
	Password string `json:"pwd"`
}

func Json() {
	// struct => json
	u1 := user{Name: "Tom", Password: "123"}
	json_bytes, err := json.Marshal(u1)
	if err != nil {
		panic(err)
	}
	json_str := string(json_bytes)
	log.Printf("json_str: %s", json_str)

	// json => struct
	u2 := user{}
	err = json.Unmarshal([]byte(json_str), &u2)
	if err != nil {
		panic(err)
	}
	log.Printf("u2: %v", u2)

	// 继承
	c1 := customer{user: user{"Simeon", "pwd123"}, Location: "abcdef", Password: "ABC"}
	log.Printf("c1: %v", c1)

	json_bytes, err = json.Marshal(c1)
	if err != nil {
		panic(err)
	}
	json_str = string(json_bytes)
	log.Printf("json_str: %s", json_str)

	log.Println(c1.Password)
}
