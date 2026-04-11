package main

import "fmt"

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}
func main() {
	var angka int
	fmt.Scan(&angka)
	for i := 0; i <= angka; i++ {
		fmt.Print(fibo(i), " ")
	}
}
