package main

import "fmt"

func main() {

	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// int var is initialized as 0
	var e int
	fmt.Println(e)

	// shorthand for 'var f string = "short"'
	f := "short"
	fmt.Println(f)

	var g string = "short"
	fmt.Println(g)
}
