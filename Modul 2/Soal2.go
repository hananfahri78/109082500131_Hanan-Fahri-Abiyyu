package main

import "fmt"

func main() {
	var (
		warna1, warna2, warna3, warna4 string
		temp bool
	)
	fmt.Println("Urutkan warna sesuai rules")
	temp = true
	i := 1
	for i <= 5{
		fmt.Print("Percobaan ", i)
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)
		if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu"{
			temp = false
		}
		i++ 
	}
	fmt.Println("Berhasil :", temp)
}
