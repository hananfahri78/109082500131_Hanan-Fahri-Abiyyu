# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Hanan Fahri Abiyyu] - [109082500131]</p>

## Unguided

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

#### soal2.go

```go
package main

import "fmt"

func main() {
	var (
		warna1, warna2, warna3, warna4 string
		temp bool
	)
	fmt.Println("Urutkan warna sesuai rules")
	temp = true
	i := 1
	for i <= 5{
		fmt.Print("Percobaan ", i)
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)
		if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu"{
			temp = false
		}
		i++ 
	}
	fmt.Println("Berhasil :", temp)
}
```

#### soal3.go

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```
### Output Unguided :

##### Output
![Screenshoot Output Unguided 1_1](https://github.com/hananfahri78/109082500131_Hanan-Fahri-Abiyyu/blob/main/Modul%201/Output/output_soal2.png)
[penjelasan]