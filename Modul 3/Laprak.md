# <h1 align="center">Laporan Praktikum Modul 3 </h1>
<p align="center">[Hanan Fahri Abiyyu] - [109082500131]</p>

## Unguided

### 1. [Soal Latihan Modul 3]
#### soal1.go

```go
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

```

### 2. [Soal Latihan Modul 2B]
#### soal2.go

```go
package main

import "fmt"

func f(x int) int {
	pangkat := x * x
	return pangkat
}

func g(x int) int {
	fungxie := x - 2
	return fungxie
}

func h(x int) int {
	function := x + 1
	return function
}
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}
```

### 3. [Soal Latihan Modul 2C]
#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}
func main() {
	var (
		cx1, cy1, r1 int
		cx2, cy2, r2 int
		x, y         int
	)
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	titik1 := didalam(float64(cx1), float64(cy1), float64(r1), float64(x), float64(y))
	titik2 := didalam(float64(cx2), float64(cy2), float64(r2), float64(x), float64(y))


	if titik1 && titik2{
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	}else if titik1{
		fmt.Println("Titik di dalam lingkaran 1")
	}else if titik2{
		fmt.Println("Titik di dalam lingkaran 2")
	}else{
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```
### Output Unguided :

##### Output

##### Soal 1

![Screenshoot Output Soal 1](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%202/Output/Output1.png)
Program tersebut berfungsi untuk menerima tiga iput/masukkan dengan masing-masing variabel bertipe data string, yaitu satu, dua, dan tiga. Setelah semua input masuk, program menampilkan urutan awal dari ketiga string tersebut. Selanjutnya program akan melakukan proses pertukaran posisi nilai menggunakan variabel temp, yaitu menyimpan nilai satu ke temp, memindahkan nilai dua ke satu, tiga ke dua, dan nilai yang ada pada temp ke tiga. Proses tersebut membuat urutan data bergeser dari posisi awal. Terakhir, program menampilkan urutan yang menunujkkan posisi ketiga string telah berubah dari urutan awal.

##### Soal 2

![Screenshoot Output Soal 2](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%202/Output/Output2.png)
Membuat program tentang SMA IPA melakukan praktikum kimia di Laboratorium sekolah. Setiap percobaan menggunakan 4 tabung reaksi, dimana susunan warna cairan memengaruhi hasil percobaan. Program dibuat untuk menguji hasil percobaan, siswa diminta menginputkan hasil dari 4 gelas dengan 4 warna berbeda sebanyak 5 kali percobaan. Program akan menampilkan hasil dari eksperimen, jika input warna sesuai urutan yaitu "merah, kuning hijau ungu" selama 5 kali percobaan, maka output atau hasil akan menampilkan true. Jika terdapat satu baris atau satu warna saja tidak sesuai urutan, output akan menampilkan false.

##### Soal 3

![Screenshoot Output Soal 3](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%202/Output/Output3.png)
Membuat program aplikasi perhitungan biaya kirim berdasarkan berat parsel. Dengan satuan awal berupa gram, kemudian akan diambil digit awalan ke dalam bentuk kilogram (kg) dengan cara membagi dengan 1000 dan sisanya sebagai gram dengan cara dimodulus 1000. Selanjutnya akan dicek ke dalam if dengan kondisi awal nilai gram >= 500 dan kg <= 10, jika memenuhi kondisi pertama, maka akan langsung mencetak output dari if tersebut. Selanjutnya terdapat else if pertama dengan kondisi 
gram sebaliknya, yaitu gram < 500 dan kg <= 10. Yang terakhir terdapat else if lagi dengan kondisi kebalikan dari kg awal, yaitu kg > 10, ini kondisi dimana kg lebih dari 10. Saat kondisi ini, walaupun gram lebih atau kurang dari 500, dia tidak akan dikenakan biaya, karena jika kita inputkan angka yang lebih dari 10 kg, maka akan digratiskan biaya sisa beratnya. Seperti pada contoh ketiga, walaupun terdapat sisa berat yaitu 750 gr, tapi karena berat kg nya adalah 11 yang berarti lebih dari 10 kg, biaya hanya terhitung yang 11 kg saja.