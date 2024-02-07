package main

import "fmt"

type Hoge struct {
	a int
	b string
}

func main() {
	a := "a"
	foo(&a)

	var hoge Hoge
	// Hogeを値として代入する
	hoge = Hoge{a: 2, b: "fuga"}
	fuga(&hoge)

	var bar *Hoge
	// Hogeをポインタ型として代入する
	bar = &Hoge{a: 3, b: "bar"}
	fuga(bar)

}

func foo(a *string) {
	fmt.Println(*a)
}

func fuga(hoge *Hoge) {
	fmt.Println(hoge.a)
	fmt.Println(hoge.b)
}
