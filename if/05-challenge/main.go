package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func main() {

	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	user, pass := os.Args[1], os.Args[2]

	if user != "jack" {
		fmt.Printf("Access denied for %q.\n", user)
	} else if pass != "1888" {
		fmt.Printf("Invalid password for %q.\n", user)
	} else {
		fmt.Printf("Access granted to %q.\n", user)
	}

}
