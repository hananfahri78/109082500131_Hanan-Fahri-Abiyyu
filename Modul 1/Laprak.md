# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Hanan Fahri Abiyyu] - [109082500131]</p>

## Unguided

### 1. [Soal Latihan Modul 2A]
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

### 2. [Soal Latihan Modul 2B]
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

### 3. [Soal Latihan Modul 2C]
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

##### Soal 1

![Screenshoot Output Soal 1](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%201/Output/output_soal1.png)
Program tersebut berfungsi untuk menerima tiga iput/masukkan dengan masing-masing variabel bertipe data string, yaitu satu, dua, dan tiga. Setelah semua input masuk, program menampilkan urutan awal dari ketiga string tersebut. Selanjutnya program akan melakukan proses pertukaran posisi nilai menggunakan variabel temp, yaitu menyimpan nilai satu ke temp, memindahkan nilai dua ke satu, tiga ke dua, dan nilai yang ada pada temp ke tiga. Proses tersebut membuat urutan data bergeser dari posisi awal. Terakhir, program menampilkan urutan yang menunujkkan posisi ketiga string telah berubah dari urutan awal.

##### Soal 2

![Screenshoot Output Soal 2](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%201/Output/output_soal2.png)
Membuat program tentang SMA IPA melakukan praktikum kimia di Laboratorium sekolah. Setiap percobaan menggunakan 4 tabung reaksi, dimana susunan warna cairan memengaruhi hasil percobaan. Program dibuat untuk menguji hasil percobaan, siswa diminta menginputkan hasil dari 4 gelas dengan 4 warna berbeda sebanyak 5 kali percobaan. Program akan menampilkan hasil dari eksperimen, jika input warna sesuai urutan yaitu "merah, kuning hijau ungu" selama 5 kali percobaan, maka output atau hasil akan menampilkan true. Jika terdapat satu baris atau satu warna saja tidak sesuai urutan, output akan menampilkan false.

##### Soal 3

![Screenshoot Output Soal 3](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%201/Output/output_soal3.png)
Membuat program aplikasi perhitungan biaya kirim berdasarkan berat parsel. Dengan satuan awal berupa gram, kemudian akan diambil digit awalan ke dalam bentuk kilogram (kg) dengan cara membagi dengan 1000 dan sisanya sebagai gram dengan cara dimodulus 1000. Selanjutnya akan dicek ke dalam if dengan kondisi awal nilai gram >= 500 dan kg <= 10, jika memenuhi kondisi pertama, maka akan langsung mencetak output dari if tersebut. Selanjutnya terdapat else if pertama dengan kondisi 
gram sebaliknya, yaitu gram < 500 dan kg <= 10. Yang terakhir terdapat else if lagi dengan kondisi kebalikan dari kg awal, yaitu kg > 10, ini kondisi dimana kg lebih dari 10. Saat kondisi ini, walaupun gram lebih atau kurang dari 500, dia tidak akan dikenakan biaya, karena jika kita inputkan angka yang lebih dari 10 kg, maka akan digratiskan biaya sisa beratnya. Seperti pada contoh ketiga, walaupun terdapat sisa berat yaitu 750 gr, tapi karena berat kg nya adalah 11 yang berarti lebih dari 10 kg, biaya hanya terhitung yang 11 kg saja.