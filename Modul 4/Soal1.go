package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	var i int
	for i = 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int, hasil *int) {
	var fn, fnr int
	faktorial(n, &fn)
	faktorial(n-r, &fnr)
	*hasil = fn / fnr
}

func kombinasi(n, r int, hasil *int) {
	var fn, fnr, fr int
	faktorial(n, &fn)
	faktorial(r, &fr)
	faktorial(n-r, &fnr)
	*hasil = fn / (fr * (fnr))
}

func main() {
	var a, b, c, d, hasil int
	fmt.Scan(&a, &b, &c, &d)
	if a >= c {
		permutasi(a, c, &hasil)
		fmt.Print(hasil, " ")
		kombinasi(a, c, &hasil)
		fmt.Println(hasil)
	}

	if b >= d {
		permutasi(b, d, &hasil)
		fmt.Print(hasil, " ")
		kombinasi(b, d, &hasil)
		fmt.Println(hasil)
	}

}
