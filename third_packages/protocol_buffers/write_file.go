package main

import (
	"io/ioutil"
	"log"

	"./pb"
	"github.com/golang/protobuf/proto"
)

func main() {
	p1 := pb.Person{
		Id:    1234,
		Name:  "Simeon",
		Email: "qixiyi@yeah.net",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "13644099810", Type: pb.Person_MOBILE},
			{Number: "0411-84620027", Type: pb.Person_HOME},
		},
	}
	book := &pb.AddressBook{People: []*pb.Person{&p1}}

	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	err = ioutil.WriteFile("./book.data", out, 0644)
	if err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
