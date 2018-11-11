package main

import (
	"fmt"
	"net/http"
)

func main() {
	r, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	defer r.Body.Close()
	fmt.Println(r.Status)
	fmt.Println(r.StatusCode)
	fmt.Println(r.Header)
	fmt.Println(r.Body)
	buf := make([]byte, 4*1024)
	var tmp string
	for {
		n, err := r.Body.Read(buf)
		if n == 0 {
			fmt.Println(err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("tmp = ", tmp)
}
