package main

import "fmt"

func main() {

	// output
	// 1
	// 2
	// 3
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1 // `i += 1` is also valid
	}

	// output
	// 7
	// 8
	// 9
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// output
	// loop
	for {
		fmt.Println("loop")
		break
	}

	// output
	// 1
	// 3
	// 5
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
