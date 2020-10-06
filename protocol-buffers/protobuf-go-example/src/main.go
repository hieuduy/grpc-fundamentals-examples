package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"protobuf-go-example/src/book"
)

const (
	fileName = "book_message.bin"
)

func main() {
	bookMessage := &bookpb.BookMessage{
		Id:     1,
		Title:  "Java 8",
		Author: "Oracle",
		Tags:   []string{"Development", "Java"},
	}

	fmt.Println(bookMessage)
	_ = writeToFile(bookMessage, fileName)

	bookMessageFromFile := &bookpb.BookMessage{}
	_ = readFromFile(bookMessageFromFile, fileName)
	fmt.Println(bookMessageFromFile)
}

func writeToFile(message proto.Message, fileName string) error {
	bytes, err := proto.Marshal(message)
	if err != nil {
		log.Fatalln("Cannot marshal to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fileName, bytes, 0644); err != nil {
		log.Fatalln("Cannot write to file", err)
		return err
	}

	return nil
}

func readFromFile(message proto.Message, fileName string) error {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Cannot read from file", err)
		return err
	}

	err = proto.Unmarshal(bytes, message)
	if err != nil {
		log.Fatalln("Cannot unmarshal to message", err)
		return err
	}

	return nil
}
