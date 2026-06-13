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
