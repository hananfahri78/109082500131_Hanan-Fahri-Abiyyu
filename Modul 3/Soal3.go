package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}
func main() {
	var (
		cx1, cy1, r1 int
		cx2, cy2, r2 int
		x, y         int
	)
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	titik1 := didalam(float64(cx1), float64(cy1), float64(r1), float64(x), float64(y))
	titik2 := didalam(float64(cx2), float64(cy2), float64(r2), float64(x), float64(y))


	if titik1 && titik2{
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	}else if titik1{
		fmt.Println("Titik di dalam lingkaran 1")
	}else if titik2{
		fmt.Println("Titik di dalam lingkaran 2")
	}else{
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
