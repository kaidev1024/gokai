package main

import (
	"fmt"
	"net"
)

func main() {
	//socket for listening
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	defer listener.Close()
	fmt.Println("server waiting for client to connect...")
	//conn is socket for communication
	// type this in command line: nc 127.0.0.1 8000
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept() err: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("connection established...")

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err: ", err)
		return
	}
	conn.Write(buf[:n])
	fmt.Println("server got data", string(buf[:n]))

}
