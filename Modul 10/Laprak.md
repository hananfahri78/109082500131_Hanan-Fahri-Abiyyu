# <h1 align="center">Laporan Praktikum Modul 10 </h1>
<p align="center">[Hanan Fahri Abiyyu] - [109082500131]</p>

## Unguided

### 1. [Soal Latihan Modul 10(2)]
#### soal1.go

```go
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
```

### 2. [Soal Latihan Modul 10(2)]
#### soal2.go

```go
package main
import "fmt"

type arrayMax [1000]float64

func main() {
	var x, y int
	var ikan arrayMax
	var totalPerWadah arrayMax
	var total, rata2 float64	

	fmt.Scan(&x, &y)
	for i := 0; i < x; i++{
		fmt.Scan(&ikan[i])
	}
	idxWadah := 0
	for i := 0; i < x; i++{
		totalPerWadah[idxWadah] += ikan[i]

		if (i+1)%y == 0{
			idxWadah++
		}
	} 

	jmlWadah := x/y
	if x%y != 0{
		jmlWadah++
	}

	for i := 0; i < jmlWadah; i++{
		fmt.Print(totalPerWadah[i], " ")
		total += totalPerWadah[i]
	}

	fmt.Println()

	rata2 = total / float64(jmlWadah)
	fmt.Println("Berat rata-rata ikan di tiap wadah : ",rata2)
}
```

### 3. [Soal Latihan Modul 10(3)]
#### soal3.go

```go
package main
import "fmt"
type arrBalita [100]float64
var dataBalita int

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < dataBalita; i++ {
		if arrBerat[i] < *bMin{
			*bMin = arrBerat[i]
		}

		if arrBerat[i] > *bMax{
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita) float64 {
	var rerata float64
	var total float64 = 0

	for i := 0; i < dataBalita; i++ {
		total += arrBerat[i]
	}
	rerata = total / float64(dataBalita)
	return rerata
}

func main() {

	var beratBalita arrBalita
	var rataRata float64
	var bMin, bMax float64
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&dataBalita)

	for i := 0; i < dataBalita; i++ {
		fmt.Print("Masukan berat balita ke-", i+1, " : ")
		fmt.Scan(&beratBalita[i])
	}

	hitungMinMax(beratBalita, &bMin, &bMax)
	rataRata = rerata(beratBalita)
	fmt.Printf("Berat balita minimum  : %.2f\n", bMin)
	fmt.Printf("Berat balita maksimum : %.2f\n", bMax)
	fmt.Printf("Rerata berat balita   : %.2f\n", rataRata)
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