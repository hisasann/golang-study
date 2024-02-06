package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])

	a[0] = 0
	fmt.Println(a[0])

	// 要素数を省略できる
	b := [...]string{"A", "B", "C"}
	fmt.Println(b)

	c := []string{"AA", "BB", "CC"}
	fmt.Println(c)
}
