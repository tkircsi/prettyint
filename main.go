package main

import (
	"fmt"

	"github.com/tkircsi/prettyint/prettyint"
)

func main() {
	var i prettyint.Int = -100000

	fmt.Println(i)
	fmt.Printf("%%d: %d\n", i)
	fmt.Printf("%%v: %v\n", i)
	fmt.Printf("%%s: %s\n", i)

	i = 99456
	fmt.Println(i)
	fmt.Printf("%%d: %d\n", i)
	fmt.Printf("%%v: %v\n", i)
	fmt.Printf("%%s: %s\n", i)

}
