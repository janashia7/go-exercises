package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {

	names := os.Args[1:]
	count := len(names)

	if count > 3 {
		fmt.Printf("Error: Expected 3 argument and got %d\n", count)
		return
	}

	fmt.Printf("There are %v people!\n", count)
	fmt.Println("Hello great", names[0], "!")
	fmt.Println("Hello great", names[1], "!")
	fmt.Println("Hello great", names[2], "!")
	fmt.Println("Nice to meet you all.")

	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	// Search on the web how to solve that.
}
