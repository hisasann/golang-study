package main

import "fmt"

// 構造体
type Hoge struct {
	a int
	b string
}

func main() {
	// アンバサンドを使うとポインタ型になる
	foo(&Hoge{a: 1, b: "hoge"})

	var hoge Hoge
	hoge.a = 2
	hoge.b = "fuga"
	foo(&hoge)
}

// *Hoge はポインタ型
func foo(input *Hoge) {
	fmt.Println(input.a)
	fmt.Println(input.b)
}
