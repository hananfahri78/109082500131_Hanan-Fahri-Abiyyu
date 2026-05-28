package main

import "fmt"

const arrMax int = 1000

type arr [arrMax]float64

func main() {
	var jumlahKelinci int
	var idxmax, idxmin int
	var berat_kelinci arr
	fmt.Print("Masukkan jumlah kelinci : ")
	fmt.Scan(&jumlahKelinci)

	for i := 0; i < jumlahKelinci; i++ {
		fmt.Scan(&berat_kelinci[i])
	}

	idxmax = 0
	for i := 1; i < jumlahKelinci; i++ {
		if berat_kelinci[i] > berat_kelinci[idxmax] {
			idxmax = i
		}
	}
	
	idxmin = 0
	for i := 1; i < jumlahKelinci; i++ {
		if berat_kelinci[i] < berat_kelinci[idxmin] {
			idxmin = i
		}
	}

	fmt.Println("Berat kelinci terbesar : ", berat_kelinci[idxmax])
	fmt.Println("Berat kelinci terkecil : ", berat_kelinci[idxmin])
}
