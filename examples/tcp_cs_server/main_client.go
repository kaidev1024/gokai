package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("are you ready"))

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read err: ", err)
		return
	}
	fmt.Println("server send: ", string(buf[:n]))
}
