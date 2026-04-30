# <h1 align="center">Laporan Praktikum Modul 9 </h1>
<p align="center">[Hanan Fahri Abiyyu] - [109082500131]</p>

## Unguided

### 1. [Soal Modul 9]
#### soal1-Lingkaran.go

```go
package main
import (
	"fmt"
	"math"
)
type titik struct {
	x int
	y int
}
type lingkaran struct {
	cx int
	cy int
	r  int
}
func jarak(p, q titik) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(float64(dx * dx) + float64(dy * dy))
	
}
func didalam(c lingkaran, p titik) bool {
	tipus := titik{c.cx, c.cy}
	return jarak(tipus, p) <= float64(c.r)
}
func main() {
	var l1, l2 lingkaran
	var T titik

	fmt.Scan(&l1.cx, &l1.cy, &l1.r)
	fmt.Scan(&l2.cx, &l2.cy, &l2.r)
	fmt.Scan(&T.x, &T.y)

	lingkaranArr := [2]lingkaran {l1, l2}

	dalam1 := didalam(lingkaranArr[0], T)
	dalam2 := didalam(lingkaranArr[1], T)

	if dalam1 && dalam2{
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	}else if dalam1{
		fmt.Println("Titik di dalam lingkaran 1")
	}else if dalam2{
		fmt.Println("Titik di dalam lingkaran 2")
	}else{
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}


```

### 2. [Soal  Modul 9]
#### soal2-Array.go

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var arr [100]int
	var N int
	fmt.Print("Tentukan jumlah array : ")
	fmt.Scan(&N)
	var i int
	// output soal a
	fmt.Print("Indeks : ")
	for i = 0; i < N; i++ {
		fmt.Scan(&arr[i])
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// output soal b
	for i = 0; i < N; i++ {
		if i%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// output soal c
	for i = 0; i < N; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// output soal d
	var x int
	fmt.Print("Tentukan Kelipatan : ")
	fmt.Scan(&x)
	for i = 0; i < N; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// output soal e
	var remove int
	fmt.Print("Tentukan urutan indeks yang akan dihapus : ")
	fmt.Scan(&remove)
	for i := remove; i < N-1; i++ {
		arr[i] = arr[i+1]
	}
	N -= 1
	for i = 0; i < N; i++{
		fmt.Print(arr[i]," ")
	}
	fmt.Println()

	// output soal f
	var total, rata2 float64
	for i = 0; i < N; i++ {
		total = total + float64(arr[i])
	}
	rata2 = total / float64(N)
	fmt.Printf("%.2f\n", rata2)

	// output soal g
	var simpangan_baku, totalS float64
	for i = 0; i < N; i++ {
		totalS += (float64(arr[i]) - rata2) * (float64(arr[i]) - rata2)
	}
	hasil := totalS / float64(N)
	simpangan_baku = math.Sqrt(hasil)
	
	fmt.Printf("%.2f\n", simpangan_baku)

	// output soal h
	fmt.Print("Tentukan nilai frekuensi indeks : ")
	fmt.Scan(&x)
	counter := 0
	for i = 0; i < N; i++ {
		if arr[i] == x {
			counter++
		}
	}
	fmt.Print(counter)

}

```

### 3. [Soal Modul 9]
#### soal3-SkorBola.go

```go
package main

import "fmt"

func main() {
	var skorA, skorB int
	var klubA, klubB string
	var arr [100]string

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	i := 1
	idx := 0
	for {
		fmt.Print("Pertandingan ", i, ": ")
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			arr[idx] = klubA
		} else if skorB > skorA {
			arr[idx] = klubB
		} else {
			arr[idx] = "Draw"
		}

		fmt.Println("Hasil ", i, ": ", arr[idx])

		idx++
		i++
	}
	fmt.Println("Pertandingan selesai")
}

```

### 4. [Soal Modul 9]
#### soal4-Urutkata.go

```go
package main
import "fmt"
const NMAX int = 127
type tabel [NMAX]rune
	
