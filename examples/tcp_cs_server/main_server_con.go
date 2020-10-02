// concurrent server
package main

import (
	"fmt"
	"net"
	"strings"
)

func HandlerConnect(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr()
	fmt.Println(addr, " established")

	buf := make([]byte, 4096)
	for {

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err: ", err)
			return
		}
		fmt.Println("got data: ", string(buf[:n]))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("err creating listener: ", err)
		return
	}
	fmt.Println("listener created...")
	defer listener.Close()

	for {
		fmt.Println("server waiting for client")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err accepting request: ", err)
			return
		}

		go HandlerConnect(conn)
	}
}
