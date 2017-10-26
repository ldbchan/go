package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// c: {1, 2, 0, 0, 0}
	c := [5]int{1, 2}
	fmt.Println("c:  ", c)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// declare and init a two-dimensional array
	d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("d:  ", d)
}
