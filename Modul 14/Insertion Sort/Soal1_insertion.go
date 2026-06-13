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
