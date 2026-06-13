package main

import "fmt"

type banyakRumah [1000000]int

func selectionSort(arr *banyakRumah, m int) {
	var i, j int
	var idx_min int
	i = 0
	for i < m-1 {
		idx_min = i 

		for j = i; j < m; j++{
			if (*arr)[idx_min] > (*arr)[j] {
				idx_min = j
			}
		}
		temp := (*arr)[idx_min]
		(*arr)[idx_min] = (*arr)[i]
		(*arr)[i] = temp
		i++
	}
}

func main() {
	var isiM banyakRumah
	var n, m int
	
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&isiM[j])
		}

		selectionSort(&isiM, m)
		
		for k := 0; k < m; k++{
			fmt.Print(isiM[k], " ")
		}
		fmt.Println()
	}

}
