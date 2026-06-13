package main

import "fmt"

type banyakRumah [1000000]int

func selectionSortAsc(arr *banyakRumah, m int) {
	var i, j int
	var idx_min int
	
	i = 0
	for i < m-1 {
		idx_min = i 

		for j = i+1; j < m; j++{
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

func selectionSortDesc(arr *banyakRumah, m int) {
	var i, j int
	var idx_max int

	i = 0
	for i < m-1{
		idx_max = i

		for j = i+1; j < m; j++{
			if (*arr)[idx_max] < (*arr)[j] {
				idx_max = j
			}
		}
		temp := (*arr)[idx_max]
		(*arr)[idx_max] = (*arr)[i]
		(*arr)[i] = temp
		i++
	}
}

func main() {
	var isiM banyakRumah
	var n, m int
	var ganjil, genap banyakRumah
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&isiM[j])
		}

		nGanjil := 0
		nGenap := 0
		for j := 0; j < m; j++{
			if isiM[j]%2 != 0{
				ganjil[nGanjil] = isiM[j]
				nGanjil++
			}else{
				genap[nGenap] = isiM[j]
				nGenap++
			}
		}
		selectionSortAsc(&ganjil, nGanjil)
		selectionSortDesc(&genap, nGenap)
		
		for j := 0; j < nGanjil; j++{
			fmt.Print(ganjil[j], " ")
		}

		for j := 0; j < nGenap; j++{
			fmt.Print(genap[j], " ")
		}
		fmt.Println()
	}

}
