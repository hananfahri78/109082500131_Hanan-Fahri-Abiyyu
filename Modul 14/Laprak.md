# <h1 align="center">Laporan Praktikum Modul 14 </h1>
<p align="center">[Hanan Fahri Abiyyu] - [109082500131]</p>

## Unguided

### Selection Sort
### 1. [Soal 1 Modul 14]
#### soal1_selection.go

```go
package main

import "fmt"

type banyakRumah [1000000]int

func selectionSort(arr *banyakRumah, m int) {
	var i, j int
	var idx_min int
	i = 0
	for i < m-1 {
		idx_min = i 

		for j = i; j < m; j++{
			if (*arr)[idx_min] > (*arr)[j] {
				idx_min = j
			}
		}
		temp := (*arr)[idx_min]
		(*arr)[idx_min] = (*arr)[i]
		(*arr)[i] = temp
		i++
	}
}

func main() {
	var isiM banyakRumah
	var n, m int
	
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&isiM[j])
		}

		selectionSort(&isiM, m)
		
		for k := 0; k < m; k++{
			fmt.Print(isiM[k], " ")
		}
		fmt.Println()
	}

}

```

### 2. [Soal 2 Modul 14]
#### soal2_selection.go

```go
package main

import "fmt"

type banyakRumah [1000000]int

func selectionSortAsc(arr *banyakRumah, m int) {
	var i, j int
	var idx_min int
	
	i = 0
	for i < m-1 {
		idx_min = i 

		for j = i+1; j < m; j++{
			if (*arr)[idx_min] > (*arr)[j] {
				idx_min = j
			}
		}
		temp := (*arr)[idx_min]
		(*arr)[idx_min] = (*arr)[i]
		(*arr)[i] = temp
		i++
	}
}

func selectionSortDesc(arr *banyakRumah, m int) {
	var i, j int
	var idx_max int

	i = 0
	for i < m-1{
		idx_max = i

		for j = i+1; j < m; j++{
			if (*arr)[idx_max] < (*arr)[j] {
				idx_max = j
			}
		}
		temp := (*arr)[idx_max]
		(*arr)[idx_max] = (*arr)[i]
		(*arr)[i] = temp
		i++
	}
}

func main() {
	var isiM banyakRumah
	var n, m int
	var ganjil, genap banyakRumah
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&isiM[j])
		}

		nGanjil := 0
		nGenap := 0
		for j := 0; j < m; j++{
			if isiM[j]%2 != 0{
				ganjil[nGanjil] = isiM[j]
				nGanjil++
			}else{
				genap[nGenap] = isiM[j]
				nGenap++
			}
		}
		selectionSortAsc(&ganjil, nGanjil)
		selectionSortDesc(&genap, nGenap)
		
		for j := 0; j < nGanjil; j++{
			fmt.Print(ganjil[j], " ")
		}

		for j := 0; j < nGenap; j++{
			fmt.Print(genap[j], " ")
		}
		fmt.Println()
	}

}

```

### Insertion Sort
### 1. [Soal 1 Modul 14]
#### soal1_insertion.go

```go
package main

import "fmt"

type MAX [100]int

func main() {
	var arr MAX
	var bilangan int
	var n int

	fmt.Scan(&bilangan)
	for bilangan >= 0 {
		arr[n] = bilangan
		n++
		fmt.Scan(&bilangan)
	}

	InsertionSort(&arr, n)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
		if i < n-1 {
			fmt.Print(" ")
		}

	}
	fmt.Println()

	Jarak(arr, n)

}
func InsertionSort(arr *MAX, n int) {
	for i := 1; i < n; i++ {
		temp := (*arr)[i]
		j := i - 1

		for j >= 0 && (*arr)[j] > temp {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = temp
	}
}

func Jarak(arr MAX, n int) {
	if n > 1 {
		jarak := arr[1] - arr[0]
		tetap := true

		for i := 1; i < n-1; i++ {
			if arr[i+1]-arr[i] != jarak {
				tetap = false
			}
		}

		if tetap == true {
			fmt.Printf("Data berjarak %d", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}

```

### 2. [Soal 2 Modul 14]
#### soal2_insertion.go

