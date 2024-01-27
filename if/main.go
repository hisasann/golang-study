package main

import (
	"errors"
	"fmt"
)

func main() {
	// エラーがない
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// エラーがある場合
	if err := run2(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	return nil
}

func run2() error {
	return errors.New("error!")
}
