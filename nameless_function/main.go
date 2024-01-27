package main

// via https://qiita.com/elephant_dev/items/64e301a5668d6e593429

import "fmt"

func main() {
	// 即時関数
	func() {
		println("Hello, 世界")
	}()

	// 関数を変数に代入する
	value := func() int {
		return 1
	}
	fmt.Println(value())

	// 即時関数に引数を渡す
	func(value string) {
		fmt.Println(value)
	}("hoge")

	// 無名関数を引数に渡す
	value2 := func(p, q string) string {
		return p + q + "Go"
	}
	GFG(value2)

	// 戻り値で関数を返却する
	value3 := GFG2()
	fmt.Println(value3("Welcome ", "to "))
}

// GFG 無名関数を引数に渡す
func GFG(i func(p, q string) string) {
	fmt.Println(i("Go", "for"))
}

// GFG2 無名関数を返却
func GFG2() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "Golang"
	}
	return myf
}
