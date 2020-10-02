package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}
	defer conn.Close()

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				fmt.Println("stdin.read err: ", err)
				continue
			}
			conn.Write(buf[:n])
		}
	}()

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn read")
			return
		}
		fmt.Println("client received from server: ", string(buf[:n]))
	}
}
