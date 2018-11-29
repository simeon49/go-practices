package main

import (
	"io/ioutil"
	"log"

	"./pb"
	"github.com/golang/protobuf/proto"
)

func main() {
	in, err := ioutil.ReadFile("book.data")
	if err != nil {
		log.Fatalln("Failed to read address book:", err)
	}
	var book *pb.AddressBook
	book = &pb.AddressBook{}
	err = proto.Unmarshal(in, book)
	if err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	log.Printf("adress: %v", book)
}
