package main

import "fmt"

func calc(x, y int) (result1 int, result2 int, result3 int) {
	result1 = x + y
	result2 = x - y
	result3 = 1
	return
}
func main() {
	result1, result2, a := calc(42, 13)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(a)

	// ブランクで不要な戻り値を捨てる
	_, _, a = calc(5, 1)
	fmt.Println(a)
}
