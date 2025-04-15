package main

import (
	"fmt"
)

func main() {
	s := "vinícius"

	fmt.Printf("%8T %[1]v %d\n", s, len(s))
	fmt.Printf("%8T %[1]v\n", []rune(s))

	b := []byte(s)
	fmt.Printf("%8T %[1]v %d\n", b, len(b))
}
