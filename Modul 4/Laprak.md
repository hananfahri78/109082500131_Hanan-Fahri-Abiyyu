# <h1 align="center">Laporan Praktikum Modul 4 </h1>
<p align="center">[Hanan Fahri Abiyyu] - [109082500131]</p>

## Unguided

### 1. [Soal Latihan Modul 4]
#### soal1.go

```go
package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	var i int
	for i = 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int, hasil *int) {
	var fn, fnr int
	faktorial(n, &fn)
	faktorial(n-r, &fnr)
	*hasil = fn / fnr
}

func kombinasi(n, r int, hasil *int) {
	var fn, fnr, fr int
	faktorial(n, &fn)
	faktorial(r, &fr)
	faktorial(n-r, &fnr)
	*hasil = fn / (fr * (fnr))
}

func main() {
	var a, b, c, d, hasil int
	fmt.Scan(&a, &b, &c, &d)
	if a >= c {
		permutasi(a, c, &hasil)
		fmt.Print(hasil, " ")
		kombinasi(a, c, &hasil)
		fmt.Println(hasil)
	}

	if b >= d {
		permutasi(b, d, &hasil)
		fmt.Print(hasil, " ")
		kombinasi(b, d, &hasil)
		fmt.Println(hasil)
	}

}

```

### 2. [Soal Latihan Modul 2B]
#### soal2.go

```go
package main
import "fmt"

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0
	var waktu int
	for i := 0; i < 8; i++{
		fmt.Scan(&waktu)
		if waktu != 301 {
			*soal++
			*skor += waktu 
		}
	}
}

func main() {
	var(
		 peserta, pemenang string
		 skor, soal int
	     maxSoal, minSkor int
	)
	maxSoal = -1
	minSkor = 999999
	for {
		fmt.Scan(&peserta)
			if peserta == "Selesai"{
				break
			}
	
	hitungSkor(&soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			pemenang = peserta
		}
	}
	fmt.Println(pemenang, maxSoal, minSkor)
}
```

### 3. [Soal Latihan Modul 2C]
#### soal3.go

```go
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}

func cetakDeret(n int) {
	var done bool
	for !done{
		fmt.Print(n, " ")
		if n == 1{
			done = true
		}
		if n%2 == 0{
			n = n / 2 
		}else if n%2 != 0{
			n = 3 * n + 1
		} 
		
	}
}
```
### Output Unguided :

##### Output

##### Soal 1

![Screenshoot Output Soal 1](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%204/Output/Output1.png)
Membuat program untuk membantu Jonas dalam pengimplementasian tugas mata kuliah Matematika Diskrit mengenai kombinasi dan permutasi, ke dalam suatu program. 

Input terdiri dari 4 variabel bilangan asli a, b, c, d dengan syarat a >= c dan b >= d. Output berupa dua baris yang menyatakan hasil permutasi dan kombinasi a terhadap c, dan hasil permutasi dan kombinasi b terhadap d.

Terdapat tiga prosedur tambahan yang dideklarasikan sebagai func, digunakan untuk penghitungan faktorial, permutasi, dan kombinasi. Pada faktorial digunakan inisiasi variabel i, dan variabel hasil yang dikirim melalui pointer, dengan nilai awal 1. Untuk menghitungnya, digunakan perulangan dengan syarat selama i <= n, maka i akan bertambah 1. 

Permutasi dan kombinasi memanfaatkan parameter hasil berupa pointer. Permutasi menghitung nilai menggunakan rumus n! / (n-r)! dengan memanggil prosedur faktorial untuk memperoleh nilai faktorial yang dibutuhkan.

Kombinasi menghitung nilai menggunakan rumus n! / r! (n-r)! dengan cara yang sama, yaitu memanggil prosedur faktorial untuk setiap komponen perhitungan.

Pada fungsi utama (main), terdapat percabangan. Jika a>= c, maka program akan memanggil prosedur permutasi dan kombinasi untuk menghitung nilai a terhadap c, kemudian menampilkan output hasil. Percabangan kedua, jika b >= d, maka program akan melakukan hal yang sama untuk b terhadap d, kemudian menampilkan output hasil.

##### Soal 2

![Screenshoot Output Soal 2](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%204/Output/Output2.png)
Membuat program menentukan pemenang dalam suatu kompetisi pemrograman. Setiap peserta diberikan 8 soal yang harus diselesaikan dalam waktu 5 jam. Penentuan pemenang dilakukan berdasarkan jumlah soal yang berhasil diselesaikan terbanyak. Jika terdapat peserta dengan jumlah soal selesai yang sama, pemenang akan ditentukan dari total waktu pengerjaan yang paling kecil.

Input terdiri dari beberapa data peserta. Setiap baris input diawali dengan nama peserta, diikuti oleh 8 bilangan bulat (integer) yang menyatakan waktu (dalam menit) yang dibutuhkan untuk menyelesaikan masing-masing soal. Jika suatu soal tidak berhasil diselesaikan, maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit).

Program dibuat secara modular dengan menggunakan prosedur hitungSkor untuk menghitung jumlah soal dan total waktu yang memiliki parameter formal berupa jumlah soal yang diselesaikan dan total waktu pengerjaan. Prosedur tersebut bertugas membaca waktu pengerjaan untuk 8 soal, kemudian menghitung berapa banyak soal yang berhasil diselesaikan (waktu yang kurang dari 301 menit) serta menjumlahkan total waktu dari soal-soal tersebut.

Pada fungsi utama (main), program membaca nama peserta, kemudian memanggil prosedur hitungSkor untuk memperoleh jumlah soal dan total waktu dari masing-masing peserta. Selanjutnya dilakukan perbandingan untuk menentukan pemenang, yaitu peserta dengan jumlah soal terbanyak atau, jika sama, dengan total waktu paling kecil.

Output program berupa satu baris yang menampilkan nama pemenang, jumlah soal diselesaikan, serta total waktu yang dibutuhkan.

##### Soal 3

![Screenshoot Output Soal 3](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%204/Output/Output3.png)
Menghitung suatu deret bilangan dan mengimplementasikannya ke dalam suatu program, menggunakan bahasa pemrograman GO.

Input berupa satu variabel bilangan asli yang dideklarasikan dengan n. Terdapat aturan perhitungan suku selanjutnya, jika n saat itu bilangan genap, maka suku berikutnya adalah n/2. Namun, jika ganjil maka suku berikutnya bernilai 3*n + 1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya, hingga berakhir ketika suku terakhir bernilai 1. 

Menggunakan prosedur cetakDeret dideklarasikan sebagai func. Di dalamnya terdapat perulangan yang digunakan untuk mencetak setiap suku deret. Perulangan akan terus berjalan hingga nilai n sama dengan 1 sebagai kondisi berhenti.

Pada setiap iterasi dilakukan percabangan sesuai syarat ganjil atau genap, untuk menghasilkan suku berikutnya, dan setiap nilai n yang dihasilkan akan dicetak sebagai bagian dari deret.

Selanjutnya pada bagian fungsi utama (main), terdapat variabel n bertipe data integer. Variabel tersebut digunakan sebagai input bilangan, kemudian program akan memanggil prosedur cetakDeret untuk mencetak hasil dari input variabel n.