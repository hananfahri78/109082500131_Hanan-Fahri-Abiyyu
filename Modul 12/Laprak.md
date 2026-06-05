# <h1 align="center">Laporan Praktikum Modul 12 </h1>
<p align="center">[Hanan Fahri Abiyyu] - [109082500131]</p>

## Unguided

### 1. [Soal Latihan Modul 12 (1)]
#### soal1.go

```go
package main

import "fmt"

const Max int = 20

type arrKandidat [Max]int

func sequentialSearch(T arrKandidat, n int, X int) int {
	var found int = -1
	var i int

	i = 0
	for i < n && found == -1{
		if T[i] == X{
			found = i
		}
		i++
	}
	return found
}

func main() {
	var kandidat arrKandidat
	var totalSuara [Max]int
	var suara, suaraMasuk, suaraSah int
	var i int

	nomor := 1

	for i = 0; i < Max; i++ {
		kandidat[i] = nomor
		nomor++
	}

	for {
		fmt.Scan(&suara)

		if suara == 0 {
			break
		}
		suaraMasuk++
		index := sequentialSearch(kandidat, Max, suara)

		if index != -1 {
			totalSuara[index]++
			suaraSah++
		}
	}

	fmt.Println("Suara masuk : ", suaraMasuk)
	fmt.Println("Suara sah : ", suaraSah)

	for i = 0; i < Max; i++ {
		if totalSuara[i] > 0{
			fmt.Printf("%d: %d\n",kandidat[i], totalSuara[i])
		}
	}
}
```

### 2. [Soal Latihan Modul 12 (2)]
#### soal2.go

```go
package main

import "fmt"

const Max int = 20

type arrKandidat [Max]int

func sequentialSearch(T arrKandidat, n int, X int) int {
	var found int = -1
	var i int

	i = 0
	for i < n && found == -1 {
		if T[i] == X {
			found = i
		}
		i++
	}
	return found
}

func main() {
	var kandidat arrKandidat
	var totalSuara [Max]int
	var suara, suaraMasuk, suaraSah int
	var i int

	nomor := 1

	for i = 0; i < Max; i++ {
		kandidat[i] = nomor
		nomor++
	}

	for {
		fmt.Scan(&suara)

		if suara == 0 {
			break
		}
		suaraMasuk++
		index := sequentialSearch(kandidat, Max, suara)

		if index != -1 {
			totalSuara[index]++
			suaraSah++
		}
	}

	ketua := -1
	idxKetua := -1

	for i = 0; i < Max; i++ {
		if totalSuara[i] > ketua {
			ketua = totalSuara[i]
			idxKetua = i
		}
	}

	wakil := -1
	idxWakil := -1

	for i := 0; i < Max; i++ {
		if i != idxKetua && totalSuara[i] > wakil {
			wakil = totalSuara[i]
			idxWakil = i
		}
	}

	fmt.Println("Suara masuk : ", suaraMasuk)
	fmt.Println("Suara sah : ", suaraSah)
	if suaraSah > 0 {
		fmt.Println("Ketua RT : ", kandidat[idxKetua])
		fmt.Println("Wakil Ketua : ", kandidat[idxWakil])
	}
}
```

### 3. [Soal Latihan Modul 12 (3)]
#### soal3.go

```go
package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	isiArray(n)

	idx := posisi(n, k)
	if idx != -1 {
		fmt.Println(idx)
	} else {
		fmt.Println("TIDAK ADA")
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	var found int = -1
	var med int
	var kiri int = 0
	var kanan int = n - 1

	for kiri <= kanan && found == -1 {
		med = (kiri + kanan) / 2
		if k < data[med] {
			kanan = med - 1
		} else if k > data[med] {
			kiri = med + 1
		} else {
			found = med
		}
	}
	
	return found
}
```
### Output Unguided :

##### Output

##### Soal 1

![Screenshoot Output Soal 1](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%2012/Output/Output_soal1.png)
Program pilkart digunakan untuk mendata pada pemilihan ketua RT baru, terdapat 20 calon ketua yang bertanding memperebutkan suara warga.

Masukan berupa satu baris, berisi bilangan bulat valid. Data valid adalah integer dengan nilai di antara 1—20 (inklusif). Sementara, data tidak valid yaitu data bilangan bulat negatif dan bilangan yang melebihi range data. Jika data yang ditulis oleh user adalah angka 0, maka program berakhir.

Dalam program, terdapat function tambahan bernama sequentialSearch, digunakan untuk pengecekan data index array yang bernilai sama dengan nilai X pada parameter, yang diambil dari nilai suara di program utama (main).

Keluaran berupa jumlah suara masuk, suara sah, dan data para calon yang mendapatkan suara.



##### Soal 2

![Screenshoot Output Soal 2](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%2012/Output/Output_soal2.png)
Pada nomor 2 ini, melanjutkan program pilkart pada nomor 1. Namun pada program ini, akan dicari siapa pemenang pemilihan Ketua RT berdasarkan banyaknya voting dari para warga, dan Wakil Ketua yang mendapat banyak voting kedua setelahnya.

Masukan berupa satu baris, berisi bilangan bulat valid. Data valid adalah integer dengan nilai di antara 1—20 (inklusif). Sementara, data tidak valid yaitu data bilangan bulat negatif dan bilangan yang melebihi range data. Jika data yang ditulis oleh user adalah angka 0, maka program berakhir.

Jika terdapat jumlah nomor calon yang sama, maka pemilihan Ketua RT akan dilakukan berdasarkan data terurut membesar. 

Contohnya : 3 dan 19 sama-sama mendapat 2 suara, tetapi karena secara urutan membesar, 3 lebih dahulu muncul, maka calon nomor 3 yang menjadi Ketua, sementara nomor 19 akan menjadi Wakil.

Dalam program, terdapat function tambahan bernama sequentialSearch, digunakan untuk pengecekan data index array yang bernilai sama dengan nilai X pada parameter, yang diambil dari nilai suara di program utama (main).

Keluaran berupa jumlah suara masuk, suara sah, dan nomor ketua dan wakil ketua RT yang baru.

##### Soal 3

![Screenshoot Output Soal 3](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%2012/Output/Output_soal3.png)
Program posisidata digunakan untuk mencari posisi suatu bilangan pada kumpulan data integer yang sudah terurut membesar (ascending). Program ini menerapkan algoritma Binary Search sehingga proses pencarian dapat dilakukan dengan lebih efisien dibandingkan pencarian secara sekuensial.

Masukan terdiri dari dua baris. Baris pertama berisi dua bilangan bulat positif, yaitu n dan k. Nilai n menyatakan banyaknya data yang akan disimpan pada array, sedangkan k merupakan bilangan yang akan dicari posisinya.

Baris kedua berisi n buah bilangan bulat positif yang telah terurut secara membesar.

Dalam program, terdapat prosedur tambahan bernama isiArray yang digunakan untuk mengisi array data sebanyak n elemen dari input user.

Selain itu terdapat function tambahan bernama posisi yang digunakan untuk mencari lokasi bilangan k pada array data menggunakan algoritma Binary Search. Function akan mengembalikan indeks data apabila bilangan yang dicari ditemukan, atau mengembalikan nilai -1 apabila bilangan tersebut tidak ditemukan.

Keluaran terdiri dari satu baris. Jika bilangan k ditemukan, maka program akan mencetak posisi atau indeks dari bilangan tersebut pada array, dengan perhitungan indeks dimulai dari 0. Apabila bilangan k tidak ditemukan, maka program akan mencetak teks "TIDAK ADA".