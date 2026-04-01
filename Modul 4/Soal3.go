package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}

func cetakDeret(n int) {
	var done bool
	for !done{
		fmt.Print(n, " ")
		if n == 1{
			done = true
		}
		if n%2 == 0{
			n = n / 2 
		}else if n%2 != 0{
			n = 3 * n + 1
		} 
		
	}
}