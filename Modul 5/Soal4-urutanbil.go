package main

import "fmt"

func urutan(x int) {
	if x == 1 {
		fmt.Print(1, " ")
		return
	} else {
		fmt.Print(x, " ")
		urutan(x - 1)
		fmt.Print(x, " ")
	}
}
func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	urutan(bilangan)
}
