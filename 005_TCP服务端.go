package main

import (
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:8888")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			continue
		}
		func() {
			defer conn.Close()
			buf := make([]byte, 1024)
			var n int = 0
			n, err = conn.Read(buf)
			if err != nil {
				fmt.Println("err = ", err)
				return
			}
			fmt.Println("buf = ", string(buf[:n]))
		}()
	}

}
