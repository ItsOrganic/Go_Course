package main

import (
	"fmt"
	"github.com/wesovilabs/koazee/stream"
)

// User data type contains all attributes about an user
type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	var users = [3]User{
		User{
			id:      1,
			name:    "Pepek",
			surname: "Vyskoč"},
		User{
			id:      2,
			name:    "Pepek",
			surname: "Vyskoč"},
		User{
			id:      3,
			name:    "Josef",
			surname: "Vyskočil"},
	}
	fmt.Println(users)

	fmt.Printf("input:  %v\n", users)

	newStream := stream.New(users)
	fmt.Printf("stream: %v\n", newStream.Out().Val())
}
