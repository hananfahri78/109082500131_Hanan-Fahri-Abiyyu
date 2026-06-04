package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	isiArray(n)

	idx := posisi(n, k)
	if idx != -1 {
		fmt.Println(idx)
	} else {
		fmt.Println("TIDAK ADA")
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	var found int = -1
	var med int
	var kiri int = 0
	var kanan int = n - 1

	for kiri <= kanan && found == -1 {
		med = (kiri + kanan) / 2
		if k < data[med] {
			kanan = med - 1
		} else if k > data[med] {
			kiri = med + 1
		} else {
			found = med
		}
	}
	
	return found
}