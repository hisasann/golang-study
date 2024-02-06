package main

import f "fmt"

func main() {
	// 遅延実行させられる
	// 関数の終了時に実行される、または return で関数を抜けるときに実行される
	// trycatchのfinallyのようなもの
	defer f.Println("Hello, 世界2")
	// 即時実行関数で遅延実行させる
	defer func() {
		// LIFOなので、一番最後に defer したものが一番最初に実行される
		f.Println("Hello, 世界3")
	}()

	println("Hello, 世界1")
}
