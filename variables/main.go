package main

import (
	"fmt"
	"math"
)

func main() {
	// 変数を定義しつつ代入
	var num = 1
	fmt.Println(num)

	// 変数を定義しつつ代入しつつデータ型も記述
	// 基本的にはデータ型は書かなくても型推論してくれるので省略可能
	var num2 int = 2
	fmt.Println(num2)

	// var を省略して定義しつつ代入
	num3 := int(math.Pow(2, 8))
	fmt.Println(num3)
}
