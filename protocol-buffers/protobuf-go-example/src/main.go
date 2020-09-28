package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
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

	jsonString, _ := toJsonString(bookMessage)
	fmt.Println(jsonString)

	bookMessageFromJsonString, _ := fromJsonString(jsonString)
	fmt.Println(bookMessageFromJsonString)
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

func toJsonString(message proto.Message) (string, error) {
	marshaller := jsonpb.Marshaler{}
	json, err := marshaller.MarshalToString(message)
	if err != nil {
		log.Fatalln("Cannot marshal to JSON", err)
		return "", err
	}
	return json, nil
}

func fromJsonString(json string) (proto.Message, error) {
	bookMessage := &bookpb.BookMessage{}
	err := jsonpb.UnmarshalString(json, bookMessage)
	if err != nil {
		log.Fatalln("Cannot unmarshal to message", err)
		return nil, err
	}
	return bookMessage, nil
}
