package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(conn net.Conn, filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open err: ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("file sent")
			} else {
				fmt.Println("os.Open err: ", err)
			}
			return
		}
		_, err = conn.Write(buf[:n])
	}
}

func main() {
	list := os.Args

	if len(list) != 2 {
		fmt.Println("go run main.go [filename]")
		return
	}

	filePath := list[1]
	fileInfo, _ := os.Stat(filePath)

	fileName := fileInfo.Name()

	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(fileName))
	if err != nil {
		fmt.Println("write err: ", err)
		return
	}

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read err: ", err)
		return
	}

	if "ok" == string(buf[:n]) {
		sendFile(conn, filePath)
	}
}
