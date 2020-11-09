package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("hello world!")

	kapil := &Person{
		Name: "Kapil",
		Age:  21,
		SocialFollowers: &SocialFollowers{
			Instagram: 680,
			Github:    2,
		},
	}

	data, err := proto.Marshal(kapil)
	if err != nil {
		log.Fatal("Marshalling error: ", err)
	}

	fmt.Println(data)

	newKapil := &Person{}
	err = proto.Unmarshal(data, newKapil)
	if err != nil {
		log.Fatal("unmarshalling error: ", err)
	}
	fmt.Println(newKapil.GetAge())
	fmt.Println(newKapil.GetName())
	fmt.Println(newKapil.SocialFollowers.GetInstagram())
	fmt.Println(newKapil.SocialFollowers.GetGithub())
}
