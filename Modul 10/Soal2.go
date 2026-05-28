package main
import "fmt"

type arrayMax [1000]float64

func main() {
	var x, y int
	var ikan arrayMax
	var totalPerWadah arrayMax
	var total, rata2 float64	

	fmt.Scan(&x, &y)
	for i := 0; i < x; i++{
		fmt.Scan(&ikan[i])
	}
	idxWadah := 0
	for i := 0; i < x; i++{
		totalPerWadah[idxWadah] += ikan[i]

		if (i+1)%y == 0{
			idxWadah++
		}
	} 

	jmlWadah := x/y
	if x%y != 0{
		jmlWadah++
	}

	for i := 0; i < jmlWadah; i++{
		fmt.Print(totalPerWadah[i], " ")
		total += totalPerWadah[i]
	}

	fmt.Println()

	rata2 = total / float64(jmlWadah)
	fmt.Println("Berat rata-rata ikan di tiap wadah : ",rata2)
}