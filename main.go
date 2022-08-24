package main

import (
	"fmt"

	"github.com/denlam/testgolang1/utilities"
)

func main() {

	s := "Hello"

	s = utilities.Strlower(s)

	count := 3

	for i := 0; i < count; i++ {
		fmt.Printf("%s %d\n", s, i+1)
	}

}
