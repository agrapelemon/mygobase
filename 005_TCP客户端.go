package main

import "net"

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8888")
	defer conn.Close()
	conn.Write([]byte("dfaodgagjahei黑黑阿斯发达的法sdgfsdfg345"))
}
