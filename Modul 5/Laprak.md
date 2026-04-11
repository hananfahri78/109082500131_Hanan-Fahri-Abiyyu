# <h1 align="center">Laporan Praktikum Modul 5 </h1>
<p align="center">[Hanan Fahri Abiyyu] - [109082500131]</p>

## Unguided

### 1. [Soal Modul 5]
#### soal1-Fibonacci.go

```go
package main

import "fmt"

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}
func main() {
	var angka int
	fmt.Scan(&angka)
	for i := 0; i <= angka; i++ {
		fmt.Print(fibo(i), " ")
	}
}

```

### 2. [Soal  Modul 5]
#### soal2-polabintang.go

```go
package main

import "fmt"

func bintang(n int) string {
	if n == 0 {
		return ""
	}
	bintang(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()

	return ""
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(bintang(n))
}

```

### 3. [Soal Modul 5]
#### soal3-Faktorbil.go

```go
package main

import "fmt"

func faktor(x int, i int) {
	if i > x {
		return
	} else {
		if x%i == 0 {
			fmt.Print(i, " ")
		}
		faktor(x, i+1)
	}
}
func main() {
	var bil int
	fmt.Scan(&bil)
	faktor(bil, 1)
}

```

### 4. [Soal Modul 5]
#### soal4-urutanbil.go

```go
package main

import "fmt"

func urutan(x int) {
	if x == 1 {
		fmt.Print(1, " ")
		return
	} else {
		fmt.Print(x, " ")
		urutan(x - 1)
		fmt.Print(x, " ")
	}
}
func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	urutan(bilangan)
}

```

### 5. [Soal Modul 5]
#### soal5-barisanganjil.go

```go
package main

import "fmt"

func ganjil(x int, i int) {
	if i > x {
		return
	} else {

		fmt.Print(i, " ")
		ganjil(x, i+2)

	}
}
func main() {
	var angka int
	fmt.Scan(&angka)
	ganjil(angka, 1)
}

```

### 6. [Soal Modul 5]
#### soal6-Caripangkat.go

```go
package main

import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)

}
func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Print(pangkat(x, y))
}

```

### Output Unguided :

##### Output

##### Soal 1

![Screenshoot Output Soal 1](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%205/Output/Output_soal1.png)
Membuat sebuah program menghitung deret Fibonacci, yang definisinya adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilaisuku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Pada program ini digunakannya fungsi rekursif untuk menjadi tempat pemanggilan  dari suatu nilai input yang dilakukan pada function main (utama), serta dicetak hasil dengan memanggil function rekursif yang telah dibuat, ke dalam main fungsi. Contoh suku n = 2, (Sn - 1) + (Sn - 2). Nilai n suku ke-2 hasil dari penjumlahan dua suku sebelumnya yang berarti 0 dan 1, hasil n = 2 adalah 1.

##### Soal 2

![Screenshoot Output Soal 2](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%205/Output/Output_soal2.png)
Pada program nomor 2, diperintahkan membuat program mencetak pola bintang segitiga siku-siku sesuai dari input user. Menggunakan fungsi rekursif untuk tempat penghitungan dan pembentukan pola bintang, yang kemudian akan dipanggil ke func main untuk pengeksekusian yang akan menghasilkan output. Pada program ini digunakannya sebuah perulangan dan percabangan, perulangan berfungsi sebagai pencetakan bintang sebanyak N kali, percabangan digunakan untuk mengecek apakah nilai sama dengan 0 atau bukan, jika bukan lanjut ke perulangan kedua dan seterusnya sampai 0. Pada fungsi bintang, perhitungan pola digunakan dengan cara menghitung dari nilai awal/input, masuk ke variabel n sebagai isian nilai di function bintang. Selanjutnya akan dihitung dari besar ke kecil (turun), program terus mengulang sampai nilai n = 0, kemudian dibalik urutannya menjadi kecil ke besar(naik) dengan menggunakan bintang(x-1) sebelum for. Namun, saat n bernilai 0 tetap akan mencetak satu baris kosong sebelum berhentinya program.

Pada bagian func main, input variabel dengan tipe data integer, serta mencetak hasil segitiga dengan memanggil fungsi bintang di dalam fmt.Println/percetakan.

##### Soal 3

![Screenshoot Output Soal 3](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%205/Output/Output_soal3.png)
Menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Untuk program ini bentuk pengimplementasian dari rekursif berupa prosedur rekursif. 
##### Soal 4

![Screenshoot Output Soal 4](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%205/Output/Output_soal4.png)

##### Soal 5

![Screenshoot Output Soal 5](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%205/Output/Output_soal5.png)

##### Soal 6

![Screenshoot Output Soal 6](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%205/Output/Output_soal6.png)