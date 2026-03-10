# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Hanan Fahri Abiyyu] - [109082500131]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```

#### soal2.go

```go
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
```

#### soal3.go

```go
package main

import "fmt"

func main() {
	var  g, kg, tambahan int
	fmt.Print("Berat parsel (gram) :")
	fmt.Scan(&g)
	kg = g/1000
	gr := g % 1000
	fmt.Println("Detail berat :", kg, "kg +", gr, "gr")
	kg2 := kg * 10000
	if gr >= 500 && kg <= 10 {
		tambahan = (gr * 5)
		fmt.Println("Detail biaya :", "Rp.", kg2, "+ Rp.", tambahan)
		total := tambahan + kg2
		fmt.Println("Total biaya :", total)
	}else if gr < 500 && kg <= 10 {
		tambahan = (gr * 15)
		fmt.Println("Detail biaya :", "Rp.", kg2, "+ Rp.", tambahan)
		kg2 = tambahan + kg2
		fmt.Print("Total biaya :", kg2)
	}else if kg > 10{
		tambahan = (gr * 5)
		fmt.Println("Detail biaya : ", "Rp.", kg2, "+", "Rp.", tambahan)
		fmt.Println("Total biaya :", kg2)
	}
	
}
```
### Output Unguided :

##### Output
![Screenshoot Output Soal 2](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%201/Output/output_soal2.png)
[penjelasan]