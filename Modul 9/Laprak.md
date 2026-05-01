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

![Screenshoot Output Soal 1](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%209/Output/Output-Soal1.png)
Menentukan posisi titik apakah berada di dalam atau di luar lingkaran. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dan radius r. Diberikan dua buah lingkaran dan sebuah titik (x, y) untuk menentukan posisinya terhadap kedua lingkaran. Digunakan dua tipe bentukan yaitu titik untuk menyimpan koordinat x dan y, serta lingkaran untuk menyimpan titik pusat dan radius. Terdapat function jarak untuk menghitung jarak dua titik menggunakan math.Sqrt, dan function didalam untuk mengecek apakah titik berada di dalam lingkaran dengan membandingkan jarak ke pusat dengan radius.

Pada function utama, diinputkan data dua lingkaran dan satu titik, kemudian kedua lingkaran disimpan dalam array lingkaranArr. Dilakukan pengecekan posisi titik terhadap masing-masing lingkaran menggunakan function didalam. Output berupa string "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2" yang ditentukan menggunakan percabangan if else. 

##### Soal 2

![Screenshoot Output Soal 2](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%209/Output/Output-Soal2.png)
Membuat program menggunakan array untuk menampilkan berbagai informasi sesuai dengan poin-poin yang tertera pada soal. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Program dibuat untuk mengisi array sebanyak N elemen sesuai input pengguna. Data disimpan dalam array bertipe integer, kemudian digunakan untuk menampilkan beberapa informasi berdasarkan indeks dan nilai elemen. Proses yang dilakukan meliputi menampilkan seluruh isi array, menampilkan elemen dengan indeks ganjil dan genap, serta menampilkan elemen pada indeks kelipatan tertentu yang diperoleh dari input pengguna.

Selain itu, program juga dapat menghapus elemen pada indeks tertentu dengan cara menggeser elemen setelahnya, kemudian menampilkan kembali isi array yang sudah diperbarui. Selanjutnya dihitung nilai rata-rata dan simpangan baku dari seluruh elemen dalam array menggunakan perhitungan matematika. Terakhir, program menghitung frekuensi kemunculan suatu nilai tertentu di dalam array. Semua proses tersebut dilakukan menggunakan perulangan dan percabangan sesuai kebutuhan.

##### Soal 3

![Screenshoot Output Soal 3](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%209/Output/Output-Soal3.png)
Membuat program menggunakan array untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan sepak bola. Program meminta input nama dua klub yang bertanding, kemudian meminta skor hasil pertandingan secara berulang. Data disimpan dalam array bertipe string yang berisi hasil pemenang dari setiap pertandingan, yaitu nama klub yang menang atau “Draw” jika skor sama.

Proses input skor akan terus dilakukan hingga salah satu skor bernilai negatif. Pada setiap input, program langsung menentukan hasil pertandingan dan menampilkannya. Di akhir program, proses dihentikan dengan menampilkan pesan bahwa pertandingan telah selesai. Semua proses dilakukan menggunakan perulangan dan percabangan sesuai kondisi yang diberikan.

##### Soal 4

![Screenshoot Output Soal 4](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%209/Output/Output-Soal4.png)
Membuat program menggunakan array untuk menampung sekumpulan karakter, kemudian dilakukan proses untuk membalikkan urutan isi array dan memeriksa apakah termasuk palindrom. Data karakter disimpan dalam array bertipe rune melalui proses input satu per satu hingga mencapai tanda titik atau batas maksimum. Program menggunakan beberapa subprogram yaitu isiArray untuk mengisi array, cetakArray untuk menampilkan isi array, balikanArray untuk membalik urutan karakter dengan hanya menukar huruf dan mengabaikan karakter non-huruf, serta palindrom untuk mengecek apakah susunan karakter sama jika dibaca dari depan maupun belakang.

Pada function utama, array diisi menggunakan isiArray, kemudian dilakukan pengecekan palindrom dengan membandingkan array asli dengan hasil array yang telah dibalik. Hasil pengecekan ditampilkan dalam bentuk boolean. Setelah itu, array dibalik menggunakan balikanArray dan ditampilkan kembali menggunakan cetakArray.