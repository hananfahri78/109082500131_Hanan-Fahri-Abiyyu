package main

import "fmt"

func main() {
	var  g, kg, tambahan int
	fmt.Print("Berat parsel (gram) :")
	fmt.Scan(&g)
	kg = g/1000
	gr := g % 1000
	fmt.Println("Detail berat :", kg, "kg +", gr, "gr")
	kg2 := kg * 10000
	if gr >= 500 && kg <= 10 {
		tambahan = (gr * 5)
		fmt.Println("Detail biaya :", "Rp.", kg2, "+ Rp.", tambahan)
		total := tambahan + kg2
		fmt.Println("Total biaya :", total)
	}else if gr < 500 && kg <= 10 {
		tambahan = (gr * 15)
		fmt.Println("Detail biaya :", "Rp.", kg2, "+ Rp.", tambahan)
		kg2 = tambahan + kg2
		fmt.Print("Total biaya :", kg2)
	}else if kg > 10{
		tambahan = (gr * 5)
		fmt.Println("Detail biaya : ", "Rp.", kg2, "+", "Rp.", tambahan)
		fmt.Println("Total biaya :", kg2)
	}
	
}