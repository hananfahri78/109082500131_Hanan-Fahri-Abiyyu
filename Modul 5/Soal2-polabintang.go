package main

import "fmt"

func bintang(n int) string {
	if n == 0 {
		return ""
	}
	bintang(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()

	return ""
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(bintang(n))
}
