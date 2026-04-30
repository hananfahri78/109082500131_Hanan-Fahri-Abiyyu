package main
import "fmt"
const NMAX int = 127
type tabel [NMAX]rune
	
func isiArray(t *tabel, n *int) {
	var char rune
	*n = 0
	
	for {
		fmt.Scanf("%c", &char)

		if char == '.' || *n >= NMAX{
			break
		}

		(*t)[*n] = char

		*n = *n + 1
	}
}
func cetakArray(t tabel, n int) {
	
	for i := 0; i < n; i++{
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}
func balikanArray(t *tabel, n int) {
	kiri := 0
	kanan := n-1
	
	for kiri < kanan {
		kiriChar := (*t)[kiri]
		kananChar := (*t)[kanan]
		if !(('A' <= kiriChar && kiriChar <= 'Z') || ('a' <= kiriChar && kiriChar <= 'z')) {
			kiri++
		}else if !(('A' <= kananChar && kananChar <= 'Z') || ('a' <= kananChar && kananChar <= 'z')) {
			kanan--
		}else{
			(*t)[kiri], (*t)[kanan] = (*t)[kanan], (*t)[kiri]
			kiri++
			kanan--
		}
		
	}
}

func palindrom(t tabel, n int) bool{
	var temp tabel

	for i := 0; i < n; i++{
		temp[i] = t[i]
	}

	balikanArray(&temp, n)
	for i := 0; i < n; i++{
		if t[i] != temp[i] {
			return false
		}
	}
	return true
}
func main() {
	var tab tabel
	var m int
	
	isiArray(&tab, &m)
	fmt.Println("Palindrom ?", palindrom(tab, m))

	balikanArray(&tab, m)
	cetakArray(tab, m)


	
}