package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Print(i)
		fmt.Print(sep)
		fmt.Println(arg)
	}
	fmt.Println(s)
}