package main

import (
	"fmt"

	f "github.com/fauna/faunadb-go/v4/faunadb"
)

type User struct {
	Name string `fauna:"name"`
}

func mainFauna() {
	client := f.NewFaunaClient("your-secret-here")

	res, err := client.Query(f.Get(f.Ref(f.Collection("user"), "42")))
	if err != nil {
		panic(err)
	}

	var user User

	if err := res.At(f.ObjKey("data")).Get(&user); err != nil {
		panic(err)
	}

	fmt.Println(user)
}