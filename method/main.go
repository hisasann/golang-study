package main

import (
	"hisasann/golang-study/logger"
	"strconv"
)

type Hoge struct {
	a int
	b string
}

// Hoge 構造体の foo メソッドを定義
// (h Hoge) はレシーバーと呼ばれる
func (h Hoge) foo() {
	// https://qiita.com/quicksort/items/c9522793a941edf074fd
	logger.WithTrace(strconv.Itoa(h.a))
	logger.WithTrace(h.b)
}

func main() {
	logger.WithTrace("Hello, 世界")

	hoge := Hoge{a: 1, b: "hoge"}
	hoge.foo()
}
