同じディレクトリに複数の main 関数がある場合、 VSCode 上はエラーにってしまうが、

```go
// +build ignore
```

と追加して保存すると、以下のようになってエラーは消える。

```go
//go:build ignore
// +build ignore
```

[Go メモ-257 (複数の main 関数を共存させる)(go:build ignore) - いろいろ備忘録日記](https://devlights.hatenablog.com/entry/2022/10/14/073000)

Task を使ってみる。

[ホーム | Task](https://taskfile.dev/ja-JP/)

```bash
brew install go-task
```

以下なしの場合。

```go
//go:build ignore
// +build ignore
```

```
$ task
task: [default] go build
# hisasann/golang-study/duplication_main
./main2.go:5:6: main redeclared in this block
        ./main1.go:5:6: other declaration of main
task: [default] go run main1.go
Hello, world!1task: [default] go run main2.go
Hello, world!2
```

以下ありの場合。

```go
//go:build ignore
// +build ignore
```

```
$ task
task: [default] go build
package hisasann/golang-study/duplication_main: build constraints exclude all Go files in /Users/hisasann/_/code/golang/golang-study/duplication_main
task: [default] go run main1.go
Hello, world!1task: [default] go run main2.go
Hello, world!2
```

## 参考記事

[How can I "go run" a project with multiple files in the main package? - Stack Overflow](https://stackoverflow.com/questions/28081486/how-can-i-go-run-a-project-with-multiple-files-in-the-main-package/28081554#28081554)