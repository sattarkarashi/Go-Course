package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	age                 int
	firstName, lastName string
}

type Person struct {
	age                 int
	firstName, lastName string
}

type JsonExmp struct {
	Id    int    `json:userId`
	It    int    `json:it`
	Title string `json:title`
	Body  string `json:body`
}

func main() {
	user1 := User{age: 10, firstName: "Sato", lastName: "Karo"}
	person1 := Person{age: 10, firstName: "Sato", lastName: "Karo"}

	fmt.Printf("%#v \n", user1)
	//fmt.Printf("%#v \n", user1 == person1) this gives error of mismatched type although they are the same
	person2 := User(person1)
	fmt.Printf("%#v \n", user1 == person2) // but this one doesn't give error because we defined inside the User and now they are not type mismatched.

	// Unknown structs: it implies the implicit conversion and wont have mismatched type error
	var unkown_struct = struct {
		age                 int
		firstName, lastName string
	}{age: 10, firstName: "Sato", lastName: "Karo"}

	var user_unkown_struct = User(unkown_struct)
	fmt.Printf("%#v \n", user_unkown_struct)

	// reading jsons

	res, _ := http.Get("https://jsonplaceholder.typicode.com/posts")
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var posts []JsonExmp
	json.Unmarshal(body, &posts)

	for _, post := range posts {
		fmt.Printf("Post id is %d and post title is %s and user id is %d \n \n", post.Id, post.Title, post.It)
	}

}
