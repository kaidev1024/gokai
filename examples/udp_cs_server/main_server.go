package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8006")
	if err != nil {
		fmt.Println("ResolveUDPAddr err: ", err)
		return
	}
	fmt.Println("Udp server created")

	udpConn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		fmt.Println("ResolveUDPAddr err: ", err)
		return
	}
	defer udpConn.Close()

	buf := make([]byte, 4096)

	n, cltAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("ResolveUDPAddr err: ", err)
		return
	}

	fmt.Printf("server read %v data: %s\n", cltAddr, string(buf[:n]))

	daytime := time.Now().String()
	_, err = udpConn.WriteToUDP([]byte(daytime), cltAddr)
	if err != nil {
		fmt.Println("ResolveUDPAddr err: ", err)
		return
	}
}
