package main

import "fmt"

func faktor(x int, i int) {
	if i > x {
		return
	} else {
		if x%i == 0 {
			fmt.Print(i, " ")
		}
		faktor(x, i+1)
	}
}
func main() {
	var bil int
	fmt.Scan(&bil)
	faktor(bil, 1)
}