```go
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

```

### 3. [Soal 3 Modul 14]
#### soal3_insertion.go

```go
package main

import "fmt"

type arrMax [1000000]int

func insertionSort(arr *arrMax, n int) {
	for i := 1; i < n; i++ {
		temp := (*arr)[i]
		j := i - 1

		for j >= 0 && (*arr)[j] > temp {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = temp
	}
}
func main() {
	var data arrMax
	var bil int
	var n int

	fmt.Scan(&bil)
	for bil != -5313 {
		if bil == 0 {
			if n > 0 {
				insertionSort(&data, n)
				if n%2 != 0 {
					fmt.Println(data[n/2])
				} else {
					fmt.Println((data[n/2-1] + data[n/2]) / 2)
				}
			}
		} else {
			data[n] = bil
			n++
		}
		fmt.Scan(&bil)
	}
}
```
### Output Unguided :

##### Output

##### Selection Sort
##### Soal 1

![Screenshoot Output Soal 1](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%2014/Selection%20Sort/Output%20Soal/Output-soal1.png)
Program rumahKerabat digunakan untuk membantu Hercules dalam mengurutkan data nomor rumah kerabat nya pada beberapa daerah menggunakan metode Selection Sort.

Masukan diawali dengan sebuah bilangan bulat n yang menyatakan banyaknya daerah yang akan diproses. Untuk setiap daerah, pengguna memasukkan sebuah bilangan bulat m yang menyatakan banyaknya rumah pada daerah tersebut, kemudian diikuti oleh m buah bilangan bulat yang merepresentasikan nomor rumah pada daerah tersebut.

Dalam program terdapat function tambahan yaitu selectionSort, yang digunakan untuk mengurutkan data nomor rumah pada suatu daerah secara menaik (ascending). Metode yang digunakan adalah Selection Sort, yaitu dengan mencari elemen terkecil pada bagian array yang belum terurut, kemudian menukarkannya dengan elemen pada posisi yang seharusnya.

Keluaran berupa beberapa baris data. Setiap baris menampilkan nomor rumah yang telah terurut dari nilai terkecil ke nilai terbesar untuk masing-masing daerah sesuai urutan masukan.

##### Soal 2

![Screenshoot Output Soal 2](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%2014/Selection%20Sort/Output%20Soal/Output-soal2.png)
Program ini digunakan untuk mengolah data nomor rumah pada beberapa daerah berdasarkan sifat bilangan ganjil dan genap. Pengguna terlebih dahulu memasukkan jumlah daerah yang akan diproses. Untuk setiap daerah, pengguna memasukkan banyaknya rumah beserta nomor rumah yang ada pada daerah tersebut.

Setelah seluruh data rumah dibaca, program akan memisahkan nomor rumah ke dalam dua kelompok, yaitu kelompok bilangan ganjil dan kelompok bilangan genap. Bilangan ganjil kemudian diurutkan secara menaik (dari kecil ke besar) menggunakan prosedur selectionSortAsc, sedangkan bilangan genap diurutkan secara menurun (dari besar ke kecil) menggunakan prosedur selectionSortDesc.

Kedua prosedur tersebut menggunakan metode Selection Sort, yaitu dengan mencari nilai terkecil atau terbesar pada bagian data yang belum terurut, kemudian menukarkannya ke posisi yang sesuai.

Keluaran program berupa daftar nomor rumah yang telah diurutkan dengan ketentuan seluruh bilangan ganjil ditampilkan terlebih dahulu secara menaik, kemudian diikuti oleh seluruh bilangan genap secara menurun. Hasil untuk setiap daerah ditampilkan pada baris yang berbeda.

##### Insertion Sort
##### Soal 1

![Screenshoot Output Soal 3](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%2014/Insertion%20Sort/Output/Output_soal1.png)
Membuat program mengurutkan sekumpulan bilangan yang dimasukkan oleh pengguna, kemudian memeriksa apakah data tersebut memiliki jarak yang tetap antar nilainya. Pengguna dapat memasukkan beberapa bilangan bulat positif atau nol secara berurutan. Proses masukan akan berhenti ketika pengguna memasukkan bilangan negatif.

