package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	ctx := context.Background()
	timeout := 30 * time.Second

	reqContext, _ := context.WithTimeout(ctx, timeout)
	m, err := Get[RequestObj](reqContext, "https://reqres.in/api/users?page=2")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m.Data[0])


	// Post data
	user := User{ Name : "morpheus", Job : "leader"}
	addContext, _ := context.WithTimeout(ctx,  timeout) 

	newUser,_ := Post[User](addContext, "https://reqres.in/api/users", user)

	fmt.Println( newUser )

}
