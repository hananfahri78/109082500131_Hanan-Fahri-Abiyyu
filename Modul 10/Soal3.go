package main
import "fmt"
type arrBalita [100]float64
var dataBalita int

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < dataBalita; i++ {
		if arrBerat[i] < *bMin{
			*bMin = arrBerat[i]
		}

		if arrBerat[i] > *bMax{
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita) float64 {
	var rerata float64
	var total float64 = 0

	for i := 0; i < dataBalita; i++ {
		total += arrBerat[i]
	}
	rerata = total / float64(dataBalita)
	return rerata
}

func main() {

	var beratBalita arrBalita
	var rataRata float64
	var bMin, bMax float64
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&dataBalita)

	for i := 0; i < dataBalita; i++ {
		fmt.Print("Masukan berat balita ke-", i+1, " : ")
		fmt.Scan(&beratBalita[i])
	}

	hitungMinMax(beratBalita, &bMin, &bMax)
	rataRata = rerata(beratBalita)
	fmt.Printf("Berat balita minimum  : %.2f\n", bMin)
	fmt.Printf("Berat balita maksimum : %.2f\n", bMax)
	fmt.Printf("Rerata berat balita   : %.2f\n", rataRata)
}
