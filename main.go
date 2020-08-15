package main

import (
    "fmt"
    "log"
    person "proto_go_ex/person"
    proto "github.com/golang/protobuf/proto"
)

func main(){
    fmt.Println("Hello world!")

    john := &person.Person{
        Name: "John",
        Age: 24,
        SocialFollowers: &person.SocialFollowers {
            Youtube: 1,
            Facebook: 5,
        },
    }

    // Marshal data
    data, err := proto.Marshal(john)
    if err != nil {
        log.Fatalf("Marshalling error: %v\n", err)
    }

    fmt.Println(data)

    // Unmarshal data
    newPerson := &person.Person{}
    err = proto.Unmarshal(data, newPerson)
    if err != nil {
        log.Fatal("Unmarshalling error:", err)
    }

    fmt.Println(newPerson)
}
