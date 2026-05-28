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

![Screenshoot Output Soal 1](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%2010/Output/Output_soal1.png)
Program tersebut berfungsi untuk mendata berat anak kelinci yang akan dijual ke pasar. Program menggunakan array berkapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.

Masukan terdiri dari sekumpulan bilangan, bilangan pertama adalah variabel bilangan bulat "jumlahKelinci" yang digunakan untuk menyatakan banyaknya anak kelinci yang akan ditimbang beratnya, atau untuk sebagai batas banyaknya array yang akan diisikan. Selanjutnya variabel "berat_kelinci" dengan tipe data array float64, yang disimpan di dalam sebuah array, digunakan untuk menyatakan berat dari anak kelinci yang akan dijual. Serta variabel "idxMin" dan "idxMax", yang digunakan untuk mencari nilai maksimum dan minimum dari data berat anak kelinci.

Keluaran berupa dua bilangan real yang menyatakan berat kelinci terkecil dan terbesar.



##### Soal 2

![Screenshoot Output Soal 2](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%2010/Output/Output_soal2.png)
Program tersebut berfungsi untuk menentukan total berat ikan pada setiap wadah serta menghitung rata-rata berat ikan di tiap wadah. Program menggunakan array berkapasitas 1000 untuk menampung data berat ikan yang akan dijual.

Masukan terdiri dari dua bilangan bulat yaitu variabel x dan y. Variabel x digunakan untuk menyatakan banyaknya ikan yang akan dijual atau banyaknya data yang akan diinputkan ke dalam array. Sedangkan variabel y digunakan untuk menyatakan kapasitas jumlah ikan di setiap wadah. Selanjutnya variabel ikan dengan tipe data float64/real, yang disimpan di dalam sebuah array, digunakan untuk menyimpan berat masing-masing ikan yang akan dijual.

Program juga menggunakan variabel "totalPerWadah" bertipe array float64 untuk menyimpan total berat ikan pada setiap wadah. Variabel idxWadah digunakan untuk menentukan indeks wadah yang sedang diisi. Kemudian variabel jmlWadah digunakan untuk menghitung jumlah wadah yang dipakai berdasarkan nilai x dan y. Selain itu, variabel total digunakan untuk menjumlahkan seluruh total berat ikan dari setiap wadah, sedangkan variabel rata2 digunakan untuk menghitung rata-rata berat ikan di tiap wadah.

Keluaran terdiri dari dua baris. Baris pertama berupa kumpulan bilangan real yang menyatakan total berat ikan pada setiap wadah sesuai urutan pengisian. Baris kedua berupa sebuah bilangan real yang menyatakan rata-rata berat ikan di tiap wadah.

##### Soal 3

![Screenshoot Output Soal 3](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%2010/Output/Output_soal3.png)
Membuat program pencatatan data berat balita yang dilakukan oleh Pos Pelayanan Terpadu (posyandu). Petugas posyandu akan memasukkan data balita ke dalam array, dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan rerata nya.

Array ditentukan berkapasitas 100.

Terdapat dua function tambahan, yaitu hitungMinMax yang berfungsi untuk mencari nilai minimum dan nilai maksimum, serta function rerata yang digunakan untuk menghitung rata-rata dari jumlah berat balita dibagi banyaknya data balita.

Masukan pertama berupa banyaknya balita yang akan didata. Variabel "dataBalita" bertipe data integer dan dijadikan sebagai variabel global (di luar function), sehingga dapat digunakan di semua function tanpa pemanggilan parameter tambahan. Selanjutnya variabel "beratBalita" bertipe data array float64, digunakan untuk menyatakan berat dari banyaknya balita. Serta "rataRata" dengan tipe data float64, digunakan sebagai tempat penyimpanan nilai yang diambil dari function rerata.

Keluaran berupa bilangan desimal dari penentuan berat balita minimum (data terkecil), berat balita maksimum (data terbesar), dan rerata berat balita.