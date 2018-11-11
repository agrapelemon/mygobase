package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RemoteAddr, "成功连接")
		fmt.Println("method = ", r.Method)
		fmt.Println("url", r.URL.Path)
		fmt.Println("header = ", r.Header)
		fmt.Println("body = ", r.Body)
		fmt.Fprintln(w, "hello world")
	})
	http.ListenAndServe("127.0.0.1:8888", nil)
}
