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

![Screenshoot Output Soal 1](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%203/Output/Output1.png)
Membuat program untuk membantu Jonas dalam pengimplementasian tugas mata kuliah Matematika Diskrit mengenai kombinasi dan permutasi, ke dalam suatu program. Input terdiri dari 4 variabel bilangan asli a, b, c, d dengan syarat a >= c dan b >= d. Output berupa dua baris yang menyatakan hasil permutasi dan kombinasi a terhadap c, dan hasil permutasi dan kombinasi b terhadap d. Terdapat tiga function tambahan, digunakan untuk penghitungan faktorial, permutasi, dan kombinasi. Pada faktorial digunakan inisiasi variabel i, serta variabel hasil, dengan nilai 1. Untuk menghitungnya, digunakan perulangan dengan syarat selama i <= n, maka i akan bertambah 1. Permutasi memanggil function faktorial dengan cara, return faktorial(n) / faktorial(r) sesuai rumus. Begitu juga kombinasi, akan memanggil function faktorial dengan return, dan disesuaikan rumus kombinasi. Pada function utama, terdapat percabangan, percabangan pertama dimana jika a>= c maka akan mengeluarkan output hasil dari permutasi dan kombinasi a, c. Percabangan kedua, jika b >= d maka mengeluarkan hasil permutasi dan kombinasi b, d. 

##### Soal 2

![Screenshoot Output Soal 2](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%203/Output/Output2.png)
Menghitung fungsi komposisi. Diberikn tiga buah fungsi matematika f, g, h. Ketiganya dibuatkan function tambahan dan akan digabungkan masing-masing sesuai pada soal. Yaitu output pertama (fogoh)(a), kedua (gohof)(b), ketiga (hofog)(c). function f(x) = x^2, function g(x) = x - 2, function h(x) = x + 1. Selanjutnya masuk ke function utama, terdapat tiga buah variabel bilangan bulat(integer) a, b, c, input dari masing-masing variabel tersebut akan dimasukkan kedalam fungsi komposisi di atas. Output akan menghasilkan fungsi komposisi dari (fogoh)(a), (gohof)(b), dan (hofog)(c).  

##### Soal 3

![Screenshoot Output Soal 3](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%203/Output/Output3.png)
Menentukan posisi titik apakah di dalam atau di luar lingkaran. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dan radius r. Diberikan dua buah lingkaran dan menentukan posisi titik sembarang menggunakan variabel (x, y) berdasarkan dua lingkaran tersebut. Terdapat dua function tambahan, jarak dan didalam. Jarak untuk menambahkan rumus mencari jarak, dengan menggunakan math.Sqrt untuk menghitung akar kuadrat. Di dalamnya terdapat variabel (a, b, c, d float64) -> float64, serta mengembalikan jarak antara titk (a,b) dan (c,d). Function didalam  terdapat variabel (cx,cy,r,x,y float64) -> bool, serta mengembalikan nilai true apabila titik (x,y) berada did alam lingkaran yang memiliki titik pusat (cx,cy) dan radius r. Pada function utama, kita menginputkan nilai cx baris 1 2, cy baris 1 2, r baris 1 2, x, dan y. Selanjutnya, ada variabel tambahan yaitu titik 1 dan titik 2 digunakan untuk menggantikan tiap variabel input dengan tipe data awal integer, menjadi float64, serta digunakan sebagai variabel yang dimasukkan dalam percabangan, untuk menentukan posisi titik. Output akan mengeluarkan berupa string "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", atau "Titik di luar lingkaran 1 dan 2" cara pemisahan tersebut menggunakan percabangan if else. 