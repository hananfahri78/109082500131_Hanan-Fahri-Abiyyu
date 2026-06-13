package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n > 0 {
		maxRating := pustaka[0].rating
		for i := 1; i < n; i++ {
			if pustaka[i].rating > maxRating {
				maxRating = pustaka[i].rating
			}
		}

		for i := 0; i < n; i++ {
			if pustaka[i].rating == maxRating {
				fmt.Printf("%s %s %s %d\n", pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].tahun)
			}
		}
	}
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		temp := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < temp.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = temp

	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	Book := 5
	if n < 5 {
		Book = n
	}
	for i := 0; i < Book; i++{
		fmt.Println(pustaka[i].judul)
	} 
} 

func CariBuku(pustaka DaftarBuku, n int, r int) {
	found := -1
	left := 0
	right := n-1
	
	for left <= right && found == -1 {
		med := (left + right) / 2
		if pustaka[med].rating == r {
			found = med
		} else if pustaka[med].rating < r {
			right = med -1
		} else {
			left = med + 1
		}
	}

	if found != -1 {
		fmt.Printf("%s %s %s %d %d %d\n", pustaka[found].judul, pustaka[found].penulis, pustaka[found].penerbit, pustaka[found].eksemplar, pustaka[found].tahun, pustaka[found].rating)
	}else{
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var Pustaka DaftarBuku
	var nPustaka int

	fmt.Scan(&nPustaka)
	DaftarkanBuku(&Pustaka, nPustaka)

	var targetRating int
	fmt.Scan(&targetRating)

	CetakTerfavorit(Pustaka, nPustaka)
	UrutBuku(&Pustaka, nPustaka)
	Cetak5Terbaru(Pustaka, nPustaka)
	CariBuku(Pustaka, nPustaka, targetRating)
}
