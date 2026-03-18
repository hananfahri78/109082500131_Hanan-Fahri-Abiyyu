package main

import "fmt"

func f(x int) int {
	pangkat := x * x
	return pangkat
}

func g(x int) int {
	fungxie := x - 2
	return fungxie
}

func h(x int) int {
	function := x + 1
	return function
}
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}
