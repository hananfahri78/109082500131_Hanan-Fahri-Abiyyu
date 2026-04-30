package main
import (
	"fmt"
	"math"
)
type titik struct {
	x int
	y int
}
type lingkaran struct {
	cx int
	cy int
	r  int
}
func jarak(p, q titik) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(float64(dx * dx) + float64(dy * dy))
	
}
func didalam(c lingkaran, p titik) bool {
	tipus := titik{c.cx, c.cy}
	return jarak(tipus, p) <= float64(c.r)
}
func main() {
	var l1, l2 lingkaran
	var T titik

	fmt.Scan(&l1.cx, &l1.cy, &l1.r)
	fmt.Scan(&l2.cx, &l2.cy, &l2.r)
	fmt.Scan(&T.x, &T.y)

	lingkaranArr := [2]lingkaran {l1, l2}

	dalam1 := didalam(lingkaranArr[0], T)
	dalam2 := didalam(lingkaranArr[1], T)

	if dalam1 && dalam2{
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	}else if dalam1{
		fmt.Println("Titik di dalam lingkaran 1")
	}else if dalam2{
		fmt.Println("Titik di dalam lingkaran 2")
	}else{
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
