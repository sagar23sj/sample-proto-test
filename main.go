package main

import (
	"fmt"
	"io/ioutil"
	"log"

	sample "github.com/sagar23sj/proto-test/protos"
	"google.golang.org/protobuf/proto"
)

func main() {
	// val := false
	aa := sample.BenchLarge{
		Name:   "John",
		Age:    10,
		Height: 160,
	}

	fmt.Println("My name is :", aa.GetName(), aa.GetAge(), aa.GetHeight())

	out, err := proto.Marshal(&aa)
	// fmt.Printf("Binarry : % b", out)
	// fmt.Printf("string : % s", out)
	// fmt.Printf("Hex : % x", out)
	if err != nil {
		log.Fatalln("Failed to encode:", err)
	}
	if err := ioutil.WriteFile("test.txt", out, 0644); err != nil {
		log.Fatalln("Failed to:", err)
	}

	// Read the existing address book.
	in, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	bb := &sample.BenchLarge{}
	if err := proto.Unmarshal(in, bb); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println("My name is :", bb.GetName(), bb.GetAge(), bb.GetHeight())
}
