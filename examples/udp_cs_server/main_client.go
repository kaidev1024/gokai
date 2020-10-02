package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8006")
	if err != nil {
		fmt.Println("Dial err: ", err)
		return
	}
	defer conn.Close()

	for {
		conn.Write([]byte("are you ready???"))

		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		fmt.Println("server sent: ", string(buf[:n]))
		time.Sleep(time.Second)
	}
}
