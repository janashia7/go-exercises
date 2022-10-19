package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1]

	ft, err := strconv.ParseFloat(args, 64)

	m := ft * 0.3048

	if err != nil {
		fmt.Printf("error : %q is not a number\n", args)
		return
	}
	fmt.Printf("%g feet is %g meters.\n", ft, m)
}