func isiArray(t *tabel, n *int) {
	var char rune
	*n = 0
	
	for {
		fmt.Scanf("%c", &char)

		if char == '.' || *n >= NMAX{
			break
		}

		(*t)[*n] = char

		*n = *n + 1
	}
}
func cetakArray(t tabel, n int) {
	
	for i := 0; i < n; i++{
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}
func balikanArray(t *tabel, n int) {
	kiri := 0
	kanan := n-1
	
	for kiri < kanan {
		kiriChar := (*t)[kiri]
		kananChar := (*t)[kanan]
		if !(('A' <= kiriChar && kiriChar <= 'Z') || ('a' <= kiriChar && kiriChar <= 'z')) {
			kiri++
		}else if !(('A' <= kananChar && kananChar <= 'Z') || ('a' <= kananChar && kananChar <= 'z')) {
			kanan--
		}else{
			(*t)[kiri], (*t)[kanan] = (*t)[kanan], (*t)[kiri]
			kiri++
			kanan--
		}
		
	}
}

func palindrom(t tabel, n int) bool{
	var temp tabel

	for i := 0; i < n; i++{
		temp[i] = t[i]
	}

	balikanArray(&temp, n)
	for i := 0; i < n; i++{
		if t[i] != temp[i] {
			return false
		}
	}
	return true
}
func main() {
	var tab tabel
	var m int
	
	isiArray(&tab, &m)
	fmt.Println("Palindrom ?", palindrom(tab, m))

	balikanArray(&tab, m)
	cetakArray(tab, m)
	
}

```


### Output Unguided :

##### Output

##### Soal 1

![Screenshoot Output Soal 1](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%205/Output/Output_soal1.png)
Membuat sebuah program menghitung deret Fibonacci, yang definisinya adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilaisuku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Pada program ini digunakannya fungsi rekursif untuk menjadi tempat pemanggilan  dari suatu nilai input yang dilakukan pada function main (utama), serta dicetak hasil dengan memanggil function rekursif yang telah dibuat, ke dalam main fungsi. Contoh suku n = 2, (Sn - 1) + (Sn - 2). Nilai n suku ke-2 hasil dari penjumlahan dua suku sebelumnya yang berarti 0 dan 1, hasil n = 2 adalah 1.

##### Soal 2

![Screenshoot Output Soal 2](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%205/Output/Output_soal2.png)
Pada program nomor 2, diperintahkan membuat program mencetak pola bintang segitiga siku-siku sesuai dari input user. Menggunakan fungsi rekursif untuk tempat penghitungan dan pembentukan pola bintang, yang kemudian akan dipanggil ke func main untuk pengeksekusian yang akan menghasilkan output. Pada program ini digunakannya sebuah perulangan dan percabangan, perulangan berfungsi sebagai pencetakan bintang sebanyak N kali, percabangan digunakan untuk mengecek apakah nilai sama dengan 0 atau bukan, jika kondisi tersebut belum terpenuhi, maka fungsi akan terus memanggil dirinya sendiri dengan nilai n-1 hingga mencapai base case. Pada fungsi bintang, perhitungan pola digunakan dengan cara menghitung dari nilai awal/input, masuk ke variabel n sebagai isian nilai di function bintang. Selanjutnya akan dihitung dari besar ke kecil (turun), program terus mengulang sampai nilai n = 0, kemudian dibalik urutannya menjadi kecil ke besar(naik) dengan menggunakan bintang(x-1) sebelum for.

Pada bagian func main, program menerima input berupa bilangan integer dari pengguna, kemudian memanggil fungsi bintang untuk mencetak pola sesuai dengan nilai yang diberikan.

##### Soal 3

![Screenshoot Output Soal 3](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%205/Output/Output_soal3.png)
Menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Untuk program ini bentuk pengimplementasian dari rekursif berupa prosedur rekursif. Seperti func faktor, diperlukannya dua variabel yaitu x dan i berfungsi untuk mencari angka dari 1 sampai ke N yang merupakan faktor dari bilangan N. Variabel x digunakan untuk menyimpan nilai dari input yang dilakukan user. percabangan i > x digunakan sebagai kondisi berhenti, yang dimana jika nilai i lebih besar dari x, maka program berhenti mencari. 

Bagian main fungsi (utama) ditambahkan variabel bil dengan tipe data integer, menginput, memasukkan ke dalam func faktor, kemudian akan dieksekusi di dalam func faktor.

##### Soal 4

![Screenshoot Output Soal 4](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%205/Output/Output_soal4.png)
Membuat program dengan menggunakan rekursif untuk menampilkan barisan bilangan terurut mengecil ke 1 dan membesar kembali ke angka input. Base case terjadi saat n = 1, di mana program mencetak angka 1 dan berhenti. Jika belum mencapai base case, program akan mencetak nilai n, kemudian memanggil fungsi dengan n-1 sehingga terbentuk urutan menurun. Setelah itu, saat proses kembali, nilai n dicetak lagi sehingga terbentuk urutan membesar.

Pada function main (utama), program menerima input integer dan memanggil prosedur secara rekursif untuk menampilkan hasil.


##### Soal 5

![Screenshoot Output Soal 5](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%205/Output/Output_soal5.png)
Membuat program mencetak bilangan ganjil dari input bilangan bulat positif. Membuat prosedur rekursif yang digunakan sebagai penyimpanan nilai variabel dari input pengguna. 

Pada function ganjil terdapat dua variabel bertipe data integer, x dan i. Dengan x digunakan untuk tempat dari nilai yang dimasukkan dan i berfungsi untuk perulangan angka ganjil. Nilai awal i adalah 1, kemudian akan bertambah 2 setiap pemanggilan rekursif sehingga menghasilkan deret bilangan ganjil.

Main fungsi (utama) digunakan untuk membuat pengguna dapat memasukan angka bilangan bulat dan memanggil subprogram rekursif untuk menampilkan hasil.
##### Soal 6

![Screenshoot Output Soal 6](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%205/Output/Output_soal6.png)
Membuat sebuah program untuk menghitung hasil perpangkatan dari dua bilangan, yaitu x dan y, menggunakan rekursif. Variabel x sebagai bilangan dasar dan y sebagai pangkat.

Pada bagian func pangkat, nilai y digunakan sebagai pengontrol rekursi. Artinya jika y = 0, fungsi akan mengembalikan nilai 1 sebagai kondisi berhenti. Jika belum, maka fungsi akan mengalikan x dengan hasil dari pemanggilan fungsi pangkat(x, y-1). Proses akan terus berulang hingga nilai y mencapai 0, sehingga hasil perpangkatan dapat diperoleh secara bertahap.

Selanjutnya bagian main fungsi, program menerima dua input bilangan bulat dari pengguna, kemudian memanggil fungsi rekursif untuk menghitung dan menampilkan hasil perpangkatan.