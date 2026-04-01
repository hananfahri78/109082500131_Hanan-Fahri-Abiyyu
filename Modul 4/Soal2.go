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