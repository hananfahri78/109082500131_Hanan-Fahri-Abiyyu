package main

import "fmt"

func faktorial(n int) int {
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if a >= c {
		fmt.Print(permutasi(a, c), " ")
		fmt.Println(kombinasi(a, c))
	}

	if b >= d {
		fmt.Print(permutasi(b, d), " ")
		fmt.Println(kombinasi(b, d))
	}
}
