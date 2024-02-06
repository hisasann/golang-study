package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	// レスポンスとしてHello, World!!を返す
	w.Write([]byte("Hello, World!!"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// http://localhost:8080/hello でアクセスできる
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	})
	server.ListenAndServe()
}