Dalam program terdapat prosedur InsertionSort yang digunakan untuk mengurutkan seluruh data secara menaik (ascending) menggunakan metode Insertion Sort. Metode ini bekerja dengan mengambil satu elemen, kemudian menyisipkannya ke posisi yang sesuai pada bagian data yang telah terurut sebelumnya hingga seluruh data berada dalam urutan yang benar.

Setelah data berhasil diurutkan, program akan menampilkan seluruh bilangan dalam keadaan terurut. Selanjutnya, prosedur Jarak digunakan untuk memeriksa apakah selisih antara setiap pasangan bilangan yang berurutan memiliki nilai yang sama. Jika seluruh selisih bernilai sama, program akan menampilkan besar jarak tersebut. Sebaliknya, jika terdapat selisih yang berbeda, program akan menampilkan informasi bahwa data memiliki jarak yang tidak tetap.

Keluaran program berupa satu baris data yang telah terurut dari kecil ke besar, kemudian diikuti informasi mengenai jarak antar data, apakah tetap atau tidak tetap. Jika data berjarak tetap maka akan menuliskan berupa "Data berjarak (banyaknya jarak)". Sebaliknya, jika tidak, maka akan mengeluarkan output "Data berjarak tidak tetap". 

##### Soal 2

![Screenshoot Output Soal 4](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%2014/Insertion%20Sort/Output/Output_soal2.png)
Program ini digunakan untuk mengelola data buku dalam sebuah pustaka berdasarkan rating yang dimiliki setiap buku. Pengguna terlebih dahulu memasukkan jumlah buku yang akan didata, kemudian memasukkan informasi setiap buku yang terdiri atas ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating, yang disimpan di dalam struct Buku. Setelah seluruh data buku dimasukkan, pengguna juga memasukkan sebuah rating yang akan digunakan sebagai kunci pencarian.

Dalam program terdapat beberapa prosedur pendukung. Prosedur DaftarkanBuku digunakan untuk menyimpan seluruh data buku yang dimasukkan pengguna ke dalam array. Prosedur CetakTerfavorit digunakan untuk mencari dan menampilkan buku dengan rating tertinggi. Selanjutnya, prosedur UrutBuku mengurutkan seluruh data buku berdasarkan rating secara menurun menggunakan metode Insertion Sort, sehingga buku dengan rating tertinggi berada di posisi awal.

Setelah data terurut, prosedur Cetak5Terbaru menampilkan lima judul buku dengan rating tertinggi, atau seluruh judul buku jika jumlah data kurang dari lima. Program juga memiliki prosedur CariBuku yang menggunakan metode Binary Search untuk mencari salah satu buku dengan rating tertentu sesuai masukan pengguna.

Keluaran program terdiri atas beberapa bagian, yaitu data buku terfavorit, lima judul buku dengan rating tertinggi, serta informasi salah satu buku yang memiliki rating sesuai dengan nilai yang dicari. Jika tidak ditemukan buku dengan rating tersebut, program akan menampilkan pesan bahwa buku tidak ditemukan


##### Soal 3

![Screenshoot Output Soal 5](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%2014/Insertion%20Sort/Output/Output_soal3.png)
Program ini dibuat berdasarkan sebuah permasalahan yang muncul pada kompetisi pemrograman yang diikuti oleh berbagai tim dari perguruan tinggi. Meskipun problem yang diberikan terlihat sederhana dan hampir seluruh tim mencoba menyelesaikannya, hanya sedikit tim yang berhasil memperoleh jawaban yang benar. Tantangan utama pada problem ini adalah mengolah data yang terus bertambah dan menentukan nilai median dari kumpulan data tersebut secara tepat.

Untuk menyelesaikan permasalahan tersebut, program menerima sejumlah bilangan yang dimasukkan secara bertahap. Setiap kali pengguna memasukkan nilai 0, program akan mengurutkan data yang telah terkumpul menggunakan metode Insertion Sort, kemudian menentukan nilai median dari data tersebut. Proses berlanjut hingga pengguna memasukkan nilai khusus sebagai penanda akhir masukan.

Keluaran program berupa nilai median dari data yang telah dimasukkan pada setiap permintaan perhitungan median