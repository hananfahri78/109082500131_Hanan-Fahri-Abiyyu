package main

import "fmt"

func ganjil(x int, i int) {
	if i > x {
		return
	} else {

		fmt.Print(i, " ")
		ganjil(x, i+2)

	}
}
func main() {
	var angka int
	fmt.Scan(&angka)
	ganjil(angka, 1)
}
