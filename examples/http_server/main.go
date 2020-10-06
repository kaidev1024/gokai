package main

import (
	"fmt"
	"net"
	"os"
)

func errFunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		os.Exit(1)
	}
}

func main() {
	// type 127.0.0.1:8000 in browser
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	errFunc(err, "tcp setup failed")
	defer listener.Close()

	conn, err := listener.Accept()
	errFunc(err, "tcp accept error")
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if n == 0 {
		return
	}
	errFunc(err, "conn.Read")

	fmt.Println(string(buf[:n]))
}
