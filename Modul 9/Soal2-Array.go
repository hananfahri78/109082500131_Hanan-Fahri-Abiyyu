package main

import (
	"fmt"
	"math"
)

func main() {
	var arr [100]int
	var N int
	fmt.Print("Tentukan jumlah array : ")
	fmt.Scan(&N)
	var i int
	// output soal a
	fmt.Print("Indeks : ")
	for i = 0; i < N; i++ {
		fmt.Scan(&arr[i])
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// output soal b
	for i = 0; i < N; i++ {
		if i%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// output soal c
	for i = 0; i < N; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// output soal d
	var x int
	fmt.Print("Tentukan Kelipatan : ")
	fmt.Scan(&x)
	for i = 0; i < N; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// output soal e
	var remove int
	fmt.Print("Tentukan urutan indeks yang akan dihapus : ")
	fmt.Scan(&remove)
	for i := remove; i < N-1; i++ {
		arr[i] = arr[i+1]
	}
	N -= 1
	for i = 0; i < N; i++{
		fmt.Print(arr[i]," ")
	}
	fmt.Println()

	// output soal f
	var total, rata2 float64
	for i = 0; i < N; i++ {
		total = total + float64(arr[i])
	}
	rata2 = total / float64(N)
	fmt.Printf("%.2f\n", rata2)

	// output soal g
	var simpangan_baku, totalS float64
	for i = 0; i < N; i++ {
		totalS += (float64(arr[i]) - rata2) * (float64(arr[i]) - rata2)
	}
	hasil := totalS / float64(N)
	simpangan_baku = math.Sqrt(hasil)
	
	fmt.Printf("%.2f\n", simpangan_baku)

	// output soal h
	fmt.Print("Tentukan nilai frekuensi indeks : ")
	fmt.Scan(&x)
	counter := 0
	for i = 0; i < N; i++ {
		if arr[i] == x {
			counter++
		}
	}
	fmt.Print(counter)

}
