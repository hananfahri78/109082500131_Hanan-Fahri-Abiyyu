package main

import "fmt"

const Max int = 20

type arrKandidat [Max]int

func sequentialSearch(T arrKandidat, n int, X int) int {
	var found int = -1
	var i int

	i = 0
	for i < n && found == -1 {
		if T[i] == X {
			found = i
		}
		i++
	}
	return found
}

func main() {
	var kandidat arrKandidat
	var totalSuara [Max]int
	var suara, suaraMasuk, suaraSah int
	var i int

	nomor := 1

	for i = 0; i < Max; i++ {
		kandidat[i] = nomor
		nomor++
	}

	for {
		fmt.Scan(&suara)

		if suara == 0 {
			break
		}
		suaraMasuk++
		index := sequentialSearch(kandidat, Max, suara)

		if index != -1 {
			totalSuara[index]++
			suaraSah++
		}
	}

	ketua := -1
	idxKetua := -1

	for i = 0; i < Max; i++ {
		if totalSuara[i] > ketua {
			ketua = totalSuara[i]
			idxKetua = i
		}
	}

	wakil := -1
	idxWakil := -1

	for i := 0; i < Max; i++ {
		if i != idxKetua && totalSuara[i] > wakil {
			wakil = totalSuara[i]
			idxWakil = i
		}
	}

	fmt.Println("Suara masuk : ", suaraMasuk)
	fmt.Println("Suara sah : ", suaraSah)
	if suaraSah > 0 {
		fmt.Println("Ketua RT : ", kandidat[idxKetua])
		fmt.Println("Wakil Ketua : ", kandidat[idxWakil])
	}
}
