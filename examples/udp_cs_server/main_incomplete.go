package main

import (
	"fmt"
	"net"
	"os"
)

func recvFile(conn net.Conn, fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("file create err: ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("finish")
		}
		f.Write(buf[:n])
	}

	conn.Read(buf)
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("net.listen err: ", err)
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("connection err: ", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read err: ", err)
		return
	}

	conn.Write([]byte("OK"))

	recvFile(conn)
}
