package logger

import "fmt"

// オーバーロードはない
// logger.WithTrace("Hello, 世界")のように呼び出せる
func WithTrace(message string) {
	fmt.Println(message)
}
