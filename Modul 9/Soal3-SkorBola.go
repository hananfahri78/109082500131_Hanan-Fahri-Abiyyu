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
